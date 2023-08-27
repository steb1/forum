package main

import (
	"crypto/tls"
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
	http.Handle("/bookmark/", rateLimiter.Wrap(handler.Bookmark))

	http.Handle("/profile", rateLimiter.Wrap(handler.ProfilePage))
	http.Handle("/edit-user", rateLimiter.Wrap(handler.EditUser))
	http.Handle("/edit-user-page", rateLimiter.Wrap(handler.EditUserPage))

	http.Handle("/trending", rateLimiter.Wrap(handler.ListPost))
	http.Handle("/post", rateLimiter.Wrap(handler.CreatePost))
	http.Handle("/delete-post/", rateLimiter.Wrap(handler.DeletePost))
	http.Handle("/delete-comment/", rateLimiter.Wrap(handler.DeleteComment))
	http.Handle("/edit-post-page/", rateLimiter.Wrap(handler.EditPostPage))
	http.Handle("/edit-post/", rateLimiter.Wrap(handler.EditPost))
	http.Handle("/edit-comment-page/", rateLimiter.Wrap(handler.EditCommentPage))
	http.Handle("/edit-comment/", rateLimiter.Wrap(handler.EditComment))
	http.Handle("/comment/", rateLimiter.Wrap(handler.Comment))
	http.Handle("/posts/", rateLimiter.Wrap(handler.GetPost))
	http.Handle("/user/", rateLimiter.Wrap(handler.UserProfilePage))
	http.Handle("/like/", rateLimiter.Wrap(handler.LikePost))
	http.Handle("/dislike/", rateLimiter.Wrap(handler.DislikePost))
	http.Handle("/like-comment/", rateLimiter.Wrap(handler.LikeComment))
	http.Handle("/dislike-comment/", rateLimiter.Wrap(handler.DislikeComment))
	http.Handle("/category/", rateLimiter.Wrap(handler.GetPostOfCategory))
	http.Handle("/notification/", rateLimiter.Wrap(handler.GetNotifs))

	httpsServer := http.Server{
		Addr: PORT,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS12, // Minimum TLS version supported
			PreferServerCipherSuites: true,             // Prefer the server's cipher suite order
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				// Add more cipher suites as needed
			},
		},
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
	if err := httpsServer.ListenAndServeTLS(os.Getenv("CERT_PATH"), os.Getenv("KEY_PATH")); err != nil {
		log.Fatal(err)
	}
}
