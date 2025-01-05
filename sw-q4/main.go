package main

import (
	"database/sql"
	"log"
	"myapp/handlers"
	"myapp/repositories"
	"myapp/services"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

// Main Application Entry Point
// Sets up database connection, routing, and starts the HTTP server

func main() {
	// Initialize database connection
	db, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Verify database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("Database connection successful!")

	// Create database schema
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	log.Println("Table 'users' created successfully!")

	// Initialize application layers
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Set up routing
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)

	// Start HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
