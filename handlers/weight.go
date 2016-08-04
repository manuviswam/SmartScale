package handlers

import (
	"net/http"
	"database/sql"
	"time"
	"log"
	"strconv"

	m "github.com/manuviswam/SmartScale/models"
	"github.com/manuviswam/SmartScale/utils"
)

func SaveWeight(db *sql.DB, eg utils.EmployeeGetter) func(http.ResponseWriter, *http.Request) { 
	return func(w http.ResponseWriter, r *http.Request) {
		wt, err := strconv.ParseFloat(r.FormValue("weight"), 64)
		if err != nil {
			log.Println("Invalid weight: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		in, err := strconv.ParseInt(r.FormValue("internalNumber"), 10, 64)
		if err != nil {
			log.Println("Invalid internal number: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		emp, err := eg.GetEmployeeFromInternalNumber(in)
		if err != nil {
			log.Println("Error decoding response: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		weightInfo := m.WeightInfo{
			EmpId : emp.EmpId,
			Weight : wt,
			RecordedAt : time.Now(),
		}

		err = weightInfo.SaveToDB(db)
		if err != nil {
			log.Println("Error writing to DB: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}