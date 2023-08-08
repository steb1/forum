package auth

import (
	"encoding/json"
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

const tokenExpiration = time.Hour

type Token struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	ExpiresAt int64     `json:"exp"`
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

		// TODO: Hash the password
		user.Password = req.FormValue("password")

		// TODO: Handle the avatar upload
		user.AvatarURL = "assets/img/community.webp"
		user.Role = models.RoleUser

		if _, exist := models.UserRepo.IsExisted(user.Email); !exist {
			// TODO: Move uuid creation directly on the create model method
			ID, err := uuid.NewV4()
			if err != nil {
				log.Fatalf("❌ Failed to generate UUID: %v", err)
			}
			user.ID = ID.String()

			token := models.Token{
				UserID:    ID.String(),
				Username:  user.Username,
				ExpiresAt: time.Now().Add(tokenExpiration * 2),
			}

			tokenJson, err := json.Marshal(token)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			user.TokenExpirationDate = token.ExpiresAt.Format("2006-01-02 15:04:05")
			err = models.UserRepo.CreateUser(&user)
			if err != nil {
				log.Fatalf("❌ Failed to created account %v", err)
			}

			cookie := http.Cookie{}
			cookie.Name = user.Username
			cookie.Value = ID.String()
			cookie.Expires = time.Now().Add(2 * time.Hour)
			cookie.Secure = true
			cookie.HttpOnly = true
			http.SetCookie(res, &cookie)

			res.Header().Set("Content-Type", "application/json")
			res.Write(tokenJson)

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

		lib.RenderPage(basePath, pagePath, nil, res)
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
		email := req.FormValue("mail")
		username := req.FormValue("username")
		password := req.FormValue("password")

		if _user, exist := models.UserRepo.IsExisted(email); exist {
			if _user.Password != password {
				log.Println("❌ Password given is wrong")
				return
			} else {
				_user, err := models.UserRepo.GetUserByEmail(email)
				user = *_user

				token := models.Token{
					UserID:    user.ID,
					Username:  username,
					ExpiresAt: time.Now().Add(tokenExpiration * 2),
				}

				tokenJson, err := json.Marshal(token)
				if err != nil {
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}

				cookie := http.Cookie{}
				cookie.Name = token.Username
				cookie.Value = token.UserID
				cookie.Expires = token.ExpiresAt
				cookie.Secure = true
				cookie.HttpOnly = true
				http.SetCookie(res, &cookie)
				user.TokenExpirationDate = token.ExpiresAt.Format("2006-01-02 15:04:05")

				err = models.UserRepo.UpdateUser(&user)
				if err != nil {
					log.Fatal(err)
				}

				res.Header().Set("Content-Type", "application/json")
				res.Write(tokenJson)

				basePath := "base"
				pagePath := "index"

				lib.RenderPage(basePath, pagePath, token, res)
				log.Println("✅ Sign in with success")
			}
		} else {
			log.Println("❌ User with the give email don't exist")
			return
		}
	}
}

func SignInPage(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/sign-in-page", http.MethodGet) {
		basePath := "base"
		pagePath := "sign-in"

		lib.RenderPage(basePath, pagePath, nil, res)
		log.Println("✅ Login page get with success")
	}
}
