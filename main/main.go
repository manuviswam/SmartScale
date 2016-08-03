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
	f, err := os.OpenFile("SmartScale.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	r := mux.NewRouter()
    r.HandleFunc("/", h.ServeIndexPage())
    log.Fatal(http.ListenAndServe(":8080", gh.LoggingHandler(f, r)))
}