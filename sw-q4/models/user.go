package models

// User Model: Defines the structure for user data
// Provides JSON mapping for API communication

type User struct {
	ID    int    `json:"id"`    // Unique identifier for the user
	Name  string `json:"name"`  // User's full name
	Email string `json:"email"` // User's email address
}
