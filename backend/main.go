package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura para representar un usuario
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Base de datos en memoria
var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

// Middleware para habilitar CORS
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// Handler para obtener todos los usuarios
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler para crear un nuevo usuario
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Servir archivos estáticos (frontend)
	frontendDir := "../frontend" // Ruta al directorio del frontend
	fs := http.FileServer(http.Dir(frontendDir))
	http.Handle("/", fs)

	// Rutas de la API
	http.HandleFunc("/users", enableCORS(getUsers))          // GET /users
	http.HandleFunc("/users/create", enableCORS(createUser)) // POST /users/create

	// Servidor HTTP
	fmt.Println("Servidor corriendo en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
