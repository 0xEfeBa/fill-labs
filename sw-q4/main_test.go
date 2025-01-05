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

func init() {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
}

func setupTestDatabase(t *testing.T) {
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		t.Fatal(err)
	}

	createTableSQL := `CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		t.Fatal(err)
	}

	insertUserSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
	_, err = db.Exec(insertUserSQL, "Alice", "alice@example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func setupHandler(t *testing.T) {
	setupTestDatabase(t)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler = handlers.NewUserHandler(userService)
}

// Test GET /users
func TestGetAllUsers(t *testing.T) {
	setupHandler(t)

	req := httptest.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(userHandler.GetAllUsers)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var users []models.User
	if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
		t.Fatal(err)
	}

	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}
}

// Test GET /users/{id} with existing user
func TestGetUserByID(t *testing.T) {
	setupHandler(t)

	req := httptest.NewRequest("GET", "/users/1", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.GetUserByID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var user models.User
	if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
		t.Fatal(err)
	}

	if user.ID != 1 || user.Name != "Alice" {
		t.Errorf("Expected user with ID 1 and name Alice, got ID %d and name %s", user.ID, user.Name)
	}
}

// Test GET /users/{id} with non-existing user
func TestGetUserByID_NotFound(t *testing.T) {
	setupHandler(t)

	req := httptest.NewRequest("GET", "/users/999", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.GetUserByID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

// Test POST /users
func TestCreateUser(t *testing.T) {
	setupHandler(t)

	newUser := models.User{Name: "John", Email: "john@example.com"}
	jsonData, _ := json.Marshal(newUser)

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(userHandler.CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

// Test PUT /users/{id} with existing user
func TestUpdateUser(t *testing.T) {
	setupHandler(t)

	updatedUser := models.User{ID: 1, Name: "Alice Updated", Email: "alice.updated@example.com"}
	jsonData, _ := json.Marshal(updatedUser)

	req := httptest.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.UpdateUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Test PUT /users/{id} with non-existing user
func TestUpdateUser_NotFound(t *testing.T) {
	setupHandler(t)

	updatedUser := models.User{Name: "Updated Name", Email: "updated@example.com"}
	jsonData, _ := json.Marshal(updatedUser)

	req := httptest.NewRequest("PUT", "/users/999", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.UpdateUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

// Test DELETE /users/{id} with existing user
func TestDeleteUser(t *testing.T) {
	setupHandler(t)

	req := httptest.NewRequest("DELETE", "/users/1", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.DeleteUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

// Test DELETE /users/{id} with non-existing user
func TestDeleteUser_NotFound(t *testing.T) {
	setupHandler(t)

	req := httptest.NewRequest("DELETE", "/users/999", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.DeleteUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
