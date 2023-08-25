package handler

import (
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"sort"
	"strings"
	"text/template"
	"time"
)

func SortComments(comments []*models.CommentItem) []*models.CommentItem {
	commentMap := make(map[string][]*models.CommentItem)

	for _, comment := range comments {
		commentMap[comment.ParentID] = append(commentMap[comment.ParentID], comment)
	}

	var sortedComments []*models.CommentItem
	var dfs func(string, int)
	dfs = func(parentID string, depth int) {
		children := commentMap[parentID]
		sort.SliceStable(children, func(i, j int) bool {
			return children[i].Index < children[j].Index
		})
		for _, child := range children {
			child.Index = depth
			child.Depth = strings.Repeat(`<span class="ml-1"></span>`, child.Index)
			sortedComments = append(sortedComments, child)
			dfs(child.ID, depth+1)
		}
	}

	dfs("", 0)
	return sortedComments
}

func Comment(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/comment/*", http.MethodPost) {
		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		text := strings.TrimSpace(req.FormValue("text"))
		parentID := strings.TrimSpace(req.FormValue("parentID"))
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if text != "" {
			if len(pathPart) == 3 && pathPart[1] == "comment" {
				creationDate := time.Now().Format("2006-01-02 15:04:05")
				modifDate := time.Now().Format("2006-01-02 15:04:05")

				authorID := models.GetUserFromSession(req).ID
				postID := pathPart[2]

				commentStruct := models.Comment{
					Text:         text,
					AuthorID:     authorID,
					PostID:       postID,
					ParentID:     parentID,
					CreateDate:   creationDate,
					ModifiedDate: modifDate,
				}

				models.CommentRepo.CreateComment(&commentStruct)
				lib.RedirectToPreviousURL(res, req)
			}
		} else {
			lib.RedirectToPreviousURL(res, req)
		}

	}
}

func EditCommentPage(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/edit-comment-page/*", http.MethodGet) {
		basePath := "base"
		pagePath := "edit-comment"
		isSessionOpen := models.ValidSession(req)
		user := models.GetUserFromSession(req)

		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "edit-comment-page" && pathPart[2] != "" {
			id := pathPart[2]
			comment, err := models.CommentRepo.GetCommentByID(id)
			if comment == nil {
				return
			}
			post, err := models.PostRepo.GetPostByID(comment.PostID)
			nbrLike, err := models.ViewRepo.GetLikesByPost(post.ID)
			nbrDislike, err := models.ViewRepo.GetDislikesByPost(post.ID)
			postCategories, err := models.PostCategoryRepo.GetCategoriesOfPost(post.ID)
			post.Description = template.HTMLEscapeString(post.Description)
			post.Title = template.HTMLEscapeString(post.Title)
			post.ModifiedDate = lib.FormatDateDB(post.ModifiedDate)
			post.ModifiedDate = lib.TimeSinceCreation(post.ModifiedDate)
			if postCategories != nil {
				for k := 0; k < len(postCategories); k++ {
					postCategories[k].Name = template.HTMLEscapeString(postCategories[k].Name)
				}
			}
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				log.Println("❌ error DB")
				return
			}

			userPageData := PostPageData{
				IsLoggedIn:     isSessionOpen,
				CurrentUser:    *user,
				Comment:        *comment,
				Post:           *post,
				NbrLike:        nbrLike,
				NbrDislike:     nbrDislike,
				CategoriesPost: postCategories,
			}

			lib.RenderPage(basePath, pagePath, userPageData, res)
			log.Println("✅ " + pagePath + " page get with success")
		} else {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "404", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
	} else {
		res.WriteHeader(http.StatusNotFound)
		lib.RenderPage("base", "404", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL.Path)
		return
	}
}

func EditComment(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/edit-comment/*", http.MethodPost) {
		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "edit-comment" && pathPart[2] != "" {
			id := pathPart[2]
			comment, err := models.CommentRepo.GetCommentByID(id)
			if comment == nil {
				return
			}
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				log.Println("❌ error DB")
				return
			}
			comment.Text = req.FormValue("text")
			comment.ModifiedDate = time.Now().Format("2006-01-02 15:04:05")
			models.CommentRepo.UpdateComment(comment)
			http.Redirect(res, req, "/profile?index=4", http.StatusSeeOther)
		} else {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "404", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
	} else {
		res.WriteHeader(http.StatusNotFound)
		lib.RenderPage("base", "404", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL.Path)
		return
	}
}

func DeleteComment(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/delete-comment/*", http.MethodGet) {
		isSessionOpen := models.ValidSession(req)
		if isSessionOpen {
			path := req.URL.Path
			pathPart := strings.Split(path, "/")
			if len(pathPart) == 3 && pathPart[1] == "delete-comment" && pathPart[2] != "" {
				id := pathPart[2]
				models.CommentRepo.DeleteComment(id)

				log.Println("✅ comment deleted with success")
				http.Redirect(res, req, "/profile?index=4", http.StatusSeeOther)
			}
		} else {
			res.WriteHeader(http.StatusNotFound)
			lib.RenderPage("base", "404", nil, res)
			log.Println("404 ❌ - Page not found ", req.URL.Path)
			return
		}
	} else {
		res.WriteHeader(http.StatusNotFound)
		lib.RenderPage("base", "404", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL.Path)
		return
	}
}

// func EditComment(res http.ResponseWriter, req *http.Request) {
// 	if lib.ValidateRequest(req, res, "/edit-comment/*", http.MethodPost) {
// 		// Check if the user is logged in
// 		currentUser := models.GetUserFromSession(req)
// 		if currentUser == nil || currentUser.ID == "" {
// 			http.Redirect(res, req, "/", http.StatusSeeOther)
// 			return
// 		}

// 		err := req.ParseMultipartForm(32 << 20) // 32 MB limit
// 		if err != nil {
// 			log.Println("❌ Failed to parse form data", err.Error())
// 			return
// 		}

// 		path := req.URL.Path
// 		pathPart := strings.Split(path, "/")
// 		if len(pathPart) == 3 && pathPart[1] == "edit-post" && pathPart[2] != "" {
// 			idPost := pathPart[2]
// 			post, err := models.PostRepo.GetPostByID(idPost)
// 			if err != nil {
// 				res.WriteHeader(http.StatusInternalServerError)
// 				log.Println("❌ error DB")
// 				return
// 			}

// 			// Update user information
// 			title := req.FormValue("title")
// 			description := req.FormValue("description")
// 			_categories := req.FormValue("categories")
// 			categories := strings.Split(_categories, "#")
// 			isEdited := false
// 			if title != "" && post.Title != title {
// 				isEdited = true
// 				post.Title = title
// 				post.Slug = lib.Slugify(title)
// 				post.ModifiedDate = time.Now().Format("2006-01-02 15:04:05")
// 				log.Println("✅ Title changed successfully")
// 			}
// 			if description != "" && post.Description != description {
// 				isEdited = true
// 				post.Description = description
// 				log.Println("✅ Description changed successfully")
// 			}
// 			imageURL := lib.UploadImage(req)
// 			if imageURL != "" {
// 				isEdited = true
// 				post.ImageURL = imageURL
// 				log.Println("✅ Image changed successfully")
// 			}

// 			// Update user information in the database

// 			categoriesOfPost, err := models.PostCategoryRepo.GetCategoriesOfPost(idPost)
// 			if err != nil {
// 				log.Println("❌ Failed to update post information ", err.Error())
// 				return
// 			}

// 			for i := 1; i < len(categories); i++ {
// 				categories[i] = strings.TrimSpace(categories[i])
// 				found := false
// 				for _, cat := range categoriesOfPost {
// 					if cat.Name == categories[i] {
// 						found = true
// 						break
// 					}
// 				}
// 				if !found {
// 					creationDate := time.Now().Format("2006-01-02 15:04:05")
// 					modificationDate := time.Now().Format("2006-01-02 15:04:05")
// 					category := &models.Category{
// 						Name:         categories[i],
// 						CreateDate:   creationDate,
// 						ModifiedDate: modificationDate,
// 					}
// 					isEdited = true
// 					models.CategoryRepo.CreateCategory(category)
// 					models.PostCategoryRepo.CreatePostCategory(category.ID, post.ID)
// 				}
// 			}

// 			for _, category := range categoriesOfPost {
// 				found := false
// 				for _, cat := range categories {
// 					if category.Name == cat {
// 						found = true
// 						break
// 					}
// 				}
// 				if !found {
// 					isEdited = true
// 					err = models.PostCategoryRepo.DeletePostCategory(category.ID, currentUser.ID)
// 					if err != nil {
// 						log.Println("❌ Failed to delete category post information ", err.Error())
// 						return
// 					}
// 				}
// 			}

// 			if isEdited {
// 				post.IsEdited = true
// 				post.ModifiedDate = time.Now().Format("2006-01-02 15:04:05")
// 				err = models.PostRepo.UpdatePost(post)
// 				if err != nil {
// 					log.Println("❌ Failed to update post information ", err.Error())
// 					return
// 				}
// 			}

// 			// Redirect to the user's profile page
// 			http.Redirect(res, req, "/profile", http.StatusSeeOther)
// 		}
// 	}
// }
