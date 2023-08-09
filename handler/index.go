package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

type HomePageData struct {
	IsLoggedIn  bool
	CurrentUser models.User
}

func Index(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/", http.MethodGet) {
		basePath := "base"
		pagePath := "index"

		isSessionOpen := lib.ValidSession(req)
		user := lib.GetUserFromSession(req)

		homePageData := HomePageData{
			IsLoggedIn:  isSessionOpen,
			CurrentUser: *user,
		}

		lib.RenderPage(basePath, pagePath, homePageData, res)
		log.Println("✅ Home page get with success")
	}
}
