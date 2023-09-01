package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
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
