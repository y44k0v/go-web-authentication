package main

import (
	"net/http"

	"errors"
)

var ErrUnauthorized = errors.New("Unauthorized")

// to use with any portected route
func Authorize(r *http.Request) error {
	username := r.FormValue("username")
	user, ok := users[username]
	if !ok {
		return ErrUnauthorized
	}

	// get the session  token from the cookie
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		return ErrUnauthorized
	}

	// get the csrf token from the header
	csrf := r.Header.Get("X-CSRF-Token")
	if csrf == "" || csrf != user.CSRFToken {
		return ErrUnauthorized
	}

	return nil
}
