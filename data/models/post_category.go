package models

import (
	"database/sql"
	"log"

	uuid "github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type PostCategory struct {
	ID         string
	CategoryID string
	PostID     string
}

type PostCategoryRepository struct {
	db *sql.DB
}

func NewPostCategoryRepository(db *sql.DB) *PostCategoryRepository {
	return &PostCategoryRepository{
		db: db,
	}
}

// Create a new post-category relationship in the database
func (pcr *PostCategoryRepository) CreatePostCategory(categoryID, postID string) error {
	ID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("❌ Failed to generate UUID: %v", err)
	}
	postCategory := PostCategory{
		ID:         ID.String(),
		CategoryID: categoryID,
		PostID:     postID,
	}
	_, err = pcr.db.Exec("INSERT INTO post_category (id, categoryID, postID) VALUES (?, ?, ?)",
		postCategory.ID, postCategory.CategoryID, postCategory.PostID)
	return err
}

// Get categories of a post from the database
func (pcr *PostCategoryRepository) GetCategoriesOfPost(postID string) ([]Category, error) {
	rows, err := pcr.db.Query(`
		SELECT c.id, c.name, c.createDate, c.modifiedDate
		FROM category c
		INNER JOIN post_category pc ON c.id = pc.categoryID
		WHERE pc.postID = ?
	`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.CreateDate, &category.ModifiedDate)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// Get posts of a category from the database
func (pcr *PostCategoryRepository) GetPostsOfCategory(categoryID string) ([]Post, error) {
	rows, err := pcr.db.Query(`
		SELECT p.id, p.title, p.description, p.imageURL, p.authorID, p.isEdited, p.createDate, p.modifiedDate
		FROM post p
		INNER JOIN post_category pc ON p.id = pc.postID
		WHERE pc.categoryID = ?
	`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID, &post.Title, &post.Description, &post.ImageURL,
			&post.AuthorID, &post.IsEdited, &post.CreateDate, &post.ModifiedDate,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
