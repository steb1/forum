package handler

import (
	"forum/lib"
	"log"
	"net/http"
)

type HomePage struct {
	IsLoggedIn bool
}

func Index(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/", http.MethodGet) {
		basePath := "base"
		pagePath := "index"

		isSessionOpen := lib.ValidSession(req)

		home := HomePage{
			IsLoggedIn: isSessionOpen,
		}

		lib.RenderPage(basePath, pagePath, home, res)
		log.Println("✅ Home page get with success")
	}
}
