package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

func SeePosts(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/posts", http.MethodGet) {
		isSessionOpen := models.ValidSession(req)
		currentuser := models.GetUserFromSession(req)
		if !isSessionOpen || currentuser.Role != 0 {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}

	} else {
		res.WriteHeader(http.StatusNotFound)
		lib.RenderPage("base", "", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL.Path)
		return
	}
}
