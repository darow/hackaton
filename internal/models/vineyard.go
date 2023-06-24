package models

import "time"

type Vineyard struct {
	ID             int
	XCoord         float64
	YCoord         float64
	WorkHistory    []WorkObject
	WeatherHistory string
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
