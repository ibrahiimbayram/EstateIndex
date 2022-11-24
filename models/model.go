package models

import "time"

type EstateIndex struct {
	Name           string    `json:"Name"`
	Date           time.Time `json:"Date"`
	TurnaroundTime int       `json:"TurnaroundTime"`
	ForSale        struct {
		RegionAverage   float64 `json:"RegionAverage"`
		MinPrice        float64 `json:"MinPrice"`
		MaxPirce        float64 `json:"MaxPirce"`
		OneMonthChange  float64 `json:"OneMonthChange"`
		ThreeYearChange float64 `json:"ThreeYearChange"`
		FiveYearChange  float64 `json:"FiveYearChange"`
	}
	ForRent struct {
		RegionAverage   float64 `json:"RegionAverage"`
		MinPrice        float64 `json:"MinPrice"`
		MaxPirce        float64 `json:"MaxPirce"`
		OneMonthChange  float64 `json:"OneMonthChange"`
		ThreeYearChange float64 `json:"ThreeYearChange"`
		FiveYearChange  float64 `json:"FiveYearChange"`
	}
}
