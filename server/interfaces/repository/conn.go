package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func CreatePostgresConnection() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("dbname"), os.Getenv("password"))
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
