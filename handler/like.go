package handler

import (
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
)

func LikePost(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/like/*", http.MethodGet) {
		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "like" && pathPart[2] != "" {
			slug := pathPart[2]
			post, err := models.PostRepo.GetPostBySlug(slug)
			if post == nil {
				return
			}
			if err != nil {
				fmt.Println("error DB")
				return
			}
			user := models.GetUserFromSession(req)
			if user == nil {
				return
			}
			view, err := models.ViewRepo.GetViewByAuthorIDandPostID(user.ID, post.ID)
			if err != nil {
				fmt.Println("error Reading from View")
				return
			}
			if view == nil {
				NewView := models.View{
					IsBookmarked: false,
					Rate:         1,
					AuthorID:     user.ID,
					PostID:       post.ID,
				}
				models.ViewRepo.CreateView(&NewView)
				lib.RedirectToPreviousURL(res, req)
			} else {
				if view.Rate == 0 || view.Rate == 2 {
					UpdateView := models.View{
						ID:           view.ID,
						IsBookmarked: false,
						Rate:         1,
						AuthorID:     user.ID,
						PostID:       post.ID,
					}
					models.ViewRepo.UpdateView(&UpdateView)
					lib.RedirectToPreviousURL(res, req)
				} else if view.Rate == 1 {
					UpdateView := models.View{
						ID:           view.ID,
						IsBookmarked: false,
						Rate:         0,
						AuthorID:     user.ID,
						PostID:       post.ID,
					}
					models.ViewRepo.UpdateView(&UpdateView)
					lib.RedirectToPreviousURL(res, req)
				}
			}
		} else {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "404", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
		}
	}
}
