package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

users := map[string]string{
	"user1": "pass1",
	"user2": "pass2"}

func main() {

	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/refresh", refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func signup(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {
	var loginuser user
	err := json.NewDecoder(r.Body).Decode(&loginuser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if loginuser.Password == users[loginuser.Username]{
        w.WriteHeader(http)
	}
}

func refresh(w http.ResponseWriter, r *http.Request) {

}
