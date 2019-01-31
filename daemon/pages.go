package daemon

import (
	"github.com/pkg/errors"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

var templatePath = path.Join("daemon", "templates")

type Page struct {
	Title    string
	SubTitle string
}

func (daemon *Daemon) RootPage(w http.ResponseWriter, r *http.Request) {
	log.Println("root page")

	username := daemon.authenticationFromCookie(r)

	if username != "" {
		daemon.renderWelcomePage(w, username)
		return
	}

	daemon.renderLoginPage(w)
}

func (daemon *Daemon) authenticationFromCookie(r *http.Request) string {
	cookie, err := r.Cookie(daemon.CookieName)
	if err != nil {
		return ""
	}

	username, err := daemon.ValidateToken(cookie.Value)
	if err != nil {
		return ""
	}

	return username
}

func (daemon *Daemon) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("login")

	err := r.ParseForm()
	if err != nil {
		handleLoginFailure(w, "", err, 400)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		handleLoginFailure(w, "", errors.New("username and password are required"), 400)
		return
	}

	err = daemon.Authenticator.Authenticate(username, password)
	if err != nil {
		handleLoginFailure(w, username, err, 400)
		return
	}

	token, err := daemon.CreateToken(username)
	if err != nil {
		handleLoginFailure(w, username, err, 500)
		return
	}

	cookie := http.Cookie{
		Name:  daemon.CookieName,
		Path:  "/",
		Value: token,
		// Domain:   daemon.Domain,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	daemon.renderWelcomePage(w, username)
}

func (daemon *Daemon) renderWelcomePage(w http.ResponseWriter, username string) {
	tmpl := parseTemplate("welcome")

	model := createBaseModel()
	model["Username"] = username

	tmpl.ExecuteTemplate(w, "layout", model)
}

func (daemon *Daemon) renderLoginPage(w http.ResponseWriter) {
	tmpl := parseTemplate("login")
	tmpl.ExecuteTemplate(w, "layout", createBaseModel())
}

func handleLoginFailure(w http.ResponseWriter, username string, err error, statusCode int) {
	log.Println("handle login failure", err)

	tmpl := parseTemplate("login")

	model := createBaseModel()
	model["Error"] = err
	model["Username"] = username

	log.Println(model)

	w.WriteHeader(statusCode)
	tmpl.ExecuteTemplate(w, "layout", model)
}

func createBaseModel() map[string]interface{} {
	model := make(map[string]interface{})
	model["Page"] = Page{
		Title:    "Jasas",
		SubTitle: "Just another simple authentication service",
	}
	return model
}

func parseTemplate(name string) *template.Template {
	layout := path.Join(templatePath, "layout.html")
	page := path.Join(templatePath, name+".html")

	return template.Must(template.ParseFiles(layout, page))
}
