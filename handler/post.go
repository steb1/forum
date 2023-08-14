package handler

import (
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PostPageData struct {
	IsLoggedIn    bool
	RandomUsers   []models.User
	CurrentUser   models.User
	Post          []models.PostItem
	NumberOfPosts int
	TopUsers      []models.TopUser
}

func Post(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/post", http.MethodPost) {
		var categories = make(map[string]models.Category)
		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		isEdited := false
		creationDate := time.Now().Format("2006-01-02 15:04:05")
		modifDate := time.Now().Format("2006-01-02 15:04:05")
		title := req.FormValue("title")
		title = lib.Slugify(title)
		description := req.FormValue("description")
		categorie := req.FormValue("categorie")

		u := uuid.New()

		//--------------------------------------
		imageUrl := lib.UploadImage(req)
		authorID := models.GetUserFromSession(req).ID
		//--------------------------------------
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

		models.PostRepo.CreatePost(&postStruct)

		for i := 0; i < len(categories); i++ {
			catStruct := models.Category{
				ID:           categories[tabUUID[i]].ID,
				Name:         categories[tabUUID[i]].Name,
				CreateDate:   categories[tabUUID[i]].CreateDate,
				ModifiedDate: categories[tabUUID[i]].ModifiedDate,
			}
			models.CategoryRepo.CreateCategory(&catStruct)
		}
	}
}

func AllPosts(res http.ResponseWriter, req *http.Request) {
	PostComments := []models.Comment{}
	if lib.ValidateRequest(req, res, "/posts", http.MethodGet) {
		basePath := "base"
		pagePath := "post"
		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "posts" {
			post, err := models.PostRepo.GetPostByTitle(pathPart[2])
			if err != nil {
				fmt.Println("error DB")
				return
			}
			comments, err := models.CommentRepo.GetAllComments("15")
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
			lib.RenderPage(basePath, pagePath, nil, res)
			log.Println("✅ Home page get with success")
		} else {
			http.NotFound(res, req)
		}
	}
}

func Comment(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/comment", http.MethodPost) {
		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		text := req.FormValue("text")

		u := uuid.New()
		creationDate := time.Now().Format("2006-01-02")
		modifDate := time.Now().Format("2006-01-02")
		//--------------------------------------
		authorID := models.GetUserFromSession(req).ID
		parentID := "chjchjchjcxjchjc"
		postID := "709433aa-9fe4-4935-b1d6-48b50e24eb20"
		//--------------------------------------
		commentStruct := models.Comment{
			ID:           u.String(),
			Text:         text,
			AuthorID:     authorID,
			PostID:       postID,
			ParentID:     parentID,
			CreateDate:   creationDate,
			ModifiedDate: modifDate}

		models.CommentRepo.CreateComment(&commentStruct)
	}
}
