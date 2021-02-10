package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type users struct {
	Users []user `json:"users"`
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Print("Server rodando liso")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Bateu na rota '/'")
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			panic(err)
		}
		var users users // var users []user
		for rows.Next() {
			var u user
			rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
			fmt.Println(u)
			users.Users = append(users.Users, u)
		}
		jsonCode, _ := json.Marshal(users)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonCode)

	})
	fmt.Printf("a porta é essa %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(string(":"+os.Getenv("PORT")), nil))
}
