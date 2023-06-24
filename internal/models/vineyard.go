package models

import "time"

type Vineyard struct {
	ID             int          `json:"id"`
	XCoord         float64      `json:"xCoord"`
	YCoord         float64      `json:"yCoord"`
	WorkHistory    []WorkObject `json:"workHistory"`
	WeatherHistory string       `json:"weatherHistory"`
}

type WorkObject struct {
	Date          time.Time
	WorkType      string
	VineYardGrade int
}

type WeatherObject struct {
	Sunny int
	Wet   int
}
