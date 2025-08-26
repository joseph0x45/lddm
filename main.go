package main

import (
	"log"
	"net/http"
	"server/db"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mux := http.NewServeMux()
	conn := db.NewDBConnection("db.sqlite")
	conn.Migrate()
	conn.SeedGroups()

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	log.Println("[INFO] Server listening on port 8080\nVisit http://0.0.0.0:8080/home")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
