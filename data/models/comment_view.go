package models

import (
	"database/sql"
	"log"

	uuid "github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Comment_View struct {
	ID        string
	Rate      RATE
	AuthorID  string
	CommentID string
}

type CommentViewRepository struct {
	db *sql.DB
}

func NewCommentViewRepository(db *sql.DB) *CommentViewRepository {
	return &CommentViewRepository{
		db: db,
	}
}

// Create a new view in the database
func (vr *CommentViewRepository) CreateCommentView(Comment_View *Comment_View) error {
	ID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("❌ Failed to generate UUID: %v", err)
	}
	Comment_View.ID = ID.String()
	_, err = vr.db.Exec("INSERT INTO comment_view (id, rate, authorID, commentID) VALUES (?, ?, ?, ?)",
		Comment_View.ID, Comment_View.Rate, Comment_View.AuthorID, Comment_View.CommentID)
	return err
}

// Get a view by ID from the database
func (vr *CommentViewRepository) GetCommentViewByID(Comment_ViewID string) (*Comment_View, error) {
	var Comment_View Comment_View
	row := vr.db.QueryRow("SELECT id, rate, authorID, commentID FROM comment_view WHERE id = ?", Comment_ViewID)
	err := row.Scan(&Comment_View.ID, &Comment_View.Rate, &Comment_View.AuthorID, &Comment_View.CommentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Comment_View not found
		}
		return nil, err
	}
	return &Comment_View, nil
}

// Get a number of dislike by post
func (vr *CommentViewRepository) GetLikesByComment(commentID string) (int, error) {
	var nbrLike int
	row := vr.db.QueryRow("SELECT COUNT(*) FROM comment_view WHERE commentID = ? AND rate=1", commentID)
	err := row.Scan(&nbrLike)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // View not found
		}
		return 0, err
	}
	return nbrLike, nil
}

func (vr *CommentViewRepository) GetDislikesByComment(commentID string) (int, error) {
	var nbrLike int
	row := vr.db.QueryRow("SELECT COUNT(*) FROM comment_view WHERE commentID = ? AND rate=2", commentID)
	err := row.Scan(&nbrLike)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // View not found
		}
		return 0, err
	}
	return nbrLike, nil
}

// Get a view by ID from the database
func (vr *CommentViewRepository) GetViewByAuthorIDandCommentID(authorID string, commentID string) (*Comment_View, error) {
	var comment_view Comment_View
	row := vr.db.QueryRow("SELECT id, rate, authorID, commentID FROM comment_view WHERE authorid = ? AND commentid = ?", authorID, commentID)
	err := row.Scan(&comment_view.ID, &comment_view.Rate, &comment_view.AuthorID, &comment_view.CommentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // comment_view not found
		}
		return nil, err
	}
	return &comment_view, nil
}

// Update a view in the database
func (vr *CommentViewRepository) UpdateView(comment_view *Comment_View) error {
	_, err := vr.db.Exec("UPDATE comment_view SET rate = ?, authorID = ?, commentID = ? WHERE id = ?",
		comment_view.Rate, comment_view.AuthorID, comment_view.CommentID, comment_view.ID)
	return err
}

// Delete a view from the database
func (vr *CommentViewRepository) DeleteCommentView(comment_viewID string) error {
	_, err := vr.db.Exec("DELETE FROM comment_view WHERE id = ?", comment_viewID)
	return err
}
