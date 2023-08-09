package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

type HomePageData struct {
	IsLoggedIn  bool
	RandomUsers []models.User
	CurrentUser models.User
}

func Index(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/", http.MethodGet) {
		basePath := "base"
		pagePath := "index"

		isSessionOpen := models.ValidSession(req)
		user := models.GetUserFromSession(req)
		randomUsers, err := models.UserRepo.SelectRandomUsers(17)
		if err != nil {
			log.Println("❌ Can't get 17 random users in the database")
		}

		homePageData := HomePageData{
			IsLoggedIn:  isSessionOpen,
			CurrentUser: *user,
			RandomUsers: randomUsers,
		}

		lib.RenderPage(basePath, pagePath, homePageData, res)
		log.Println("✅ Home page get with success")
	}
}
