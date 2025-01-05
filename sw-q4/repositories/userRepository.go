package repositories

import (
	"database/sql"
	"myapp/models"
)

// User Repository: Handles database operations for user data
// Implements CRUD operations using SQL

// UserRepository, veritabanı işlemlerini yönetir.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository, yeni bir UserRepository örneği döndürür.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Retrieve all users from database
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	// SQL sorgusu
	query := "SELECT id, name, email FROM users"

	// Sorguyu çalıştır
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Kullanıcıları saklamak için bir dilim oluştur
	var users []models.User

	// Satırları dolaş ve kullanıcıları dilime ekle
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Hata kontrolü
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Find user by ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	// SQL sorgusu
	query := "SELECT id, name, email FROM users WHERE id = ?"

	// Sorguyu çalıştır
	row := r.db.QueryRow(query, id)

	// Kullanıcıyı saklamak için bir değişken oluştur
	var user models.User

	// Satırı tarayıp kullanıcıya ata
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Kullanıcı bulunamadı
		}
		return nil, err
	}

	return &user, nil
}

// Insert new user record
func (r *UserRepository) CreateUser(user *models.User) error {
	// SQL sorgusu
	query := "INSERT INTO users (name, email) VALUES (?, ?)"

	// Sorguyu çalıştır
	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}

	// Eklenen kullanıcının ID'sini al
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Kullanıcıya ID'yi ata
	user.ID = int(id)

	return nil
}

// Update existing user record
func (r *UserRepository) UpdateUser(user *models.User) error {
	// SQL sorgusu
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"

	// Sorguyu çalıştır
	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// Remove user record
func (r *UserRepository) DeleteUser(id int) error {
	// SQL sorgusu
	query := "DELETE FROM users WHERE id = ?"

	// Sorguyu çalıştır
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
