package daemon

import (
	"github.com/pkg/errors"
	"github.com/sdorra/jasas/templates"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

var templatePath = "templates"

type Page struct {
	Title    string
	SubTitle string
}

func (daemon *Daemon) RootPage(w http.ResponseWriter, r *http.Request) {
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

func (daemon *Daemon) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  daemon.CookieName,
		Path:  "/",
		Value: "deleted",
		// Domain:   daemon.Domain,
		Expires:  time.Date(1970, time.January, 1, 1, 0, 0, 0, time.UTC),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	tmpl := parseTemplate("logout")
	tmpl.ExecuteTemplate(w, "layout", createBaseModel())
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
	tmpl := parseTemplate("login")

	model := createBaseModel()
	model["Error"] = err
	model["Username"] = username

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

	template, err := vfstemplate.ParseFiles(templates.Templates, nil, layout, page)
	if err != nil {
		log.Fatal("failed to parse template", err)
	}

	return template
}
