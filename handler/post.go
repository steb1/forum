package handler

import (
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type PostPageData struct {
	IsLoggedIn bool
	Post       models.Post
	Comments   []*models.CommentItem
	UserPoster *models.User
	NbrComment int
}

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

func Post(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/post", http.MethodPost) {
		isSessionOpen := models.ValidSession(req)
		if isSessionOpen {
			// Parse form data
			err := req.ParseMultipartForm(32 << 20) // 32 MB limit
			if err != nil {
				log.Println("❌ Failed to parse form data", err.Error())
				return
			}
			creationDate := time.Now().Format("2006-01-02 15:04:05")
			modificationDate := time.Now().Format("2006-01-02 15:04:05")
			title := req.FormValue("title")
			slug := lib.Slugify(title)
			description := req.FormValue("description")
			_categories := req.FormValue("categories")

			imageUrl := lib.UploadImage(req)
			authorID := models.GetUserFromSession(req).ID
			categories := strings.Split(_categories, "#")

			post := models.Post{
				Title:        title,
				Slug:         slug,
				Description:  description,
				ImageURL:     imageUrl,
				AuthorID:     authorID,
				IsEdited:     false,
				CreateDate:   creationDate,
				ModifiedDate: modificationDate,
			}

			models.PostRepo.CreatePost(&post)

			for i := 1; i < len(categories); i++ {
				name := strings.TrimSpace(categories[i])
				category, _ := models.CategoryRepo.GetCategoryByName(name)
				if category == nil {
					category = &models.Category{
						Name:         name,
						CreateDate:   creationDate,
						ModifiedDate: modificationDate,
					}
					models.CategoryRepo.CreateCategory(category)
				}
				models.PostCategoryRepo.CreatePostCategory(category.ID, post.ID)
			}

			log.Println("✅ Post created with success")
			lib.RedirectToPreviousURL(res, req)
		}
	}
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/posts/*", http.MethodGet) {
		basePath := "base"
		pagePath := "post"

		isSessionOpen := models.ValidSession(req)

		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		path := req.URL.Path
		pathPart := strings.Split(path, "/")
		if len(pathPart) == 3 && pathPart[1] == "posts" {
			slug := pathPart[2]
			post, err := models.PostRepo.GetPostBySlug(slug)
			if err != nil {
				fmt.Println("error DB")
				return
			}
			PostComments, err := models.CommentRepo.GetCommentsOfPost(post.ID, "15")
			PostComments = SortComments(PostComments)
			post.ModifiedDate = strings.ReplaceAll(post.ModifiedDate, "T", " ")
			post.ModifiedDate = strings.ReplaceAll(post.ModifiedDate, "Z", "")
			post.ModifiedDate = lib.TimeSinceCreation(post.ModifiedDate)
			if err != nil {
				fmt.Println("error DB")
				return
			}
			userPost, err := models.UserRepo.GetUserByID(post.AuthorID)
			if err != nil {
				fmt.Println("error reading from user")
				return
			}
			PostPageData := PostPageData{
				IsLoggedIn: isSessionOpen,
				Post:       *post,
				UserPoster: userPost,
				Comments:   PostComments,
				NbrComment: len(PostComments),
			}
			lib.RenderPage(basePath, pagePath, PostPageData, res)
			log.Println("✅ Post page get with success")
		} else {
			http.NotFound(res, req)
		}
	}
}

func Comment(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/comment", http.MethodPost) {
		err := req.ParseForm()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		text := req.FormValue("text")

		creationDate := time.Now().Format("2006-01-02")
		modifDate := time.Now().Format("2006-01-02")
		//--------------------------------------
		authorID := models.GetUserFromSession(req).ID
		parentID := "chjchjchjcxjchjc"
		postID := "709433aa-9fe4-4935-b1d6-48b50e24eb20"
		//--------------------------------------
		commentStruct := models.Comment{
			Text:         text,
			AuthorID:     authorID,
			PostID:       postID,
			ParentID:     parentID,
			CreateDate:   creationDate,
			ModifiedDate: modifDate}

		models.CommentRepo.CreateComment(&commentStruct)
	}
}
