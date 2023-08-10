package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strconv"
)

type UserPageData struct {
	IsLoggedIn  bool
	CurrentUser models.User
	TabIndex    int
}

func ProfilePage(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/profile", http.MethodGet) {
		basePath := "base"
		pagePath := "user/profile"
		queryParams := req.URL.Query()
		TabIndex := 1
		if len(queryParams["index"]) != 0 {
			_tabIndex, err := strconv.Atoi(queryParams.Get("index"))
			if err != nil {
				log.Println("❌ Can't convert index to int")
			} else {
				TabIndex = _tabIndex
			}
		} else {
			log.Println("❌ Index parameter missing")
		}
		isSessionOpen := models.ValidSession(req)
		user := models.GetUserFromSession(req)

		userPageData := UserPageData{
			IsLoggedIn:  isSessionOpen,
			CurrentUser: *user,
			TabIndex:    TabIndex,
		}

		lib.RenderPage(basePath, pagePath, userPageData, res)
		log.Println("✅ Login page get with success")
	}
}

func EditUser(res http.ResponseWriter, req *http.Request) {
	//TODO: Return http error
	if lib.ValidateRequest(req, res, "/edit-user", http.MethodPost) {
		// Check if the user is logged in
		currentUser := models.GetUserFromSession(req)
		if currentUser == nil || currentUser.ID == "" {
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}

		// Parse form data
		err := req.ParseMultipartForm(32 << 20) // 32 MB limit
		if err != nil {
			log.Println("❌ Failed to parse form data")
			return
		}

		// Update user information
		username := req.Form.Get("username")
		email := req.Form.Get("email")
		newPassword := req.Form.Get("new_password")
		confirmPassword := req.Form.Get("confirm_password")
		oldPassword := req.Form.Get("old_password")
		if username != "" && currentUser.Username != username {
			currentUser.Username = username
			log.Println("✅ Username changed successfully")
		}
		if email != "" && currentUser.Email != email {
			currentUser.Email = email
			log.Println("✅ Email changed successfully")
		}
		if newPassword != "" {
			if lib.IsPasswordsMatch(currentUser.Password, oldPassword) {
				if !lib.IsPasswordsMatch(currentUser.Password, newPassword) {
					if newPassword == confirmPassword {
						newPassword, err = lib.HashPassword(newPassword)
						if err != nil {
							log.Println("❌ Failed to hash password")
							return
						}
						currentUser.Password = newPassword
						log.Println("✅ Password changed successfully")
					}
				}
			}
		}
		avatarURL := lib.UploadImage(req)
		if avatarURL != "" {
			currentUser.AvatarURL = avatarURL
			log.Println("✅ Avatar changed successfully")
		}

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
		pagePath := "user/edit-user"
		queryParams := req.URL.Query()
		TabIndex := 1
		if len(queryParams["index"]) != 0 {
			_tabIndex, err := strconv.Atoi(queryParams.Get("index"))
			if err != nil {
				log.Println("❌ Can't convert index to int")
			} else {
				TabIndex = _tabIndex
			}
		} else {
			log.Println("❌ Index parameter missing")
		}
		isSessionOpen := models.ValidSession(req)
		user := models.GetUserFromSession(req)

		userPageData := UserPageData{
			IsLoggedIn:  isSessionOpen,
			CurrentUser: *user,
			TabIndex:    TabIndex,
		}

		lib.RenderPage(basePath, pagePath, userPageData, res)
		log.Println("✅ Login page get with success")
	}
}
