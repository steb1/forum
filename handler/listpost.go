package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type ListPostsPageData struct {
	Title         string
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
				res.WriteHeader(http.StatusBadRequest)
				log.Println("❌ Can't convert index to int")
				return
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
			res.WriteHeader(http.StatusInternalServerError)
			log.Println("❌ Can't get top users")
			return
		}

		if limit != -1 {
			if limit+5 > numberOfPosts {
				limit = numberOfPosts
			} else {
				limit += 5
			}
		}
		if posts != nil {
			for j := 0; j < len(posts); j++ {
				posts[j].Title = template.HTMLEscapeString(posts[j].Title)
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
