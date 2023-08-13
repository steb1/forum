package models

import (
	"database/sql"
	"forum/lib"
	"log"
	"strconv"
	"strings"

	uuid "github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type PostItem struct {
	ID                string
	Title             string
	AuthorName        string
	ImageURL          string
	LastEditionDate   string
	NumberOfComments  int
	ListOfCommentator []string
}

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

func (pr *PostRepository) GetUserPost(userId, userName string) ([]PostItem, error) {
	var posts []*Post
	var numberComments []int

	rows, err := pr.db.Query(`
	SELECT p.id AS id, title, description, imageURL, p.authorID AS authorID, isEdited, p.createDate AS createDate, p.modifiedDate AS modifiedDate, COUNT(*) AS numberComment FROM post p
LEFT JOIN comment c ON c.postID = p.ID
WHERE p.authorID = ?
GROUP BY p.ID ;
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		var nbComment int
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImageURL, &post.AuthorID, &post.IsEdited, &post.CreateDate, &post.ModifiedDate, &nbComment)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
		numberComments = append(numberComments, nbComment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	tabPostItem := []PostItem{}

	for i := 0; i < len(posts); i++ {
		lastModificationDate := strings.ReplaceAll(posts[i].ModifiedDate, "T", " ")
		lastModificationDate = strings.ReplaceAll(lastModificationDate, "Z", "")
		urlImage := strings.ReplaceAll(posts[i].ImageURL, "jpg", "jpg")
		postItem := PostItem{
			ID:                posts[i].ID,
			Title:             posts[i].Title,
			AuthorName:        userName,
			ImageURL:          urlImage,
			LastEditionDate:   lib.TimeSinceCreation(lastModificationDate),
			NumberOfComments:  numberComments[i],
			ListOfCommentator: []string{},
		}
		tabPostItem = append(tabPostItem, postItem)
	}

	return tabPostItem, nil
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

// Get all posts from database
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

// Get all posts from database
func (pr *PostRepository) GetAllPostsItems(more string) ([]PostItem, error) {
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
	tabPostItem := []PostItem{}

	for i := 0; i < len(posts); i++ {
		tabUser, _ := UserRepo.SelectAllUsers()
		tabAllComments, _ := CommentRepo.GetAllComments("15")
		user := ""

		tabComments := []string{}
		for k := 0; k < len(tabAllComments); k++ {
			if posts[i].ID == tabAllComments[k].PostID {
				tabComments = append(tabComments, posts[i].ImageURL)
			}
		}
		for j := 0; j < len(tabUser); j++ {
			if posts[i].AuthorID == tabUser[j].ID {
				user = tabUser[j].Username
				break
			}
		}

		TopUser, _ := UserRepo.SelectAllUsersBypost(posts[i].ID)
		tabTopUser := []string{}
		cpt := 0
		for l := 0; l < len(TopUser); l++ {
			if cpt < 3 {
				tabTopUser = append(tabTopUser, TopUser[l].AvatarURL)
			}
			cpt++
		}

		lastmodif := strings.ReplaceAll(posts[i].ModifiedDate, "T", " ")
		lastmodif = strings.ReplaceAll(lastmodif, "Z", "")
		urlImage := strings.ReplaceAll(posts[i].ImageURL, "jpg", "jpg")
		PostItemi := PostItem{
			ID:                posts[i].ID,
			Title:             posts[i].Title,
			AuthorName:        user,
			ImageURL:          urlImage,
			LastEditionDate:   lib.TimeSinceCreation(lastmodif),
			NumberOfComments:  len(tabComments),
			ListOfCommentator: tabTopUser}
		tabPostItem = append(tabPostItem, PostItemi)
	}

	return tabPostItem, nil
}

// Get the number of posts in the database
func (pr *PostRepository) GetNumberOfPosts() int {
	var numberOfPosts int

	row := pr.db.QueryRow("SELECT COUNT(*) FROM post")
	err := row.Scan(&numberOfPosts)
	if err != nil {
		return 0
	}
	return numberOfPosts
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
