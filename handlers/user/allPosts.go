package user

import (
	"database/sql"
	"fmt"
	"forum/data/models"
	"log"
	"net/http"
)

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
