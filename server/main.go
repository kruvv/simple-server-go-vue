package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Структура User представляет одного пользователя
type User struct {
	Name string `json:"name"`
}

var users = []User{
	{Name: "Alex"},
	{Name: "Bob"},
	{Name: "Elis"},
}

func main() {
	http.HandleFunc("/users", handleUsersRequest)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Настройка CORS
func setCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
}

// Обработчик запросов к /users
func handleUsersRequest(w http.ResponseWriter, r *http.Request) {

	setCorsHeaders(w) // Устанавливаем заголовки CORS

	switch r.Method {

	case "GET":
		handleGetUsers(w, r)
	case "POST":
		handlePostUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

// GET-запросы возвращают всех пользователей в JSON
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// POST-запросы добавляют нового пользователя
func handlePostUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User added successfully!")
}
