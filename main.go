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
	http.HandleFunc("/Google-Sign-up", auth.HandleGoogleLogin)
	http.HandleFunc("/callback", auth.HandleCallback)
	http.HandleFunc("/sign-in", auth.SignIn)
	http.HandleFunc("/sign-in-page", auth.SignInPage)
	http.HandleFunc("/logout", auth.Logout)

	http.HandleFunc("/profile", handler.ProfilePage)
	http.HandleFunc("/edit-user", handler.EditUser)
	http.HandleFunc("/edit-user-page", handler.EditUserPage)

	http.HandleFunc("/post", handler.Post)
	http.HandleFunc("/comment", handler.Comment)
	http.HandleFunc("/posts/", handler.AllPosts)

	go models.DeleteExpiredSessions()

	log.Print("Server started and running on ")
	log.Println(ADDRESS + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
