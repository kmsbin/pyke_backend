package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	defer killInstance()
	getDBInstance()
	r := mux.NewRouter()

	r.Handle("/auth", AuthMiddleware(usersHandler))
	r.Handle("/auth/{access_token}", AuthMiddleware(usersHandler))

	r.HandleFunc("/users", usersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", userGetHandler).Methods("GET")
	r.HandleFunc("/user/", registerUser).Methods("POST")
	r.HandleFunc("/login/", loginUser).Methods("POST")
	r.HandleFunc("/history/{id}", historyHandler).Methods("GET")
	r.HandleFunc("/history/", historyRegister).Methods("POST")
	r.HandleFunc("/history/{id}", updateHistory).Methods("PUT")

	http.ListenAndServe(string(":"+os.Getenv("PORT")), r)
	// http.ListenAndServe(":8080", r)
}
