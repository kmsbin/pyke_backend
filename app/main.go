package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	getDBInstance()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer killInstance()
	r := mux.NewRouter()

	r.HandleFunc("/users", usersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", userGetHandler).Methods("GET")
	r.HandleFunc("/user/", registerUser).Methods("POST")
	r.HandleFunc("/login/", loginUser).Methods("POST")
	r.HandleFunc("/history/{id}", historyHandler).Methods("GET")
	r.HandleFunc("/history/", historyRegister).Methods("POST")
	r.HandleFunc("/history/{id}", updateHistory).Methods("PUT")

	http.ListenAndServe(string(":"+os.Getenv("PORT")), r)
}
