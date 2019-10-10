package domain

import "time"

type Declaration struct {
	UserID int `json:"UserID" dynamo:"UserID"`
	// Date      time.Time `dynamo:"Date" json:"Date"`
	// StartAt   time.Time `dynamo:"StartAt" json:"StartAt"`
	// EndAt     time.Time `dynamo:"EndAt" json:"EndAt"`
	// Breaktime time.Time `dynamo:"Breaktime" json:"Breaktime"`
	Place     string    `dynamo:"Place" json:"Place"`
	Comment   string    `dynamo:"Comment" json:"Comment"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
}
