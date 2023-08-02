package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Comment struct {
	ID           string
	Text         string
	AuthorID     string
	PostID       string
	ParentID     string
	CreateDate   string
	ModifiedDate string
}

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

// Create a new comment in the database
func (cr *CommentRepository) CreateComment(comment *Comment) error {
	_, err := cr.db.Exec("INSERT INTO comment (id, text, authorID, postID, parentID, createDate, modifiedDate) VALUES (?, ?, ?, ?, ?, ?, ?)",
		comment.ID, comment.Text, comment.AuthorID, comment.PostID, comment.ParentID, comment.CreateDate, comment.ModifiedDate)
	return err
}

// Get a comment by ID from the database
func (cr *CommentRepository) GetCommentByID(commentID string) (*Comment, error) {
	var comment Comment
	row := cr.db.QueryRow("SELECT id, text, authorID, postID, parentID, createDate, modifiedDate FROM comment WHERE id = ?", commentID)
	err := row.Scan(&comment.ID, &comment.Text, &comment.AuthorID, &comment.PostID, &comment.ParentID, &comment.CreateDate, &comment.ModifiedDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Comment not found
		}
		return nil, err
	}
	return &comment, nil
}

// Update a comment in the database
func (cr *CommentRepository) UpdateComment(comment *Comment) error {
	_, err := cr.db.Exec("UPDATE comment SET text = ?, authorID = ?, postID = ?, parentID = ?, createDate = ?, modifiedDate = ? WHERE id = ?",
		comment.Text, comment.AuthorID, comment.PostID, comment.ParentID, comment.CreateDate, comment.ModifiedDate, comment.ID)
	return err
}

// Delete a comment from the database
func (cr *CommentRepository) DeleteComment(commentID string) error {
	_, err := cr.db.Exec("DELETE FROM comment WHERE id = ?", commentID)
	return err
}
