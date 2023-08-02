package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID                  string
	Username            string
	Email               string
	Password            string
	AvatarURL           string
	Type                ROLE
	Token               string
	TokenExpirationDate string
}

type ROLE int

const (
	RoleAdmin     ROLE = 0
	RoleModerator ROLE = 1
	RoleUser      ROLE = 2
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create a new user in the database
func (ur *UserRepository) CreateUser(user *User) error {
	_, err := ur.db.Exec("INSERT INTO user (id, username, email, password, avatarURL, type, token, tokenExpirationDate) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.ID, user.Username, user.Email, user.Password, user.AvatarURL, user.Type, user.Token, user.TokenExpirationDate)
	return err
}

// Get a user by ID from the database
func (ur *UserRepository) GetUserByID(userID string) (*User, error) {
	var user User
	row := ur.db.QueryRow("SELECT id, username, email, password, avatarURL, type, token, tokenExpirationDate FROM user WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.AvatarURL, &user.Type, &user.Token, &user.TokenExpirationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

// Update a user in the database
func (ur *UserRepository) UpdateUser(user *User) error {
	_, err := ur.db.Exec("UPDATE user SET username = ?, email = ?, password = ?, avatarURL = ?, type = ?, token = ?, tokenExpirationDate = ? WHERE id = ?",
		user.Username, user.Email, user.Password, user.AvatarURL, user.Type, user.Token, user.TokenExpirationDate, user.ID)
	return err
}

// Delete a user from the database
func (ur *UserRepository) DeleteUser(userID string) error {
	_, err := ur.db.Exec("DELETE FROM user WHERE id = ?", userID)
	return err
}
