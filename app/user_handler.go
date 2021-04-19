package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func userGetHandler(w http.ResponseWriter, r *http.Request) {

	db := getDBInstance()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	rows, err := db.Query(`SELECT * FROM users WHERE id_user = $1`, id)
	if err != nil {
		panic(err)
	}
	var users users
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
func registerUser(w http.ResponseWriter, r *http.Request) {

	var newUser user
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Println(newUser)

	db := getDBInstance()
	// var id int
	row, _ := db.Query(`SELECT * FROM users WHERE email = $1`, newUser.Email)
	if row.Next() {
		var httpErro errHTTP
		httpErro.HTTPError.Message = "this email already registered"

		jsonEncoded, _ := json.Marshal(httpErro)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonEncoded)
		w.WriteHeader(400)
		return
	}

	config := &PasswordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 2,
		keyLen:  32,
	}
	// Example 1: Generating a Password Hash
	hash, errGenerate := GeneratePassword(config, newUser.Password)
	if errGenerate != nil {
		fmt.Println(errGenerate)
	}
	fmt.Println(hash)

	sqlstmt := `INSERT INTO users(user_name, email, password_key) VALUES ( $1, $2, $3)`
	_, err = db.Exec(sqlstmt, newUser.Name, newUser.Email, hash)

	if err != nil {
		panic(err)
	}
	jsonEncode, _ := json.Marshal(newUser)
	w.Write(jsonEncode)
	w.WriteHeader(200)
}

type errHTTP struct {
	HTTPError message `json:"error"`
}
type message struct {
	Message string `json:"message"`
}

func loginUser(w http.ResponseWriter, r *http.Request) {

	var newUser user
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Println(newUser)

	db := getDBInstance()
	// var id int
	row, _ := db.Query(`SELECT * FROM users WHERE email = $1`, newUser.Email)
	if row.Next() {
		var u user
		row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
		isValite, _ := VerifyPassword(newUser.Password, u.Password)
		if isValite {
			fmt.Println("SUJEITO VALIDADO")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			jsonEncode, _ := json.Marshal(u)
			w.Write(jsonEncode)
			return
		} else {
			var err errHTTP

			err.HTTPError.Message = "wrong password"
			jsonEncode, _ := json.Marshal(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			w.Write(jsonEncode)
		}
		return
	}
	var errCred errHTTP
	w.Header().Set("Content-Type", "application/json")
	errCred.HTTPError.Message = "not registered email"
	jsonEncode, _ := json.Marshal(errCred)
	w.WriteHeader(400)

	w.Write(jsonEncode)
}
