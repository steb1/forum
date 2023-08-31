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
	Username string
	ImageURL string
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
	_, err = rr.db.Exec("INSERT INTO request (id, authorID, time, username, imageurl) VALUES (?, ?, ?, ?, ?)",
		request.ID, request.AuthorID, request.Time, request.Username, request.ImageURL)
	return err
}

// Get a report by ID from the database
func (rr *RequestRepository) GetRequestByID(requestID string) (*Request, error) {
	var request Request
	row := rr.db.QueryRow("SELECT id, authorID, time, username, imageurl FROM request WHERE id = ?", requestID)
	err := row.Scan(&request.ID, &request.AuthorID, &request.Time, &request.Username, &request.ImageURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Report not found
		}
		return nil, err
	}
	return &request, nil
}

// Select All Requestss
func (rr *RequestRepository) GetAllRequest() ([]Request, error) {
	var request []Request
	row, err := rr.db.Query("SELECT * FROM request")
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		var ID string
		var AuthorID string
		var Time string
		var Username string
		var ImageURL string

		err = row.Scan(&ID, &AuthorID, &Time, &Username, &ImageURL)
		if err != nil {
			log.Fatal(err)
		}

		var tab = Request{
			ID:       ID,
			AuthorID: AuthorID,
			Time:     Time,
			Username: Username,
			ImageURL: ImageURL,
		}

		request = append(request, tab)
	}
	return request, nil
}

// Delete a report from the database
func (rr *RequestRepository) DeleteRequest(requestID string) error {
	_, err := rr.db.Exec("DELETE FROM request WHERE id = ?", requestID)
	return err
}
