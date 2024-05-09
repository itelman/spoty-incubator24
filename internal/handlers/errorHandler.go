package handlers

import (
	"log"
	"net/http"
	"spoty/internal/models"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, statusCode int) {
	errTmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	errStr := models.ErrorConstructor(statusCode)
	err = errTmpl.Execute(w, errStr)
	log.Printf("\"[CODE %d]: %s\"\n", errStr.Code, errStr.Name)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
