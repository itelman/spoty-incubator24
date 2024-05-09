package main

import (
	"fmt"
	"log"
	"net/http"
	"spoty/internal/handlers"
)

const API_KEY = "6c8e19bc52595c3544a0e449c8b9e696"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/albums", handlers.AlbumsHandler)
	mux.HandleFunc("/albuminfo", handlers.AlbumInfoHandler)
	mux.HandleFunc("/search", handlers.SearchHandler)
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	fmt.Println("Server is running on http://localhost:8080/...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
