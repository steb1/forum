package handler

import (
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

type RequestPostPageData struct {
	AllPost []models.PostItem
}

func SeePosts(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/posts", http.MethodGet) {
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
		PostPage := RequestPostPageData{
			AllPost: tabPostItems,
		}
		fmt.Println(PostPage)

	} else {
		res.WriteHeader(http.StatusNotFound)
		lib.RenderPage("base", "", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL.Path)
		return
	}
}
