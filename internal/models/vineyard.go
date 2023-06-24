package models

import "time"

type Vineyard struct {
	ID             int
	XCoord         float64
	YCoord         float64
	WorkHistory    []WorkObject
	WeatherHistory []WeatherObject
	GroundHistory  []GroundAnalysis
}

type WorkObject struct {
	Date          time.Time
	WorkType      string
	VineYardGrade int
}

type WeatherObject struct {
	Sunny       int
	Wet         int
	Temperature float32
}

type GroundAnalysis struct {
	Date       time.Time
	Organic    float32
	Acidity    float32
	NO3        float32
	Phosphorus float32
}
