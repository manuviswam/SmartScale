package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"database/sql"
	
	h "github.com/manuviswam/SmartScale/handlers"
	c "github.com/manuviswam/SmartScale/config"
	"github.com/manuviswam/SmartScale/utils"

	"github.com/gorilla/mux"
	gh "github.com/gorilla/handlers"
	_ "github.com/lib/pq"
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

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=smart_scale sslmode=disable",
        conf.DBUsername, conf.DBPassword)
    db, err := sql.Open("postgres", dbinfo)
    defer db.Close()
    if err != nil {
    	log.Fatal(err)
    }

    employeeGetter := utils.EmployeeDetailFromServer{
    	Server: conf.AdjuvantServer,
    	API: conf.AdjuvantInternalNumberAPI,
    	AuthKey: conf.AdjuvantAuthKey,
    }

	r := mux.NewRouter()
    r.HandleFunc("/", h.ServeIndexPage())
    r.HandleFunc("/api/weight", h.SaveWeight(db, &employeeGetter))
    r.HandleFunc("/api/getWeight", h.GetWeight())
    r.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",conf.Port), gh.LoggingHandler(f, r)))
}