package daemon

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sdorra/jasas/auth"
	"github.com/twinj/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Daemon struct {
	Authenticator auth.Authenticator
	key           jose.JSONWebKey
	jwkPath       string
	Domain        string
	CookieName    string
}

func New(authenticator auth.Authenticator) (*Daemon, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate private rsa key")
	}

	domain := env("JASAS_DOMAIN", "example.net")
	keyID := env("JASAS_KEYID", "auth."+domain)

	key := jose.JSONWebKey{
		Key:       privateKey,
		KeyID:     keyID,
		Algorithm: string(jose.RS256),
	}

	keySet := jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{key},
	}

	data, err := json.Marshal(keySet)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal jwk key set")
	}

	jwkPath := env("JASAS_JWK_PATH", "keyset.jwk")
	err = ioutil.WriteFile(jwkPath, data, 0700)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write jwk key set")
	}
	cookie := env("JASAS_COOKIE", "X-JASAS-Token")
	return &Daemon{
		Authenticator: authenticator,
		key:           key,
		jwkPath:       jwkPath,
		Domain:        domain,
		CookieName:    cookie,
	}, nil
}

func (daemon *Daemon) Start() error {
	r := mux.NewRouter()

	r.HandleFunc("/v1/authentication", daemon.AuthenticationHandler)
	r.HandleFunc("/v1/logout", daemon.LogoutHandler).Methods("POST")
	r.HandleFunc("/v1/validation", daemon.ValidationHandler)

	r.HandleFunc("/", daemon.RootPage).Methods("GET")
	r.HandleFunc("/", daemon.Login).Methods("POST")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return errors.Wrap(err, "http server returned error")
	}
	return nil
}

func (daemon *Daemon) ValidateToken(token string) (string, error) {
	signature, err := jose.ParseSigned(token)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse token")
	}

	key := daemon.key.Key
	switch key.(type) {
	case *rsa.PrivateKey:
		key = key.(*rsa.PrivateKey).Public()
	}

	payload, err := signature.Verify(key)
	if err != nil {
		return "", errors.Wrap(err, "could not verify token")
	}

	entries := make(map[string]interface{})
	err = json.Unmarshal(payload, &entries)
	if err != nil {
		return "", errors.Wrap(err, "could not unmarshal payload")
	}

	return entries["sub"].(string), nil
}

func (daemon *Daemon) CreateToken(subject string) (string, error) {
	algorithm := jose.SignatureAlgorithm(daemon.key.Algorithm)

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: algorithm, Key: daemon.key.Key}, &jose.SignerOptions{})
	if err != nil {
		return "", errors.Wrap(err, "failed to create jwt signer")
	}

	now := time.Now()
	expires := now.Add(24 * time.Hour)
	cl := jwt.Claims{
		ID:       uuid.NewV4().String(),
		Subject:  subject,
		IssuedAt: jwt.NewNumericDate(now),
		Expiry:   jwt.NewNumericDate(expires),
	}

	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", errors.Wrap(err, "failed to sign claim")
	}

	return raw, nil
}

func env(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
