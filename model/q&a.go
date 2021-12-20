package model

import "time"

type QuestionAndAnswer struct {
	Name     string    `json:"name"`
	Context  string    `json:"context"`
	PostTime time.Time `json:"post_time"`
}
