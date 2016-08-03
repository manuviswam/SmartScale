package main

import (
	"net/http"
	h "github.com/manuviswam/SmartScale/handlers"
)

func main() {
	http.HandleFunc("/", h.ServeIndexPage)
    http.ListenAndServe(":8080", nil)
}