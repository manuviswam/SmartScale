package handlers

import (
	"fmt"
	"net/http"
)

func ServeIndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Have you checked your weight today?")
}