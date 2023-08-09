package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

type EditUserPageData struct {
	IsLoggedIn  bool
	CurrentUser models.User
}

func EditUser(res http.ResponseWriter, req *http.Request) {
	//TODO: Return http error
	if lib.ValidateRequest(req, res, "/edit-user", http.MethodPost) {
		// Check if the user is logged in
		currentUser := lib.GetUserFromSession(req)
		if currentUser == nil || currentUser.ID == "" {
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}

		// Parse form data
		err := req.ParseForm()
		if err != nil {
			log.Println("❌ Failed to parse form data")
			return
		}

		// Update user information
		username := req.Form.Get("username")
		email := req.Form.Get("email")
		currentUser.Username = username
		currentUser.Email = email

		// Update user information in the database
		err = models.UserRepo.UpdateUser(currentUser)
		if err != nil {
			log.Println("❌ Failed to update user information")
			return
		}

		// Redirect to the user's profile page
		http.Redirect(res, req, "/edit-user-page", http.StatusSeeOther)
		return
	}
}

func EditUserPage(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/edit-user-page", http.MethodGet) {
		basePath := "base"
		pagePath := "edit-user"

		isSessionOpen := lib.ValidSession(req)
		user := lib.GetUserFromSession(req)

		editUserPageData := EditUserPageData{
			IsLoggedIn:  isSessionOpen,
			CurrentUser: *user,
		}

		lib.RenderPage(basePath, pagePath, editUserPageData, res)
		log.Println("✅ Login page get with success")
	}
}
