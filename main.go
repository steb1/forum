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

	http.HandleFunc("/", handler.Index)

	log.Println("Server started and running on", PORT)
	log.Println(ADDRESS + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
