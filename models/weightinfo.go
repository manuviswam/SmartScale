package models

import (
	"database/sql"
	"time"
)

const (
	SaveWeightQuery = "INSERT INTO weights(empid,weight,recorded_at) VALUES($1,$2,$3)"
	GetWeightsQuery = "WITH t AS (SELECT weight,recorded_at FROM weights WHERE empid=$1 ORDER BY recorded_at DESC LIMIT 10) SELECT * FROM t ORDER BY recorded_at ASC"
)

type WeightInfo struct {
	Id int64
	EmpId string
	Weight float64
	RecordedAt time.Time
}

type Weight struct {
	Weight float64
	RecordedAt time.Time
}

func (w *WeightInfo)SaveToDB(db *sql.DB) error {
	_,err := db.Exec(SaveWeightQuery, w.EmpId, w.Weight, w.RecordedAt)
	return err
}

func GetWeightsByEmpId(db *sql.DB, empId string) ([]Weight, error) {
	weights := make([]Weight, 0)
	rows, err := db.Query(GetWeightsQuery, empId)
	defer rows.Close()
	if err != nil {
		return weights, err
	}
	for rows.Next() {
		w := Weight{}
		err := rows.Scan(&w.Weight, &w.RecordedAt)
		if err != nil {
			return weights, err
		}
		weights = append(weights, w)
	}
	return weights, nil
}