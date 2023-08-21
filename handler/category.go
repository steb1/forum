package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
)

func GetPostOfCategory(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/category/*", http.MethodGet) {
		basePath := "base"
		pagePath := "list-post"

		isSessionOpen := models.ValidSession(req)
		user := models.GetUserFromSession(req)
		limit := 10

		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "category" && pathPart[2] != "" {
			name := pathPart[2]
			posts, _ := models.PostCategoryRepo.GetPostsOfCategory(name)

			TopUsers, err := models.UserRepo.TopUsers()
			if err != nil {
				log.Println("❌ Can't get top users")
			}

			homePageData := ListPostsPageData{
				Title:       "Category: " + name,
				IsLoggedIn:  isSessionOpen,
				CurrentUser: *user,
				Post:        posts,
				TopUsers:    TopUsers,
				Limit:       limit,
			}

			lib.RenderPage(basePath, pagePath, homePageData, res)
			log.Println("✅ Home page get with success")
		}
	}
}
