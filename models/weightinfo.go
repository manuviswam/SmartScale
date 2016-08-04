package models

import (
	"database/sql"
	"time"
)

const (
	SaveWeightQuery = "INSERT INTO weights(empid,weight,recorded_at) VALUES($1,$2,$3)"
)

type WeightInfo struct {
	Id int64
	EmpId int64
	Weight float64
	RecordedAt time.Time
}

func (w *WeightInfo)SaveToDB(db *sql.DB) error {
	_,err := db.Exec(SaveWeightQuery, w.EmpId, w.Weight, w.RecordedAt)
	return err
}