package auth

import (
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

type SignPageData struct {
	IsLoggedIn bool
	RandomUsers []models.User
}

func SignUp(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/sign-up", http.MethodPost) {
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "❌ On Signing Up %v", err)
			return
		}
		user := models.User{}
		user.Email = req.FormValue("email")
		user.Username = req.FormValue("username")

		_password, err := lib.HashPassword(req.FormValue("password"))
		if err != nil {
			log.Fatalf("❌ Failed to generate UUID: %v", err)
		}
		user.Password = _password

		// TODO: Handle the avatar upload
		avatarURL := lib.UploadImage(req)
		if avatarURL == "" {
			avatarURL = models.DEFAULT_AVATAR
		}
		user.AvatarURL = avatarURL
		user.Role = models.RoleUser

		if _, exist := models.UserRepo.IsExisted(user.Email); !exist {
			err := models.UserRepo.CreateUser(&user)
			if err != nil {
				log.Fatalf("❌ Failed to created account %v", err)
			}

			lib.NewSessionToken(res, user.ID, user.Username)

			http.Redirect(res, req, "/", http.StatusSeeOther)
			log.Println("✅ Account created with success")
		} else {
			fmt.Println("❌ User already exist")
			return
		}
	}
}

func SignUpPage(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/sign-up-page", http.MethodGet) {
		basePath := "base"
		pagePath := "sign-up"

		randomUsers, err := models.UserRepo.SelectRandomUsers(15)
		if err != nil {
			log.Println("❌ Can't get 15 random users in the database")
		}

		signPageData := SignPageData{
			IsLoggedIn:  false,
			RandomUsers: randomUsers,
		}

		lib.RenderPage(basePath, pagePath, signPageData, res)
		log.Println("✅ Register page get with success")
	}
}

func SignIn(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/sign-in", http.MethodPost) {
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "❌ On Signing In %v", err)
			return
		}
		user := models.User{}
		email := req.FormValue("email")
		password := req.FormValue("password")

		if _user, exist := models.UserRepo.IsExisted(email); exist {
			if !lib.IsPasswordsMatch(_user.Password, password) {
				log.Println("❌ Password given is wrong")
			} else {
				_user, err := models.UserRepo.GetUserByEmail(email)
				if err != nil {
					log.Println("❌ ", err)
				}
				user = *_user

				lib.NewSessionToken(res, user.ID, user.Username)

				http.Redirect(res, req, "/", http.StatusSeeOther)
				log.Println("✅ Sign in with success")
			}
		} else {
			log.Println("❌ User with the given email don't exist")
			return
		}
	}
}

func SignInPage(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/sign-in-page", http.MethodGet) {
		basePath := "base"
		pagePath := "sign-in"

		randomUsers, err := models.UserRepo.SelectRandomUsers(15)
		if err != nil {
			log.Println("❌ Can't get 15 random users in the database")
		}

		signPageData := SignPageData{
			IsLoggedIn:  false,
			RandomUsers: randomUsers,
		}

		lib.RenderPage(basePath, pagePath, signPageData, res)
		log.Println("✅ Login page get with success")
	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/logout", http.MethodGet) {
		if ok := lib.DeleteSession(req); ok {
			http.Redirect(res, req, "/", http.StatusSeeOther)
			log.Println("✅ Logout done with success")
		} else {
			log.Println("❌ Logout failure")
		}
	}
}
