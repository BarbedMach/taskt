package main

import (
	_ "github.com/lib/pq"
)

func main() {
	initDB()

	createTableFromSQL(db, "users.sql")
	createTableFromSQL(db, "tasks.sql")

	closeDB()
}
