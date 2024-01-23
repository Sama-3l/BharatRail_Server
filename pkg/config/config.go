package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CreateConnection() *sql.DB {
	str, _ := os.Getwd()
	str = str + "\\.env"
	err := godotenv.Load(str)

	if err != nil {
		log.Printf("Error loading .env file: \n%v \n", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))

	if err != nil {
		log.Printf("Error Connecting to postgres server: \n%v \n", err)
	}

	err = db.Ping()

	if err != nil {
		log.Printf("Can't ping database: \n%v \n", err)
	}

	return db
}
