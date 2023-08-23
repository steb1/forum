package main

import (
	"forum/data/models"
	"forum/handler"
	"forum/handler/auth"
	"forum/lib"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	PORT := ":" + os.Getenv("PORT")
	ADDRESS := os.Getenv("ADDRESS")

	rateLimiter := lib.NewRateLimiter(30*time.Second, 100) // Allow 5 requests per minute

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./assets/styles/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./assets/img/"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	http.Handle("/", rateLimiter.Wrap(handler.Index))
	http.Handle("/sign-up", rateLimiter.Wrap(auth.SignUp))
	http.Handle("/sign-up-page", rateLimiter.Wrap(auth.SignUpPage))
	http.Handle("/google-sign-in", rateLimiter.Wrap(auth.HandleGoogleLogin))
	http.Handle("/github-sign-in", rateLimiter.Wrap(auth.HandleGithubLoginHandler))
	http.Handle("/callback", rateLimiter.Wrap(auth.HandleCallback))
	http.Handle("/github-callback", rateLimiter.Wrap(auth.HandleGithubCallback))
	http.Handle("/sign-in", rateLimiter.Wrap(auth.SignIn))
	http.Handle("/sign-in-page", rateLimiter.Wrap(auth.SignInPage))
	http.Handle("/logout", rateLimiter.Wrap(auth.Logout))

	http.Handle("/profile", rateLimiter.Wrap(handler.ProfilePage))
	http.Handle("/edit-user", rateLimiter.Wrap(handler.EditUser))
	http.Handle("/edit-user-page", rateLimiter.Wrap(handler.EditUserPage))

	http.Handle("/trending", rateLimiter.Wrap(handler.ListPost))
	http.Handle("/post", rateLimiter.Wrap(handler.CreatePost))
	http.Handle("/delete-post/", rateLimiter.Wrap(handler.DeletePost))
	http.Handle("/edit-post-page/", rateLimiter.Wrap(handler.EditPostPage))
	http.Handle("/edit-post/", rateLimiter.Wrap(handler.EditPost))
	http.Handle("/comment/", rateLimiter.Wrap(handler.Comment))
	http.Handle("/posts/", rateLimiter.Wrap(handler.GetPost))
	http.Handle("/user/", rateLimiter.Wrap(handler.UserProfilePage))
	http.Handle("/like/", rateLimiter.Wrap(handler.LikePost))
	http.Handle("/dislike/", rateLimiter.Wrap(handler.DislikePost))
	http.Handle("/like-comment/", rateLimiter.Wrap(handler.LikeComment))
	http.Handle("/dislike-comment/", rateLimiter.Wrap(handler.DislikeComment))
	http.Handle("/category/", rateLimiter.Wrap(handler.GetPostOfCategory))

	httpsServer := http.Server{
		Addr: PORT,
	}

	go models.DeleteExpiredSessions()

	go func() {
		err := http.ListenAndServe(":8080", lib.RedirectToHTTPS(http.DefaultServeMux))
		if err != nil {
			panic(err)
		}
	}()

	log.Print("Server started and running on ")
	log.Println(ADDRESS + PORT)
	if err := httpsServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
