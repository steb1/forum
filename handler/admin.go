package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
)

type RequestPageData struct {
	IsLoggedIn  bool
	CurrentUser models.User
	Requests    []models.Request
	AllPosts    []*models.Post
}

func SeeRequests(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/requests", http.MethodGet) {
		basePath := "base"
		pagePath := "admin/request"
		isSessionOpen := models.ValidSession(req)
		currentuser := models.GetUserFromSession(req)
		if !isSessionOpen || currentuser.Role != 0 {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
		requests, err := models.RequestRepo.GetAllRequest()
		if err != nil {
			return
		}
		allPosts, err := models.PostRepo.GetAllPosts("")
		if err != nil {
			return
		}
		requestPageData := RequestPageData{
			IsLoggedIn:  isSessionOpen,
			CurrentUser: *currentuser,
			Requests:    requests,
			AllPosts:    allPosts,
		}
		lib.RenderPage(basePath, pagePath, requestPageData, res)
		log.Println("✅ Home111 page get with success")
	}
}

func Validate(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/validate/*", http.MethodGet) {
		isSessionOpen := models.ValidSession(req)
		currentuser := models.GetUserFromSession(req)
		if !isSessionOpen || currentuser.Role != 0 {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "validate" && pathPart[2] != "" {
			id := pathPart[2]
			user, err := models.UserRepo.GetUserByID(id)
			if user == nil {
				res.WriteHeader(http.StatusNotFound)
				lib.RenderPage("base", "", nil, res)
				log.Println("404 ❌ - Page not found ", req.URL.Path)
				return
			}
			if err != nil {
				return
			}
			UpdateUser := models.User{
				ID:        user.ID,
				Username:  user.Username,
				Email:     user.Email,
				Password:  user.Password,
				AvatarURL: user.AvatarURL,
				Role:      models.RoleModerator,
			}
			err = models.UserRepo.UpdateUser(&UpdateUser)
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				log.Println("❌ error Update user")
				return
			}
			request, err := models.RequestRepo.GetRequestByUser(user.ID)
			if err != nil {
				return
			}
			UpdateRequest := models.Request{
				ID:       request.ID,
				AuthorID: request.AuthorID,
				Time:     request.Time,
				Username: request.Username,
				ImageURL: request.ImageURL,
				Role:     models.RoleModerator,
			}
			err = models.RequestRepo.UpdateRequest(&UpdateRequest)
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				log.Println("❌ error Update user")
				return
			}
			lib.RedirectToPreviousURL(res, req)
		} else {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
	} else {
		res.WriteHeader(http.StatusNotFound)
		lib.RenderPage("base", "", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL.Path)
		return
	}
}

func Invalidate(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/invalidate/*", http.MethodGet) {
		isSessionOpen := models.ValidSession(req)
		currentuser := models.GetUserFromSession(req)
		if !isSessionOpen || currentuser.Role != 0 {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "invalidate" && pathPart[2] != "" {
			id := pathPart[2]
			user, err := models.UserRepo.GetUserByID(id)
			if user == nil {
				res.WriteHeader(http.StatusNotFound)
				lib.RenderPage("base", "", nil, res)
				log.Println("404 ❌ - Page not found ", req.URL.Path)
				return
			}
			if err != nil {
				return
			}
			request, err := models.RequestRepo.GetRequestByUser(user.ID)
			if request == nil {
				res.WriteHeader(http.StatusNotFound)
				lib.RenderPage("base", "", nil, res)
				log.Println("404 ❌ - Page not found ", req.URL.Path)
				return
			}
			if err != nil {
				return
			}
			models.RequestRepo.DeleteRequest(request.ID)
			lib.RedirectToPreviousURL(res, req)
		} else {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
	} else {
		res.WriteHeader(http.StatusNotFound)
		lib.RenderPage("base", "", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL.Path)
		return
	}
}
