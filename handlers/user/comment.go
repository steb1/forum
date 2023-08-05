package user

import (
	"database/sql"
	"fmt"
	"forum/data/models"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

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
		postID := "djnnjdd"
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
