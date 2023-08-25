package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
)

type NotifPageData struct {
	IsLoggedIn    bool
	CurrentUser   models.User
	Notifications []models.Notification
	UserAuthor    []models.User
	Posts         []models.Post
	Allposts      []*models.Post
}

func GetNotifs(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/notification/*", http.MethodGet) {
		basePath := "base"
		pagePath := "notification"
		isSessionOpen := models.ValidSession(req)
		if !isSessionOpen {
			return
		}
		user := models.GetUserFromSession(req)
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "notification" && pathPart[2] != "" {
			id := pathPart[2]
			notifications, err := models.NotifRepo.GetAllNotifsByUser(id)
			if notifications == nil {
				return
			}
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				log.Println("❌ Can't get notifs")
				return
			}
			posts := []models.Post{}
			users := []models.User{}
			for i := 0; i < len(notifications); i++ {
				post, _ := models.PostRepo.GetPostByID(notifications[i].PostID)
				posts = append(posts, *post)
				userAuthor, _ := models.UserRepo.GetUserByID(notifications[i].AuthorID)
				users = append(users, *userAuthor)
			}
			allPost, err := models.PostRepo.GetAllPosts("")
			if err != nil {
				return
			}
			notifpagedata := NotifPageData{
				IsLoggedIn:    isSessionOpen,
				CurrentUser:   *user,
				Notifications: notifications,
				UserAuthor:    users,
				Posts:         posts,
				Allposts:      allPost,
			}

			lib.RenderPage(basePath, pagePath, notifpagedata, res)
			log.Println("✅ Notification page get with success")
		}
	}
}
