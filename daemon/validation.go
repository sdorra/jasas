package daemon

import "net/http"

func (daemon *Daemon) ValidationHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(daemon.CookieName)
	if err != nil {
		http.Error(w, "failed to read cookie", 403)
		return
	}

	username, err := daemon.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, "cookie is invalid", 403)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(username))
}
