package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	getDBInstance()
	defer killInstance()
	r := mux.NewRouter()

	r.HandleFunc("/users", usersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", userGetHandler).Methods("GET")
	r.HandleFunc("/user/", registerUser).Methods("POST")
	r.HandleFunc("/login/", loginUser).Methods("POST")

	http.ListenAndServe(string(":"+os.Getenv("PORT")), r)
}
