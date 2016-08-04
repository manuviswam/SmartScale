package handlers

import (
	"fmt"
	"net/http"
	"database/sql"
	"time"
	"log"
	"strconv"

	m "github.com/manuviswam/SmartScale/models"
)

func SaveWeight(db *sql.DB) func(http.ResponseWriter, *http.Request) { 
	return func(w http.ResponseWriter, r *http.Request) {
		wt, err := strconv.ParseFloat(r.FormValue("weight"), 64)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		weightInfo := m.WeightInfo{
			EmpId : 16134,
			Weight : wt,
			RecordedAt : time.Now(),
		}

		err = weightInfo.SaveToDB(db)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Done")
	}
}