package handlers

import (
	"log"
	"net/http"
	"html/template"
)

func ServeIndexPage() func(http.ResponseWriter, *http.Request) { 
	return func(w http.ResponseWriter, r *http.Request) {
		home, err := template.ParseFiles("public/js/index.html")
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		home.Execute(w, r.Host)
	}
}