package main

import (
	"fmt"
	"forum/handlers/auth"
	"forum/handlers/user"
	"log"
	"net/http"
)

func main() {
	//DBManipulation()

	http.HandleFunc("/Signup", auth.SignupHandler)
	http.HandleFunc("/Signin", auth.SigninHandler)

	http.HandleFunc("/post", user.Post)
	http.HandleFunc("/comment", user.Comment)
	http.HandleFunc("/posts", user.AllPosts)

	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
