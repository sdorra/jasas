package daemon

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type credentials struct {
	Username string
	Password string
}

var errBodyRequired = errors.New("body required")
var errUsernameAndPasswordRequired = errors.New("username and password are required")

// AuthenticationHandler authenticates the client and creates an jwt token
func (daemon *Daemon) AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		handleFailure(w, errBodyRequired, 400)
		return
	}

	var credentials credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		handleFailure(w, err, 400)
		return
	}

	if credentials.Username == "" && credentials.Password == "" {
		handleFailure(w, errUsernameAndPasswordRequired, 400)
		return
	}

	err = daemon.Authenticator.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		handleFailure(w, err, 400)
		return
	}

	token, err := daemon.CreateToken(credentials.Username)
	if err != nil {
		handleFailure(w, err, 500)
		return
	}

	cookie := http.Cookie{
		Name:     daemon.CookieName,
		Path:     "/",
		Value:    token,
		Domain:   daemon.Domain,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	handleSuccess(w, credentials.Username)
}
