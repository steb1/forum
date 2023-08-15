package handler

import (
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
	"time"
)

type PostPageData struct {
	IsLoggedIn bool
	Post       models.Post
	Comments   []models.Comment
	UserPoster *models.User
	NbrComment int
}

func Post(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/post", http.MethodPost) {
		isSessionOpen := models.ValidSession(req)
		if isSessionOpen {
			// Parse form data
			err := req.ParseMultipartForm(32 << 20) // 32 MB limit
			if err != nil {
				log.Println("❌ Failed to parse form data", err.Error())
				return
			}
			creationDate := time.Now().Format("2006-01-02 15:04:05")
			modificationDate := time.Now().Format("2006-01-02 15:04:05")
			title := req.FormValue("title")
			slug := lib.Slugify(title)
			description := req.FormValue("description")
			_categories := req.FormValue("categories")

			imageUrl := lib.UploadImage(req)
			authorID := models.GetUserFromSession(req).ID
			categories := strings.Split(_categories, "#")

			post := models.Post{
				Title:        title,
				Slug:         slug,
				Description:  description,
				ImageURL:     imageUrl,
				AuthorID:     authorID,
				IsEdited:     false,
				CreateDate:   creationDate,
				ModifiedDate: modificationDate,
			}

			models.PostRepo.CreatePost(&post)

			for i := 1; i < len(categories); i++ {
				name := strings.TrimSpace(categories[i])
				category, _ := models.CategoryRepo.GetCategoryByName(name)
				if category == nil {
					category = &models.Category{
						Name:         name,
						CreateDate:   creationDate,
						ModifiedDate: modificationDate,
					}
					models.CategoryRepo.CreateCategory(category)
				}
				models.PostCategoryRepo.CreatePostCategory(category.ID, post.ID)
			}

			log.Println("✅ Post created with success")
			lib.RedirectToPreviousURL(res, req)
		}
	}
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, req.URL.Path, http.MethodGet) {
		PostComments := []models.Comment{}
		basePath := "base"
		pagePath := "post"

		isSessionOpen := models.ValidSession(req)

		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "posts" {
			post, err := models.PostRepo.GetPostBySlug(pathPart[2])
			if err != nil {
				fmt.Println("error DB")
				return
			}
			comments, err := models.CommentRepo.GetAllComments("15")
			if comments == nil {
				return
			}
			if err != nil {
				fmt.Println("error DB")
				return
			}
			for j := 0; j < len(comments); j++ {
				if post.ID == comments[j].PostID {
					PostComments = append(PostComments, *comments[j])
				}
			}
			userPost, err := models.UserRepo.GetUserByID(post.AuthorID)
			if err != nil {
				fmt.Println("error reading from user")
				return
			}
			PostPageData := PostPageData{
				IsLoggedIn: isSessionOpen,
				Post:       *post,
				UserPoster: userPost,
				Comments:   PostComments,
				NbrComment: len(PostComments),
			}
			lib.RenderPage(basePath, pagePath, PostPageData, res)
			log.Println("✅ Post page get with success")
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

		creationDate := time.Now().Format("2006-01-02")
		modifDate := time.Now().Format("2006-01-02")
		//--------------------------------------
		authorID := models.GetUserFromSession(req).ID
		parentID := "chjchjchjcxjchjc"
		postID := "709433aa-9fe4-4935-b1d6-48b50e24eb20"
		//--------------------------------------
		commentStruct := models.Comment{
			Text:         text,
			AuthorID:     authorID,
			PostID:       postID,
			ParentID:     parentID,
			CreateDate:   creationDate,
			ModifiedDate: modifDate}

		models.CommentRepo.CreateComment(&commentStruct)
	}
}
