package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strconv"
)

type ListPostsPageData struct {
	IsLoggedIn    bool
	CurrentUser   models.User
	Post          []models.PostItem
	NumberOfPosts int
	Limit         int
	TopUsers      []models.TopUser
}

func ListPost(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/trending", http.MethodGet) {
		basePath := "base"
		pagePath := "list-post"

		isSessionOpen := models.ValidSession(req)
		user := models.GetUserFromSession(req)
		queryParams := req.URL.Query()
		limit := 10
		numberOfPosts := models.PostRepo.GetNumberOfPosts()

		if len(queryParams["limit"]) != 0 {
			_limit, err := strconv.Atoi(queryParams.Get("limit"))
			if limit <= 0 || err != nil {
				log.Println("❌ Can't convert index to int")
			} else {
				if _limit == numberOfPosts {
					limit = -1
				} else {
					limit = _limit
				}
			}
		}

		posts, _ := models.PostRepo.GetAllPostsItems(limit)

		TopUsers, err := models.UserRepo.TopUsers()
		if err != nil {
			log.Println("❌ Can't get top users")
		}

		if limit != -1 {
			if limit+5 > numberOfPosts {
				limit = numberOfPosts
			} else {
				limit += 5
			}
		}

		homePageData := ListPostsPageData{
			IsLoggedIn:    isSessionOpen,
			CurrentUser:   *user,
			Post:          posts,
			NumberOfPosts: numberOfPosts,
			TopUsers:      TopUsers,
			Limit:         limit,
		}

		lib.RenderPage(basePath, pagePath, homePageData, res)
		log.Println("✅ Home page get with success")
	}
}
