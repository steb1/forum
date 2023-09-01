package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
)

type RequestPostPageData struct {
	IsLoggedIn  bool
	CurrentUser models.User
	PostsList   []models.PostItem
	AllPosts    []*models.Post
}

func SeePosts(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/posts", http.MethodGet) {
		basePath := "base"
		pagePath := "admin/posts"
		isSessionOpen := models.ValidSession(req)
		currentuser := models.GetUserFromSession(req)
		if !isSessionOpen || currentuser.Role != 1 {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
		tabPostItems, err := models.PostRepo.GetAllNoValidedPosts()
		if err != nil {
			log.Println("Error getting no valided posts ", req.URL.Path)
			return
		}
		allPost, err := models.PostRepo.GetAllPosts("")
		if err != nil {
			return
		}
		PostPage := RequestPostPageData{
			IsLoggedIn:  isSessionOpen,
			CurrentUser: *currentuser,
			PostsList:   tabPostItems,
			AllPosts:    allPost,
		}
		lib.RenderPage(basePath, pagePath, PostPage, res)
		log.Println("✅ success")
	} else {
		res.WriteHeader(http.StatusNotFound)
		lib.RenderPage("base", "", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL.Path)
		return
	}
}

func Publish(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/publish/*", http.MethodGet) {
		isSessionOpen := models.ValidSession(req)
		currentuser := models.GetUserFromSession(req)
		if !isSessionOpen || currentuser.Role != 1 {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "publish" && pathPart[2] != "" {
			slug := pathPart[2]
			post, err := models.PostRepo.GetPostBySlug(slug)
			if post == nil {
				res.WriteHeader(http.StatusNotFound)
				lib.RenderPage("base", "", nil, res)
				log.Println("404 ❌ - Page not found ", req.URL.Path)
				return
			}
			if err != nil {
				return
			}
			UpdatePost := models.Post{
				ID:           post.ID,
				Title:        post.Title,
				Slug:         post.Slug,
				Description:  post.Description,
				ImageURL:     post.ImageURL,
				AuthorID:     post.AuthorID,
				IsEdited:     post.IsEdited,
				CreateDate:   post.CreateDate,
				ModifiedDate: post.ModifiedDate,
				Validate:     true,
			}
			err = models.PostRepo.UpdatePost(&UpdatePost)
			if err != nil {
				return
			}
			lib.RedirectToPreviousURL(res, req)
		}
		log.Println("✅ publish success")
	} else {
		res.WriteHeader(http.StatusNotFound)
		lib.RenderPage("base", "", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL.Path)
		return
	}
}
