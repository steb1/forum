package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

type PostItem struct {
	ID                string
	Title             string
	ImageURL          string
	AuthorName        string
	LastEditionDate   string
	NumberOfComments  int
	ListOfCommentator []string
}

type HomePageData struct {
	IsLoggedIn    bool
	RandomUsers   []models.User
	CurrentUser   models.User
	Post          []models.PostItem
	NumberOfPosts int
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

		posts, _ := models.PostRepo.GetAllPostsItems("5")

		homePageData := HomePageData{
			IsLoggedIn:    isSessionOpen,
			CurrentUser:   *user,
			RandomUsers:   randomUsers,
			Post:          posts,
			NumberOfPosts: models.PostRepo.GetNumberOfPosts(),
		}

		lib.RenderPage(basePath, pagePath, homePageData, res)
		log.Println("✅ Home page get with success")

	}
}
