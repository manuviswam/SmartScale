package handlers

import (
	"net/http"
	"database/sql"
	"time"
	"log"
	"strconv"
	"encoding/json"

	"github.com/gorilla/websocket"

	m "github.com/manuviswam/SmartScale/models"
	"github.com/manuviswam/SmartScale/utils"
)

func SaveWeight(db *sql.DB, eg utils.EmployeeGetter) func(http.ResponseWriter, *http.Request) { 
	return func(w http.ResponseWriter, r *http.Request) {
		wt, err := strconv.ParseFloat(r.FormValue("weight"), 64)
		if err != nil {
			log.Println("Invalid weight: ", err)
			w.WriteHeader(http.StatusBadRequest)
			pushErrorMessage("Something bad happened. Please try again.")
			return
		}

		in, err := strconv.ParseInt(r.FormValue("internalNumber"), 10, 64)
		if err != nil {
			log.Println("Invalid internal number: ", err)
			w.WriteHeader(http.StatusBadRequest)
			pushErrorMessage("Something bad happened. Please try again.")
			return
		}

		emp, err := eg.GetEmployeeFromInternalNumber(in)
		if err != nil {
			log.Println("Error decoding response: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			pushErrorMessage("Unable to retrieve your details. Please contact admin.")
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
			pushErrorMessage("Something is wrong with the server. Please try after sometime.")
			return
		}

		weights, _ := m.GetWeightsByEmpId(db, emp.EmpId)

		w.WriteHeader(http.StatusOK)
		pushSuccessMessage(emp, wt, weights)
	}
}

func pushMessage(data m.WeightResponse) {
	pushMsg, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	weightChan := m.WeightChan()
	weightChan <- pushMsg
}

func pushErrorMessage(msg string) {
	wr := m.WeightResponse{
		IsError: true,
		ErrorMsg: msg,
	}

	pushMessage(wr)
}

func pushSuccessMessage(emp m.Employee, currentWeight float64, weights []m.Weight) {
	wr := m.WeightResponse{
		IsError: false,
		EmpId: emp.EmpId,
		EmpName: emp.EmployeeName,
		CurrentWeight: currentWeight,
		Weights: weights,
	}

	pushMessage(wr)
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func GetWeight() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
	    if err != nil {
	        log.Println(err)
	        return
	    }
	    defer conn.Close()
	    weightChan := m.WeightChan()
	    for {
			msg := <-weightChan
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
	    }
	}
}
