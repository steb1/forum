package main

import (
	"fmt"
	"forum/handlers/user"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/post", user.Post)
	http.HandleFunc("/comment", user.Comment)

	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
