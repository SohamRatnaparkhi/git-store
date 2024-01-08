package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *database.Queries {
	//database connection
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		db_url = "postgres://postgres:foobarbaz@localhost:5432/git_store?sslmode=disable"
	}

	db, dbErr := sql.Open("postgres", db_url)

	if dbErr != nil {
		return nil
	}
	dbQueries := database.New(db)
	return dbQueries
}

// var DbClient *database.Queries = ConnectDB()
