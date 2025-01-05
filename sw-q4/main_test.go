package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"myapp/handlers"
	"myapp/models"
	"myapp/repositories"
	"myapp/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db          *sql.DB
	userHandler *handlers.UserHandler
)

// --- INITIALIZATION ---

// Initialize the test environment
func init() {
	var err error
	db, err = sql.Open("sqlite3", ":memory:") // In-memory SQLite database
	if err != nil {
		log.Fatal(err)
	}
}

// --- SETUP FUNCTIONS ---

// Prepare the test database
func setupTestDatabase(t *testing.T) {
	// Drop existing tables
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		t.Fatalf("Failed to drop tables: %v", err)
	}

	// Create the users table
	createTableSQL := `CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		t.Fatalf("Failed to create tables: %v", err)
	}

	// Insert test data
	insertUserSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
	_, err = db.Exec(insertUserSQL, "Alice", "alice@example.com")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}
}

// Prepare the handler
func setupHandler(t *testing.T) {
	setupTestDatabase(t) // Set up the test database
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler = handlers.NewUserHandler(userService)
}

// --- TEST FUNCTIONS ---

// Test for getting all users
func TestGetAllUsers(t *testing.T) {
	setupHandler(t)

	t.Log("Sending GET request to fetch all users...")
	req := httptest.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(userHandler.GetAllUsers)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		t.Logf("GET /users returned status code: %v", rr.Code)
	}

	var users []models.User
	if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	t.Logf("Fetched users: %+v", users)
	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}
}

// Test for getting a user by ID
func TestGetUserByID(t *testing.T) {
	setupHandler(t)

	t.Log("Sending GET request to fetch user with ID 1...")
	req := httptest.NewRequest("GET", "/users/1", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.GetUserByID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		t.Logf("GET /users/1 returned status code: %v", rr.Code)
	}

	var user models.User
	if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	t.Logf("Fetched user: %+v", user)
	if user.ID != 1 || user.Name != "Alice" {
		t.Errorf("Expected user with ID 1 and name Alice, got ID %d and name %s", user.ID, user.Name)
	}
}

// Test for creating a new user
func TestCreateUser(t *testing.T) {
	setupHandler(t)

	newUser := models.User{Name: "John", Email: "john@example.com"}
	jsonData, _ := json.Marshal(newUser)

	t.Logf("Sending POST request to create a new user: %+v", newUser)
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(userHandler.CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	} else {
		t.Logf("POST /users returned status code: %v", rr.Code)
	}

	t.Logf("Response body: %s", rr.Body.String())
}

// Test for updating an existing user
func TestUpdateUser(t *testing.T) {
	setupHandler(t)

	updatedUser := models.User{ID: 1, Name: "Alice Updated", Email: "alice.updated@example.com"}
	jsonData, _ := json.Marshal(updatedUser)

	t.Logf("Sending PUT request to update user with ID 1 to: %+v", updatedUser)
	req := httptest.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.UpdateUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		t.Logf("PUT /users/1 returned status code: %v", rr.Code)
	}

	t.Logf("Response body: %s", rr.Body.String())
}

// Test for deleting a user
func TestDeleteUser(t *testing.T) {
	setupHandler(t)

	t.Log("Sending DELETE request to remove user with ID 1...")
	req := httptest.NewRequest("DELETE", "/users/1", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.DeleteUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	} else {
		t.Logf("DELETE /users/1 returned status code: %v", rr.Code)
	}

	t.Logf("Response body: %s", rr.Body.String())
}
