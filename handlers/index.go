package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func ServeIndexPage() { 
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, Have you checked your weight today?")
	}
}