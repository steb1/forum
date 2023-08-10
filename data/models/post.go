package models

import (
	"database/sql"
	"log"
	"strconv"

	uuid "github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	ID           string
	Title        string
	Description  string
	ImageURL     string
	AuthorID     string
	IsEdited     bool
	CreateDate   string
	ModifiedDate string
}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

// Create a new post in the database
func (pr *PostRepository) CreatePost(post *Post) error {
	ID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("❌ Failed to generate UUID: %v", err)
	}
	post.ID = ID.String()
	_, err = pr.db.Exec("INSERT INTO post (id, title, description, imageURL, authorID, isEdited, createDate, modifiedDate) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		post.ID, post.Title, post.Description, post.ImageURL, post.AuthorID, post.IsEdited, post.CreateDate, post.ModifiedDate)
	return err
}

// Get a post by ID from the database
func (pr *PostRepository) GetPostByID(postID string) (*Post, error) {
	var post Post
	row := pr.db.QueryRow("SELECT id, title, description, imageURL, authorID, isEdited, createDate, modifiedDate FROM post WHERE id = ?", postID)
	err := row.Scan(&post.ID, &post.Title, &post.Description, &post.ImageURL, &post.AuthorID, &post.IsEdited, &post.CreateDate, &post.ModifiedDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		return nil, err
	}
	return &post, nil
}

// Get a post by TITLE from the database
func (pr *PostRepository) GetPostByTitle(title string) (*Post, error) {
	var post Post
	row := pr.db.QueryRow("SELECT id, title, description, imageURL, authorID, isEdited, createDate, modifiedDate FROM post WHERE title = ?", title)
	err := row.Scan(&post.ID, &post.Title, &post.Description, &post.ImageURL, &post.AuthorID, &post.IsEdited, &post.CreateDate, &post.ModifiedDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		return nil, err
	}
	return &post, nil
}

// Get all posts from databse
func (pr *PostRepository) GetAllPosts(more string) ([]*Post, error) {
	morePost, err := strconv.Atoi(more)
	if err != nil {
		return nil, err
	}
	var posts []*Post

	rows, err := pr.db.Query("SELECT id, title, description, imageURL, authorID, isEdited, createDate, modifiedDate FROM post LIMIT ?", morePost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImageURL, &post.AuthorID, &post.IsEdited, &post.CreateDate, &post.ModifiedDate)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// Update a post in the database
func (pr *PostRepository) UpdatePost(post *Post) error {
	_, err := pr.db.Exec("UPDATE post SET title = ?, description = ?, imageURL = ?, authorID = ?, isEdited = ?, createDate = ?, modifiedDate = ? WHERE id = ?",
		post.Title, post.Description, post.ImageURL, post.AuthorID, post.IsEdited, post.CreateDate, post.ModifiedDate, post.ID)
	return err
}

// Delete a post from the database
func (pr *PostRepository) DeletePost(postID string) error {
	_, err := pr.db.Exec("DELETE FROM post WHERE id = ?", postID)
	return err
}
