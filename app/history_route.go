package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type history struct {
	IDHistory      int     `json:"id_history"`
	IsFarorite     bool    `json:"is_favorite"`
	UserID         int     `json:"id_user"`
	LatitudeFrom   float64 `json:"latitude_from"`
	LongitudeFrom  float64 `json:"longitude_from"`
	LatitudeWhere  float64 `json:"latitude_where"`
	LongitudeWhere float64 `json:"longitude_where"`
}
type historys struct {
	Values []history `json:"historyRoute"`
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	db := getDBInstance()

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	rows, err := db.Query(`SELECT * FROM route_history WHERE id_user = $1`, id)
	if err != nil {
		panic(err)
	}
	var hists historys
	for rows.Next() {
		var his history
		rows.Scan(&his.IDHistory, &his.IsFarorite, &his.UserID, &his.LatitudeFrom, &his.LongitudeFrom, &his.LatitudeWhere, &his.LongitudeWhere)
		fmt.Println(his)
		hists.Values = append(hists.Values, his)
	}
	jsonCode, _ := json.Marshal(hists)
	// fmt.Println(jsonCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonCode)
}
func historyRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("Egasdgsd")
	db := getDBInstance()
	var his history
	err := json.NewDecoder(r.Body).Decode(&his)
	fmt.Println(his)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlstmt := `INSERT INTO route_history(is_favorite, id_user, latitude_from,longitude_from, latitude_where, longitude_where) VALUES ($1, $2, $3,$4, $5, $6);`
	db.Exec(sqlstmt, his.IsFarorite, his.UserID, his.LatitudeFrom, his.LongitudeFrom, his.LatitudeWhere, his.LongitudeWhere)
	w.WriteHeader(200)
}
func updateHistory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	db := getDBInstance()

	db.Exec(`UPDATE route_history SET is_favorite = true WHERE id_history = $1 ;`, id)
	w.WriteHeader(200)
}
