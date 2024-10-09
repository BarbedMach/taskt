package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	_ "github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

func usersHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fetchUsers(writer, req)
	case http.MethodPost:
		insertUser(writer, req)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func tasksHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fetchTasks(writer, req)
	case http.MethodPost:
		insertTask(writer, req)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	initDB()

	corsObj := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type"})

	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/login", loginUser)
	http.HandleFunc("/signup", insertUser)

	err := http.ListenAndServe(":8080", handlers.CORS(corsHeaders, corsMethods, corsObj)(http.DefaultServeMux))
	if err != nil {
		panic(err)
	}

	closeDB()
}
