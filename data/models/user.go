package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID                  string
	Username            string
	Email               string
	Password            string
	AvatarURL           string
	Role                ROLE
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
	_, err := ur.db.Exec("INSERT INTO user (id, username, email, password, avatarURL, role, token, tokenExpirationDate) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.ID, user.Username, user.Email, user.Password, user.AvatarURL, user.Role, user.Token, user.TokenExpirationDate)
	return err
}

// Get a user by ID from the database
func (ur *UserRepository) GetUserByID(userID string) (*User, error) {
	var user User
	row := ur.db.QueryRow("SELECT id, username, email, password, avatarURL, type, token, tokenExpirationDate FROM user WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.AvatarURL, &user.Role, &user.Token, &user.TokenExpirationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

// Get a user by email from the database
func (ur *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	row := ur.db.QueryRow("SELECT id, username, email, password, avatarURL, type, token, tokenExpirationDate FROM user WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.AvatarURL, &user.Role, &user.Token, &user.TokenExpirationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

// Select All users
func (ur *UserRepository) SelectAllUsers() ([]User, error) {
	var user []User
	row, err := ur.db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		var ID string
		var Email string
		var Username string
		var Password string
		var AvatarUrl string
		var Type ROLE
		var Token string
		var TokenExpirationDate string

		err = row.Scan(&ID, &Email, &Username, &Password, &AvatarUrl, &Type, &Token, &TokenExpirationDate)

		if err != nil {
			log.Fatal(err)
		}

		var tab = User{
			ID:                  ID,
			Email:               Email,
			Username:            Username,
			Password:            Password,
			AvatarURL:           AvatarUrl,
			Role:                Type,
			Token:               Token,
			TokenExpirationDate: TokenExpirationDate,
		}

		user = append(user, tab)
	}
	return user, nil
}

// Update a user in the database
func (ur *UserRepository) UpdateUser(user *User) error {
	_, err := ur.db.Exec("UPDATE user SET username = ?, email = ?, password = ?, avatarURL = ?, type = ?, token = ?, tokenExpirationDate = ? WHERE id = ?",
		user.Username, user.Email, user.Password, user.AvatarURL, user.Role, user.Token, user.TokenExpirationDate, user.ID)
	return err
}

// Delete a user from the database
func (ur *UserRepository) DeleteUser(userID string) error {
	_, err := ur.db.Exec("DELETE FROM user WHERE id = ?", userID)
	return err
}

// Check if user exists
func (ur *UserRepository) IsExisted(email string) (*User, bool) {
	var user User
	row := ur.db.QueryRow("SELECT password FROM user WHERE email = ?", email)
	err := row.Scan(&user.Password)
	if err != nil {
		log.Println("❌ ", err)
		if err == sql.ErrNoRows {
			return nil, false
		}
		return nil, false
	}
	return &user, true
}
