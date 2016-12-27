package daemon

import (
	"net/http"
	"time"
)

func (daemon *Daemon) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     daemon.CookieName,
		Path:     "/",
		Value:    "deleted",
		Domain:   daemon.Domain,
		Expires:  time.Date(1970, time.January, 1, 1, 0, 0, 0, time.UTC),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}
