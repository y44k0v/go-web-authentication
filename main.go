package main

import (
	"fmt"
	"net/http"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// key is the username -TODO update to DB
var users = map[string]Login{}

func main() {
	// entry points
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	http.ListenAndServe(":9080", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	// verify POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// grab username and password from form
	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) < 6 || len(password) < 8 {
		http.Error(w, "Invalid username or password", http.StatusNotAcceptable)
		return
	}
	// check if user already exists
	if _, ok := users[username]; ok {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	HashedPassword, _ := hashPassword(password)
	users[username] = Login{
		HashedPassword: HashedPassword,
	}

	fmt.Fprintf(w, "User %s registered successfully!", username)

}

func login(w http.ResponseWriter, r *http.Request) {}

func logout(w http.ResponseWriter, r *http.Request) {}

func protected(w http.ResponseWriter, r *http.Request) {}
