package main

import (
	"fmt"
	"forum/handler"
	"forum/lib"
	"log"
	"net/http"
	"os"
)

func main() {
	lib.LoadEnv(".env")
	port := os.Getenv("PORT")
	PORT := fmt.Sprintf(":%v", port)
	ADDRESS := os.Getenv("ADDRESS")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./assets/styles/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./assets/img/"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./assets/uploads/"))))

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/sign-up", auth.SignUp)
	http.HandleFunc("/sign-up-page", auth.SignUpPage)
	http.HandleFunc("/sign-in", auth.SignIn)
	http.HandleFunc("/sign-in-page", auth.SignInPage)

	// http.HandleFunc("/post", handler.Post)
	// http.HandleFunc("/comment", handler.Comment)
	// http.HandleFunc("/posts", handler.AllPosts)

	log.Println("Server started and running on", PORT)
	log.Println(ADDRESS + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
