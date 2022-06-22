package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwt_key = []byte("mysecret key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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

	//Validate Login
	if loginuser.Password != users[loginuser.Username] {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("incorect login details"))
		return
	}
	//Create a jwt token and set it to cookie
	expiration_time := time.Now().Add(5 * time.Minute)

	claim := Claims{
		Username: loginuser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_time.Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signed_token, err := token.SignedString(jwt_key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("claim is ", claim)
	fmt.Println("token is ", token)
	fmt.Println("signed_token is ", signed_token)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   signed_token,
		Expires: expiration_time,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(signed_token))
}

func Refresh(w http.ResponseWriter, r *http.Request) {

}
