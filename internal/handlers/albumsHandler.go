package handlers

import (
	"net/http"
	"spoty/internal/models"
	"text/template"
)

func AlbumsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/albums" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/album.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	artist := r.FormValue("artist")

	albums, err := models.GetArtistsTopAlbums(artist)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, albums)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
