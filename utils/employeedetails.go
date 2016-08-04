package utils

import (
	"fmt"
	"net/http"
	"encoding/json"

	m "github.com/manuviswam/SmartScale/models"
)

type EmployeeGetter interface{
	GetEmployeeFromInternalNumber(int64) (m.Employee, error)
}

type EmployeeDetailFromServer struct{
	Server string
	API string
	AuthKey string
}

func (e *EmployeeDetailFromServer) GetEmployeeFromInternalNumber(internalNumber int64) (m.Employee, error) {
	employee := m.Employee{}
	url := fmt.Sprintf("%s/%s/%d", e.Server, e.API, internalNumber)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", e.AuthKey)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return employee, err
	}
	
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&employee)
	if err != nil {
	  return employee, err
	}
	return employee, nil
}