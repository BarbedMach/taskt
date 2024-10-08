package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func insertUserFromHTTPReq(writer http.ResponseWriter, req *http.Request) {
	var user User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Invalid Input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (username, email, pswd, created_at) VALUES ($1, $2, $3, $4)`

	_, err = db.Exec(query, user.Username, user.Email, user.Password, time.Now())
	if err != nil {
		http.Error(writer, "Could not add user", http.StatusInternalServerError)
		log.Println("Error adding user: ", err)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(map[string]string{"message": "User created successfully"})
	fmt.Println("User added successfully!")
}

func insertTaskFromHTTPReq(writer http.ResponseWriter, req *http.Request) {
	var task Task

	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(writer, "Invalid Input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO tasks (title, description, start_time, end_time, status, user_id)
				VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = db.Exec(query, task.Title, task.Description, task.StartTime, task.EndTime, task.Status, task.UserID)
	if err != nil {
		http.Error(writer, "Could not add task", http.StatusInternalServerError)
		log.Println("Error adding task: ", err)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(map[string]string{"message": "User created successfully"})
	fmt.Println("Task added successfully!")
}
