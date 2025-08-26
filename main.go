package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"server/db"
	"server/handler"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed ui/*
var uiFS embed.FS

func main() {
	conn := db.NewDBConnection("db.sqlite")
	conn.Migrate()
	conn.SeedGroups()

	handler := handler.NewHandler(conn, uiFS)

	staticContent, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Println("[ERROR] Error while loading static files")
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.FS(staticContent))))
	mux.HandleFunc("GET /home", handler.RenderHomePage)

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
