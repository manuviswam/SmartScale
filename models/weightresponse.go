package models

type WeightResponse struct{
	IsError bool
	ErrorMsg string
	EmpId string
	EmpName string
	CurrentWeight float64
	Weights []Weight
}