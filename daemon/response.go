package daemon

import (
	"encoding/json"
	"net/http"
)

type AuthenticationState struct {
	Username string `json:"username"`
}

type AuthenticationFailure struct {
	Message string `json:"message"`
}

func handleSuccess(w http.ResponseWriter, username string) {
	state := AuthenticationState{username}
	content, err := json.Marshal(&state)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(content)
}

func handleFailure(w http.ResponseWriter, cause error, code int) {
	failure := AuthenticationFailure{cause.Error()}
	data, err := json.Marshal(failure)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(data)
}
