package handler

import (
	"forum/lib"
	"log"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/", http.MethodGet) {
		basePath := "base"
		pagePath := "index"

		lib.RenderPage(basePath, pagePath, nil, res)
		log.Println("✅ Home page get with success")
	}
}
