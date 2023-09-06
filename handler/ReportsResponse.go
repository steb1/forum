package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

type ResponseReportPage struct {
	IsLoggedIn  bool
	CurrentUser models.User
	AllResponse []models.Response
	AllNotifs []*models.Notification
	AllPosts []*models.Post
}

func SeeReportsResponse(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/seeReportsResponse", http.MethodGet) {
		isSessionOpen := models.ValidSession(req)
		currentUser := models.GetUserFromSession(req)
		AllResponse, err := models.ResponseRepo.GetAllResponse()

		if err != nil {
			log.Println("1")
			return
		}
		notifications, err := models.NotifRepo.GetAllNotifs()

		if err != nil {
			log.Println("2")
			return
		}

		allPosts, err := models.PostRepo.GetAllPosts("")
		if err != nil {
			log.Println("3")
			return
		}

		ResponseReportPage := ResponseReportPage{
			IsLoggedIn:  isSessionOpen,
			CurrentUser: *currentUser,
			AllResponse: AllResponse,
			AllNotifs: notifications,
			AllPosts: allPosts,
		}

		log.Println("Response Page get with success")
		lib.RenderPage("base", "admin/adminResponse", ResponseReportPage, res)
	}
}
