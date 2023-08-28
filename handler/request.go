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
	if lib.ValidateRequest(req, res, "/request/*", http.MethodPost) {
		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		text := strings.TrimSpace(req.FormValue("motif"))
		path := req.URL.Path
		if text != "" {
			pathPart := strings.Split(path, "/")
			if len(pathPart) == 3 && pathPart[1] == "request" && pathPart[2] != "" {
				id := pathPart[2]
				request := models.Request{
					AuthorID:   id,
					Motivation: text,
					Time:       time.Now().Format("2006-01-02 15:04:05"),
				}
				models.RequestRepo.CreateRequest(&request)
				log.Println("✅ Request success")
				lib.RedirectToPreviousURL(res, req)
			}
		}
	}
}
