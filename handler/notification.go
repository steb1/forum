package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
)

type NotifPageData struct {
	IsLoggedIn        bool
	CurrentUser       models.User
	NotifsID          []string
	NotificationsType []string
	UserAuthor        []string
	Posts             []string
	Allposts          []*models.Post
	FormatedNotif     []string
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
			tabNotifType := []string{}
			tabNotifID := []string{}
			posts := []string{}
			users := []string{}
			
			for i := 0; i < len(notifications); i++ {
				post, _ := models.PostRepo.GetPostByID(notifications[i].PostID)
				posts = append(posts, post.Title)
				userAuthor, _ := models.UserRepo.GetUserByID(notifications[i].AuthorID)
				users = append(users, userAuthor.Username)
				tabNotifType = append(tabNotifType, notifications[i].Notif_type)
				tabNotifID = append(tabNotifID, notifications[i].ID)

			}
			FormatedNotif := (models.FormatNotifications(notifications))
			allPost, err := models.PostRepo.GetAllPosts("")
			if err != nil {
				return
			}
			notifpagedata := NotifPageData{
				IsLoggedIn:        isSessionOpen,
				CurrentUser:       *user,
				NotifsID:          tabNotifID,
				NotificationsType: tabNotifType,
				UserAuthor:        users,
				Posts:             posts,
				Allposts:          allPost,
				FormatedNotif:     FormatedNotif,
				
			}

			lib.RenderPage(basePath, pagePath, notifpagedata, res)
			log.Println("✅ Notification page get with success")
			notifs, _ := models.NotifRepo.GetAllNotifsByUser(id)
			notif := models.FormatNotifications(notifs)
			log.Println("----------------------------------------------------------------------")
			for _, val := range notif {
				log.Println(val)
				log.Println()
				log.Println("----------------------------------------------------------------------")
			}

		}
	}
}
