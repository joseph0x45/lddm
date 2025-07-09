package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"server/handlers"
	"server/repo"
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
		panic("Could not find DB URL")
	}
	db, err := sqlx.Connect("sqlite3", dbURL)
	if err != nil {
		panic(err)
	}
	log.Println("[INFO] Connected to Database")
	return db
}

func main() {
	db := getDBConnection()
	staticContent, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Println("[ERROR] Error while loading static files")
		panic(err)
	}

	productsRepo := repo.NewProductRepo(db)

	productsHandler := handlers.NewProductHandler(productsRepo)

	uiHandler := handlers.NewUIHandler(&templatesFS)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticContent))))
	mux.HandleFunc("POST /products", productsHandler.CreateProduct)
	mux.HandleFunc("GET /products", productsHandler.GetAllProducts)

	mux.HandleFunc("GET /login", uiHandler.RenderLoginPage)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	log.Println("[INFO] Server listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
