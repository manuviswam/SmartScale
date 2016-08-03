package main

import (
	"net/http"
	"log"
	"os"
	
	h "github.com/manuviswam/SmartScale/handlers"

	"github.com/gorilla/mux"
	gh "github.com/gorilla/handlers"
)

func main() {
	r := mux.NewRouter()
    r.HandleFunc("/", h.ServeIndexPage())
    log.Fatal(http.ListenAndServe(":8080", gh.LoggingHandler(os.Stdout, r)))
}