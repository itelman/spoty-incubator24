package handlers

import (
	"net/http"
	"spoty/internal/models"
	"text/template"
)

func AlbumInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/albuminfo" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/info.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	artist := r.FormValue("artist")
	album := r.FormValue("album")

	info, err := models.GetAlbumInfo(artist, album)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, info)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
