package main

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"server/db"
	"server/handler"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed ui/*
var uiFS embed.FS

const dataJSONURL = "https://joseph0x45.github.io/mom_business_data/data.json"

func setupData() {
  log.Println("[INFO]: Initializing data")
	req, err := http.NewRequest(
		"GET",
		dataJSONURL,
		nil,
	)
	if err != nil {
		log.Panicf("[ERROR]: Failed to create HTTP request: %s\n", err.Error())
	}
	client := http.Client{
		Timeout: time.Minute,
	}
	response, err := client.Do(req)
	if err != nil {
		log.Panicf("[ERROR]: Failed to send HTTP request: %s\n", err.Error())
	}
	if response.StatusCode != 200 {
		log.Panicf("[ERROR]: Expected HTTP 200 got: %d\n", response.StatusCode)
	}
	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Panicf("[ERROR]: Failed to read response body: %s\n", err.Error())
	}
	// write json to file
	jsonFile, err := os.Create("./data.json")
	if err != nil {
		log.Panicf("[ERROR]: Failed to create data file: %s\n", err.Error())
	}
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Panicf("[ERROR]: Failed to write data to JSON file : %s\n", err.Error())
	}
  log.Println("[INFO]: Done!")
}

func main() {
	setupData()
  return

	conn := db.NewDBConnection("db.sqlite")
	conn.Migrate()
	conn.SeedData()

	handler := handler.NewHandler(conn, uiFS)

	staticContent, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Println("[ERROR] Error while loading static files")
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.FS(staticContent))))
	mux.HandleFunc("GET /home", handler.RenderHomePage)

	mux.HandleFunc("GET /api/groups", handler.FetchGroups)

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
