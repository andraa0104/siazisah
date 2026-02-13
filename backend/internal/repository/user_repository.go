package repository

import (
	"database/sql"
	"github.com/yourusername/siazisah/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (username, email, password, full_name, role, masjid_id, is_active) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := r.DB.Exec(query, user.Username, user.Email, user.Password, user.FullName, user.Role, user.MasjidID, user.IsActive)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password, full_name, role, masjid_id, is_active, created_at, updated_at 
			  FROM users WHERE username = ?`
	err := r.DB.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.FullName,
		&user.Role, &user.MasjidID, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password, full_name, role, masjid_id, is_active, created_at, updated_at 
			  FROM users WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.FullName,
		&user.Role, &user.MasjidID, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	query := `SELECT id, username, email, full_name, role, masjid_id, is_active, created_at, updated_at 
			  FROM users ORDER BY created_at DESC`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.Role, &user.MasjidID, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetByMasjidID(masjidID int) ([]models.User, error) {
	query := `SELECT id, username, email, full_name, role, masjid_id, is_active, created_at, updated_at 
			  FROM users WHERE masjid_id = ? ORDER BY created_at DESC`
	rows, err := r.DB.Query(query, masjidID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.Role, &user.MasjidID, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Update(user *models.User) error {
	query := `UPDATE users SET username=?, email=?, full_name=?, role=?, masjid_id=?, is_active=? WHERE id=?`
	_, err := r.DB.Exec(query, user.Username, user.Email, user.FullName, user.Role, user.MasjidID, user.IsActive, user.ID)
	return err
}

func (r *UserRepository) UpdatePassword(id int, hashedPassword string) error {
	query := `UPDATE users SET password=? WHERE id=?`
	_, err := r.DB.Exec(query, hashedPassword, id)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id=?`
	_, err := r.DB.Exec(query, id)
	return err
}
