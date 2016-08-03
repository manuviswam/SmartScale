package handlers

import (
	"fmt"
	"net/http"
)

func ServeIndexPage() func(http.ResponseWriter, *http.Request) { 
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, Have you checked your weight today?")
	}
}