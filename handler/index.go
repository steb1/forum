package handler

import (
	"fmt"
	"forum/data/models"
	"forum/lib"
	"log"
	"net/http"
)

type PostItem struct {
	ID                string
	Title             string
	ImageURL          string
	AuthorName        string
	LastEditionDate   string
	NumberOfComments  int
	ListOfCommentator []string
}

type HomePageData struct {
	IsLoggedIn   bool
	NumberOfPost int
	RandomUsers  []models.User
	CurrentUser  models.User
	AllPostItems []PostItem
}

func Index(res http.ResponseWriter, req *http.Request) {
	if lib.ValidateRequest(req, res, "/", http.MethodGet) {
		basePath := "base"
		pagePath := "index"

		isSessionOpen := models.ValidSession(req)
		user := models.GetUserFromSession(req)
		randomUsers, err := models.UserRepo.SelectRandomUsers(17)
		if err != nil {
			log.Println("❌ Can't get 17 random users in the database")
		}

		fmt.Println("---------------------1")
		// TODO: USE DIRECTLY A FUNCTION THAT RETURN A LIST OF POST ITEM
		allPosts, err := models.PostRepo.GetAllPosts("5")
		if err != nil {
			log.Println("❌ Can't get all post")
		}
		allPostItems := []PostItem{}

		for _, post := range allPosts {
			_postItem := PostItem{
				ID:               post.ID,
				Title:            post.Title,
				AuthorName:       "Ping",
				ImageURL:         post.ImageURL,
				LastEditionDate:  lib.TimeSinceCreation(post.CreateDate),
				NumberOfComments: 3,
				ListOfCommentator: []string{
					"/uploads/avatar.2.jpeg",
					"/uploads/avatar.3.jpeg",
					"/uploads/avatar.4.jpeg",
				},
			}
			allPostItems = append(allPostItems, _postItem)
		}

		homePageData := HomePageData{
			IsLoggedIn:   isSessionOpen,
			CurrentUser:  *user,
			RandomUsers:  randomUsers,
			NumberOfPost: models.PostRepo.GetNumberOfPosts(),
			AllPostItems: allPostItems,
		}

		lib.RenderPage(basePath, pagePath, homePageData, res)
		log.Println("✅ Home page get with success")

	}
}
