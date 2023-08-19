package handler

import (
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"strings"
)

func LikeComment(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/like-comment/*", http.MethodGet) {
		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "like-comment" && pathPart[2] != "" {
			commentID := pathPart[2]
			comment, err := models.CommentRepo.GetCommentByID(commentID)
			if comment == nil {
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
			commentView, err := models.CommentViewRepo.GetViewByAuthorIDandCommentID(user.ID, comment.ID)
			if err != nil {
				fmt.Println("error Reading from View")
				return
			}
			if commentView == nil {
				NewView := models.Comment_View{
					Rate:      1,
					AuthorID:  user.ID,
					CommentID: comment.ID,
				}
				models.CommentViewRepo.CreateCommentView(&NewView)
				lib.RedirectToPreviousURL(res, req)
			} else {
				if commentView.Rate == 0 || commentView.Rate == 2 {
					UpdateView := models.Comment_View{
						ID:        commentView.ID,
						Rate:      1,
						AuthorID:  user.ID,
						CommentID: comment.ID,
					}
					models.CommentViewRepo.UpdateView(&UpdateView)
					lib.RedirectToPreviousURL(res, req)
				} else if commentView.Rate == 1 {
					UpdateView := models.Comment_View{
						ID:        commentView.ID,
						Rate:      0,
						AuthorID:  user.ID,
						CommentID: comment.ID,
					}
					models.CommentViewRepo.UpdateView(&UpdateView)
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

func DislikeComment(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/dislike-comment/*", http.MethodGet) {
		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "dislike-comment" && pathPart[2] != "" {
			commentID := pathPart[2]
			comment, err := models.CommentRepo.GetCommentByID(commentID)
			if comment == nil {
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
			commentView, err := models.CommentViewRepo.GetViewByAuthorIDandCommentID(user.ID, comment.ID)
			if err != nil {
				fmt.Println("error Reading from View")
				return
			}
			if commentView == nil {
				NewView := models.Comment_View{
					Rate:      2,
					AuthorID:  user.ID,
					CommentID: comment.ID,
				}
				models.CommentViewRepo.CreateCommentView(&NewView)
				lib.RedirectToPreviousURL(res, req)
			} else {
				if commentView.Rate == 0 || commentView.Rate == 1 {
					UpdateView := models.Comment_View{
						ID:        commentView.ID,
						Rate:      2,
						AuthorID:  user.ID,
						CommentID: comment.ID,
					}
					models.CommentViewRepo.UpdateView(&UpdateView)
					lib.RedirectToPreviousURL(res, req)
				} else if commentView.Rate == 2 {
					UpdateView := models.Comment_View{
						ID:        commentView.ID,
						Rate:      0,
						AuthorID:  user.ID,
						CommentID: comment.ID,
					}
					models.CommentViewRepo.UpdateView(&UpdateView)
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
