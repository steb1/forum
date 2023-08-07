package main

import (
	"fmt"
	"forum/handler"
	"forum/handlers/auth"
	"forum/handlers/user"
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

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/Signup", auth.SignupHandler)
	http.HandleFunc("/Signin", auth.SigninHandler)

	http.HandleFunc("/post", user.Post)
	http.HandleFunc("/comment", user.Comment)
	http.HandleFunc("/posts", user.AllPosts)

	log.Println("Server started and running on", PORT)
	log.Println(ADDRESS + PORT)

	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
