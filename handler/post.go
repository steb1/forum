package handler

import (
	"database/sql"
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Post(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "./data/sql/forum.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	var categories = make(map[string]models.Category)

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		isEdited := false
		creationDate := time.Now().Format("2006-01-02")
		modifDate := time.Now().Format("2006-01-02")
		title := r.FormValue("title")
		title = lib.Slugify(title)
		description := r.FormValue("description")
		categorie := r.FormValue("categorie")

		u := uuid.New()

		imageUrl := "chemin de l'image..."
		authorID := "ejn3b3h3h3"
		tabcat := strings.Split(categorie, "#")

		tabUUID := []string{}
		for i := 1; i < len(tabcat); i++ {
			c := uuid.New()
			tabUUID = append(tabUUID, c.String())
			categories[c.String()] = models.Category{ID: c.String(), Name: strings.TrimSpace(tabcat[i]), CreateDate: creationDate, ModifiedDate: modifDate}
		}

		postStruct := models.Post{
			ID:           u.String(),
			Title:        title,
			Description:  description,
			ImageURL:     imageUrl,
			AuthorID:     authorID,
			IsEdited:     isEdited,
			CreateDate:   creationDate,
			ModifiedDate: modifDate}

		models.NewPostRepository(db).CreatePost(&postStruct)

		for i := 0; i < len(categories); i++ {
			catStruct := models.Category{
				ID:           categories[tabUUID[i]].ID,
				Name:         categories[tabUUID[i]].Name,
				CreateDate:   categories[tabUUID[i]].CreateDate,
				ModifiedDate: categories[tabUUID[i]].ModifiedDate,
			}
			models.NewCategoryRepository(db).CreateCategory(&catStruct)
		}

	} else {
		fmt.Println("Method not allowed")
	}
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./data/sql/forum.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	PostComments := []models.Comment{}
	if r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		path := r.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "posts" {
			post, err := models.NewPostRepository(db).GetPostByTitle(pathPart[2])
			comments, err := models.NewCommentRepository(db).GetAllComments("15")
			if err != nil {
				fmt.Println("error DB")
				return
			}

			for j := 0; j < len(comments); j++ {
				if post.ID == comments[j].PostID {
					PostComments = append(PostComments, *comments[j])
				}
			}

			fmt.Println(PostComments)
		} else {
			http.NotFound(w, r)
		}

	} else {
		fmt.Println("Method not allowed")
	}
}

func Comment(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./data/sql/forum.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		text := r.FormValue("text")

		u := uuid.New()
		creationDate := time.Now().Format("2006-01-02")
		modifDate := time.Now().Format("2006-01-02")
		authorID := "cdjndjd"
		parentID := "fdinjff"
		postID := "3356e5b9-57c9-4c1f-b67c-7e485f66eab9"
		commentStruct := models.Comment{
			ID:           u.String(),
			Text:         text,
			AuthorID:     authorID,
			PostID:       postID,
			ParentID:     parentID,
			CreateDate:   creationDate,
			ModifiedDate: modifDate}

		models.NewCommentRepository(db).CreateComment(&commentStruct)

	} else {
		fmt.Println("Method not allowed")
	}
}
