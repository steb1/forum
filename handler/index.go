package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strconv"
)

type HomePageData struct {
	IsLoggedIn    bool
	RandomUsers   []models.User
	CurrentUser   models.User
	Post          []models.PostItem
	NumberOfPosts int
	Limit         int
	TopUsers      []models.TopUser
}

func Index(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/", http.MethodGet) {
		basePath := "base"
		pagePath := "index"

		isSessionOpen := models.ValidSession(req)
		user := models.GetUserFromSession(req)
		queryParams := req.URL.Query()
		limit := 5
		if len(queryParams["limit"]) != 0 {
			_limit, err := strconv.Atoi(queryParams.Get("limit"))
			if err != nil {
				log.Println("❌ Can't convert index to int")
			} else {
				limit = _limit
			}
		}
		randomUsers, err := models.UserRepo.SelectRandomUsers(17)
		if err != nil {
			log.Println("❌ Can't get 17 random users in the database")
		}

		posts, _ := models.PostRepo.GetAllPostsItems(limit)

		TopUsers, err := models.UserRepo.TopUsers()
		if err != nil {
			log.Println("❌ Can't get top users")
		}
		numberOfPosts := models.PostRepo.GetNumberOfPosts()

		if limit + 5 > numberOfPosts {
			limit = numberOfPosts
		} else {
			limit += 5
		}

		homePageData := HomePageData{
			IsLoggedIn:    isSessionOpen,
			CurrentUser:   *user,
			RandomUsers:   randomUsers,
			Post:          posts,
			NumberOfPosts: models.PostRepo.GetNumberOfPosts(),
			TopUsers:      TopUsers,
			Limit:         limit,
		}

		lib.RenderPage(basePath, pagePath, homePageData, res)
		log.Println("✅ Home page get with success")
	}
}
