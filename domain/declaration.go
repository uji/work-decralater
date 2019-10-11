package domain

import "time"

type Declaration struct {
	UserID    int       `json:"user_id" dynamo:"UserID"`
	Date      string    `json:"date" dynamo:"Date"`
	StartAt   string    `json:"start_at" dynamo:"StartAt"`
	EndAt     string    `json:"end_at" dynamo:"EndAt"`
	Breaktime string    `json:"breaktime" dynamo:"Breaktime"`
	Place     string    `json:"place" dynamo:"Place"`
	Comment   string    `json:"comment" dynamo:"Comment"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
}
