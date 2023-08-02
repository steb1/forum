package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type View struct {
	ID           string
	IsBookmarked bool
	Rate         string
	AuthorID     string
	PostID       string
}

type ViewRepository struct {
	db *sql.DB
}

func NewViewRepository(db *sql.DB) *ViewRepository {
	return &ViewRepository{
		db: db,
	}
}

// Create a new view in the database
func (vr *ViewRepository) CreateView(view *View) error {
	_, err := vr.db.Exec("INSERT INTO view (id, isBookmarked, rate, authorID, postID) VALUES (?, ?, ?, ?, ?)",
		view.ID, view.IsBookmarked, view.Rate, view.AuthorID, view.PostID)
	return err
}

// Get a view by ID from the database
func (vr *ViewRepository) GetViewByID(viewID string) (*View, error) {
	var view View
	row := vr.db.QueryRow("SELECT id, isBookmarked, rate, authorID, postID FROM view WHERE id = ?", viewID)
	err := row.Scan(&view.ID, &view.IsBookmarked, &view.Rate, &view.AuthorID, &view.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // View not found
		}
		return nil, err
	}
	return &view, nil
}

// Update a view in the database
func (vr *ViewRepository) UpdateView(view *View) error {
	_, err := vr.db.Exec("UPDATE view SET isBookmarked = ?, rate = ?, authorID = ?, postID = ? WHERE id = ?",
		view.IsBookmarked, view.Rate, view.AuthorID, view.PostID, view.ID)
	return err
}

// Delete a view from the database
func (vr *ViewRepository) DeleteView(viewID string) error {
	_, err := vr.db.Exec("DELETE FROM view WHERE id = ?", viewID)
	return err
}
