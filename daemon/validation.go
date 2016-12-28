package daemon

import "net/http"

func (daemon *Daemon) ValidationHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(daemon.CookieName)
	if err != nil {
		handleFailure(w, err, 403)
		return
	}

	username, err := daemon.ValidateToken(cookie.Value)
	if err != nil {
		handleFailure(w, err, 403)
		return
	}

	handleSuccess(w, username)
}
