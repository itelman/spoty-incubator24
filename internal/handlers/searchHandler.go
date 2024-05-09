package handlers

import (
	"net/http"
	"spoty/internal/models"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	search := r.FormValue("q")

	results, err := models.GetAlbumResults(search)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, results)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
