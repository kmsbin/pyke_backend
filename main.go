package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Print("Server rodando liso")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Bateu na rota '/'")
		w.Write([]byte("horário oficial do oleo de macaco: " + time.Now().Format("15:04:05 02-01-2006")))
	})
	log.Fatal(http.ListenAndServe(string(":"+os.Getenv("PORT")), nil))
}
