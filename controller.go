package main

import (
	"encoding/json"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = map[string]string{
	"user1": "pass1",
	"user2": "pass2"}

func Signup(w http.ResponseWriter, r *http.Request) {
	var newuser Credentials
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if users[newuser.Username] != "" {
		w.WriteHeader(http.StatusIMUsed)
		w.Write([]byte("username already taken, try another one"))
		return
	}
	users[newuser.Username] = newuser.Password
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully created user"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginuser Credentials
	err := json.NewDecoder(r.Body).Decode(&loginuser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if loginuser.Password == users[loginuser.Username] {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("login successfull"))
	}

}

func Refresh(w http.ResponseWriter, r *http.Request) {

}
