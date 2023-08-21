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
			commentRate, err := models.CommentRateRepo.GetRateByAuthorIDandCommentID(user.ID, comment.ID)
			if err != nil {
				fmt.Println("error Reading from Rate")
				return
			}
			if commentRate == nil {
				NewRate := models.CommentRate{
					Rate:      1,
					AuthorID:  user.ID,
					CommentID: comment.ID,
				}
				models.CommentRateRepo.CreateCommentRate(&NewRate)
				lib.RedirectToPreviousURL(res, req)
			} else {
				if commentRate.Rate == 0 || commentRate.Rate == 2 {
					UpdateRate := models.CommentRate{
						ID:        commentRate.ID,
						Rate:      1,
						AuthorID:  user.ID,
						CommentID: comment.ID,
					}
					models.CommentRateRepo.UpdateRate(&UpdateRate)
					lib.RedirectToPreviousURL(res, req)
				} else if commentRate.Rate == 1 {
					UpdateRate := models.CommentRate{
						ID:        commentRate.ID,
						Rate:      0,
						AuthorID:  user.ID,
						CommentID: comment.ID,
					}
					models.CommentRateRepo.UpdateRate(&UpdateRate)
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
			commentRate, err := models.CommentRateRepo.GetRateByAuthorIDandCommentID(user.ID, comment.ID)
			if err != nil {
				fmt.Println("error Reading from Rate")
				return
			}
			if commentRate == nil {
				NewRate := models.CommentRate{
					Rate:      2,
					AuthorID:  user.ID,
					CommentID: comment.ID,
				}
				models.CommentRateRepo.CreateCommentRate(&NewRate)
				lib.RedirectToPreviousURL(res, req)
			} else {
				if commentRate.Rate == 0 || commentRate.Rate == 1 {
					UpdateRate := models.CommentRate{
						ID:        commentRate.ID,
						Rate:      2,
						AuthorID:  user.ID,
						CommentID: comment.ID,
					}
					models.CommentRateRepo.UpdateRate(&UpdateRate)
					lib.RedirectToPreviousURL(res, req)
				} else if commentRate.Rate == 2 {
					UpdateRate := models.CommentRate{
						ID:        commentRate.ID,
						Rate:      0,
						AuthorID:  user.ID,
						CommentID: comment.ID,
					}
					models.CommentRateRepo.UpdateRate(&UpdateRate)
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
