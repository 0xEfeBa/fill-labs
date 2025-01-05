package services

import (
	"myapp/models"
	"myapp/repositories"
)

// User Service: Business logic layer for user operations
// Handles communication between handlers and repository

type UserService struct {
	userRepo *repositories.UserRepository
}

// Create new service instance with repository dependency
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Get all users from repository
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

// Find specific user by their ID
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

// Create new user in the system
func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

// Update existing user information
func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}

// Remove user from the system
func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}
