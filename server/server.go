package server

import (
	"fmt"
	"log"
	"net/http"
)

var Routes Urls = Urls{
	Home:     "/",
	AsciiArt: "/ascii-art",
}

func Run() {
	http.HandleFunc(Routes.Home, ViewHandler)
	http.HandleFunc(Routes.AsciiArt, CreateHandler)
	http.HandleFunc("/authors", AuthorsHandler)
	http.HandleFunc("/about", AboutHandler)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public/"))))
	fmt.Println("Listened on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err) // very bad
}
