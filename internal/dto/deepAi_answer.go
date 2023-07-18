package dto

import "time"

type Answer struct {
	ID           string `json:"id"`
	Output       string `json:"output"`
	QuestionDate time.Time
	ResponseDate time.Time
}
