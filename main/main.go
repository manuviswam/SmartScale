package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	
	h "github.com/manuviswam/SmartScale/handlers"
	c "github.com/manuviswam/SmartScale/config"

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

	conf,err := c.ReadFromFile("../config.json")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
    r.HandleFunc("/", h.ServeIndexPage())
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",conf.Port), gh.LoggingHandler(f, r)))
}