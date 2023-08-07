package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"forum/data/models"

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
		description := r.FormValue("description")
		categorie := r.FormValue("categorie")
		// imageUrl := r.FormValue("image")
		fmt.Printf("Title : %s\n", title)
		fmt.Printf("description : %s\n", description)
		fmt.Printf("isEdited : %v\n", isEdited)
		fmt.Printf("creationDate : %s\n", creationDate)
		fmt.Printf("modifDate : %s\n", modifDate)
		fmt.Printf("categories : %s\n", categorie)
		// fmt.Printf("image : %s\n", imageUrl)

		u := uuid.New()
		fmt.Print("UUID = ")
		fmt.Println(u.String())
		//authorID : to do with front TRICK
		imageUrl := "chemin de l'image..."
		fmt.Println(imageUrl)
		authorID := "ejn3b3h3h3"
		fmt.Println(authorID)
		tabcat := strings.Split(categorie, "#")
		fmt.Println(tabcat)

		tabUUID := []string{}
		for i := 1; i < len(tabcat); i++ {
			c := uuid.New()
			tabUUID = append(tabUUID, c.String())
			categories[c.String()] = models.Category{ID: c.String(), Name: strings.TrimSpace(tabcat[i]), CreateDate: creationDate, ModifiedDate: modifDate}
		}
		fmt.Print("categories : ")
		fmt.Println(categories)
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
		// fmt.Println(categories["0"].Name)
		for i := 0; i < len(categories); i++ {
			catStruct := models.Category{
				ID:           categories[tabUUID[i]].ID,
				Name:         categories[tabUUID[i]].Name,
				CreateDate:   categories[tabUUID[i]].CreateDate,
				ModifiedDate: categories[tabUUID[i]].ModifiedDate,
			}
			fmt.Println("---------------------")
			fmt.Println(catStruct)
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

	var PostsAndComments = make(map[models.Post]models.Comment)
	if r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Structposts := []models.Post{}
		posts, err := models.NewPostRepository(db).GetAllPosts()
		comments, err := models.NewCommentRepository(db).GetAllComments()
		if err != nil {
			fmt.Println("error DB")
			return
		}
		for i := 0; i < len(posts); i++ {
			for j := 0; j < len(comments); j++ {
				if posts[i].ID == comments[j].PostID {
					PostsAndComments[*posts[i]] = *comments[j]
				}
			}
		}

		fmt.Println(PostsAndComments)

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
		fmt.Printf("Comment : %s\n", text)

		u := uuid.New()
		fmt.Print("UUID = ")
		fmt.Println(u.String())
		creationDate := time.Now().Format("2006-01-02")
		modifDate := time.Now().Format("2006-01-02")
		fmt.Printf("creationDate : %s\n", creationDate)
		fmt.Printf("modifDate : %s\n", modifDate)
		// authorID, parentID, postID : to do with front TRICK
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
		//task to do : format date -> only day-month-year

	} else {
		fmt.Println("Method not allowed")
	}
}
