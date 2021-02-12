package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func init() {
	os.Setenv("PORT", "8080")
	os.Setenv("DATABASE_URL", "postgres://dljcxcutmuyvce:de66310f2b01549519c9fe4267722691e0e87ebca6ccf80e3e1f5a97f37f1b84@ec2-54-164-241-193.compute-1.amazonaws.com:5432/d8f9megkkbe7ij")
}
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
