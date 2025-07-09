package main

import (
	"log"
	"net/http"
	"os"
	"server/handlers"
	"server/repo"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

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

	productsRepo := repo.NewProductRepo(db)

	productsHandler := handlers.NewProductHandler(productsRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", productsHandler.CreateProduct)
	mux.HandleFunc("GET /products", productsHandler.GetAllProducts)

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
