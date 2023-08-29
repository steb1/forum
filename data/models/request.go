package models

import (
	"database/sql"
	"log"

	uuid "github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Request struct {
	ID       string
	AuthorID string
	Time     string
}

type RequestRepository struct {
	db *sql.DB
}

func NewRequestRepository(db *sql.DB) *RequestRepository {
	return &RequestRepository{
		db: db,
	}
}

// Create a new report in the database
func (rr *RequestRepository) CreateRequest(request *Request) error {
	ID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("❌ Failed to generate UUID: %v", err)
	}
	request.ID = ID.String()
	_, err = rr.db.Exec("INSERT INTO request (id, authorID, time) VALUES (?, ?, ?)",
		request.ID, request.AuthorID, request.Time)
	return err
}

// Get a report by ID from the database
func (rr *RequestRepository) GetRequestByID(requestID string) (*Request, error) {
	var request Request
	row := rr.db.QueryRow("SELECT id, authorID, time FROM request WHERE id = ?", requestID)
	err := row.Scan(&request.ID, &request.AuthorID, &request.Time)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Report not found
		}
		return nil, err
	}
	return &request, nil
}

// Update a report in the database
func (rr *RequestRepository) UpdateRequest(request *Request) error {
	_, err := rr.db.Exec("UPDATE request SET authorID = ?, time = ? WHERE id = ?",
		request.AuthorID, request.Time, request.ID)
	return err
}

// Delete a report from the database
func (rr *RequestRepository) DeleteRequest(requestID string) error {
	_, err := rr.db.Exec("DELETE FROM request WHERE id = ?", requestID)
	return err
}
