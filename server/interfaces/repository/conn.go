package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func CreatePostgresConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=postgres user=%s dbname=%s password=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"))
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
