package main

import (
	"fmt"
<<<<<<< HEAD
	"forum/handlers/auth"
=======
	"forum/handlers/user"
>>>>>>> a9177fda9e3f23b174d76b3c19ecef1552ec19a4
	"log"
	"net/http"
)

func main() {
<<<<<<< HEAD
	//DBManipulation()

	http.HandleFunc("/Signup", auth.SignupHandler)
	http.HandleFunc("/Signin", auth.SigninHandler)

=======

	http.HandleFunc("/post", user.Post)
	http.HandleFunc("/comment", user.Comment)
>>>>>>> a9177fda9e3f23b174d76b3c19ecef1552ec19a4

	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
<<<<<<< HEAD
}
=======

}
>>>>>>> a9177fda9e3f23b174d76b3c19ecef1552ec19a4
