package main

import (
	"forum/data/models"
	"forum/handler"
	"forum/handler/auth"
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := ":" + os.Getenv("PORT")
	ADDRESS := os.Getenv("ADDRESS")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./assets/styles/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./assets/img/"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/sign-up", auth.SignUp)
	http.HandleFunc("/sign-up-page", auth.SignUpPage)
	http.HandleFunc("/google-Sign-in", auth.HandleGoogleLogin)
	http.HandleFunc("/github-Sign-in", auth.HandleGithubLoginHandler)
	http.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		auth.LoggedinHandler(w, r, "")
	})
	http.HandleFunc("/callback", auth.HandleCallback)
	http.HandleFunc("/login/github/callback", auth.HandleGithubCallback)
	http.HandleFunc("/sign-in", auth.SignIn)
	http.HandleFunc("/sign-in-page", auth.SignInPage)
	http.HandleFunc("/logout", auth.Logout)

	http.HandleFunc("/profile", handler.ProfilePage)
	http.HandleFunc("/edit-user", handler.EditUser)
	http.HandleFunc("/edit-user-page", handler.EditUserPage)

	http.HandleFunc("/trending", handler.ListPost)
	http.HandleFunc("/post", handler.CreatePost)
	http.HandleFunc("/delete-post/", handler.DeletePost)
	http.HandleFunc("/edit-post-page/", handler.EditPostPage)
	http.HandleFunc("/edit-post/", handler.EditPost)
	http.HandleFunc("/comment/", handler.Comment)
	http.HandleFunc("/posts/", handler.GetPost)
	http.HandleFunc("/user/", handler.UserProfilePage)
	http.HandleFunc("/like/", handler.LikePost)
	http.HandleFunc("/dislike/", handler.DislikePost)
	http.HandleFunc("/like-comment/", handler.LikeComment)
	http.HandleFunc("/dislike-comment/", handler.DislikeComment)
	http.HandleFunc("/category/", handler.GetPostOfCategory)

	go models.DeleteExpiredSessions()

	log.Print("Server started and running on ")
	log.Println(ADDRESS + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
