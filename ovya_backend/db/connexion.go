package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbName)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			fmt.Printf("Attempt %d : opening error : %v\n", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		if err = db.Ping(); err == nil {
			fmt.Println("Connexion established")
			return db, nil
		}

		fmt.Printf("Attempt %d : ping failed : %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to the database after several attempts : %v", err)
}
