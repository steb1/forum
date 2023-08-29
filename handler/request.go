package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
	"time"
)

func CreateRequest(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/request/*", http.MethodGet) {
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "request" && pathPart[2] != "" {
			id := pathPart[2]
			request := models.Request{
				AuthorID: id,
				Time:     time.Now().Format("2006-01-02 15:04:05"),
			}
			models.RequestRepo.CreateRequest(&request)
			log.Println("✅ Request success")
			lib.RedirectToPreviousURL(res, req)
		}
	}
}
