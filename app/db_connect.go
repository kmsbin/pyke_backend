package main

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
)

var once sync.Once

var dbConn *sql.DB

func getDBInstance() *sql.DB {
	if dbConn == nil {
		once.Do(
			func() {
				var err error
				fmt.Println(os.Getenv("DATABASE_URL"))
				dbConn, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
				if err != nil {
					panic(err)
				}

				fmt.Println("Creting Single Instance Now")
			})
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return dbConn
}

func killInstance() {
	defer dbConn.Close()
}
