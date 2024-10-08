package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	initDB()

	createTableFromSQL(db, "users.sql")
	createTableFromSQL(db, "tasks.sql")

	http.HandleFunc("/users", insertUserFromHTTPReq)
	http.HandleFunc("/tasks", insertTaskFromHTTPReq)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

	closeDB()
}
