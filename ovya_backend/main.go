package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"ovya_backend/db"
	"ovya_backend/middleware"
	"ovya_backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "docker" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Erreur de chargement du fichier .env")
		}
	}
	db, err := db.DbConnection()
	if err != nil {
		fmt.Println("Connexion failed :", err)
		return
	}
	defer db.Close()

	fmt.Println("Server started succesfully")

	mux := http.NewServeMux()

	routes.RegisterAllRoutes(db, mux)

	handlerWithCors := middleware.WithCORS(mux)

	log.Fatal(http.ListenAndServe(":8080", handlerWithCors))

	// http.ListenAndServe(":8080", handlerWithCors)

}
