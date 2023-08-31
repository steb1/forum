package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

type RequestPageData struct {
	IsLoggedIn  bool
	CurrentUser models.User
	Requests    []models.Request
	Allposts    []*models.Post
}

func SeeRequests(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/requests", http.MethodGet) {
		basePath := "base"
		pagePath := "admin/request"
		isSessionOpen := models.ValidSession(req)
		currentuser := models.GetUserFromSession(req)
		if !isSessionOpen || currentuser.Role != 0 {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
		requests, err := models.RequestRepo.GetAllRequest()
		if err != nil {
			return
		}
		allPosts, err := models.PostRepo.GetAllPosts("")
		if err != nil {
			return
		}
		requestPageData := RequestPageData{
			IsLoggedIn:  isSessionOpen,
			CurrentUser: *currentuser,
			Requests:    requests,
			Allposts:    allPosts,
		}
		lib.RenderPage(basePath, pagePath, requestPageData, res)
		log.Println("✅ Home111 page get with success")
	}
}
