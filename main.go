package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"server/handler"
	"server/store"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed templates/*
var templatesFS embed.FS

func getDBConnection() *sqlx.DB {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "db.sqlite"
	}
	db, err := sqlx.Connect("sqlite3", dbURL)
	if err != nil {
		panic(err)
	}
	log.Println("[INFO] Connected to Database:", dbURL)
	return db
}

func main() {
	db := getDBConnection()
	printerPort := os.Getenv("PRINTER")
	staticContent, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Println("[ERROR] Error while loading static files")
		panic(err)
	}

	store := store.NewStore(db)
	handler := handler.NewHandler(store, &templatesFS, printerPort)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticContent))))
	mux.HandleFunc("GET /home", handler.RenderHomePage)
	mux.HandleFunc("GET /products", handler.RenderProductsPage)
	mux.HandleFunc("GET /orders", handler.RenderOrdersPage)
	mux.HandleFunc("GET /stats", handler.RenderStatsPage)

	mux.HandleFunc("POST /api/products", handler.CreateProduct)
	mux.HandleFunc("POST /api/orders", handler.SaveOrder)
	mux.HandleFunc("DELETE /api/orders/{id}", handler.DeleteOrder)
	mux.HandleFunc("GET /api/orders/{id}/print", handler.PrintOrder)

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
