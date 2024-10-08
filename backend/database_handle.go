package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var db *sql.DB

func generateConnectionString() string {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("DB_HOST")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)
}

func initDB() {
	connStr := generateConnectionString()

	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Error connecting to database: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			break
		}

		log.Printf("Error pinging database: %v", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}

func closeDB() {
	db.Close()
}

func createTableFromSQL(db *sql.DB, file string) {
	filepath := "./db/tables/" + file

	sqlFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("Table created successfully!")
}
