package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"csv-importer/internal/config"
)

func Connect(cfg config.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.MySQLDSN)
	if err != nil {
		log.Fatal("DB connect error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	log.Println("Connected to MySQL")
	return db
}
