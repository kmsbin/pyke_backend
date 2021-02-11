package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {

	db := getDBInstance()
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

}
