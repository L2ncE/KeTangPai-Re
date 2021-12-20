package model

import "time"

type Grade struct {
	Id         int
	Name       string
	Subject    string
	Grade      int
	Poster     string
	PostTime   time.Time
	UpdateTime time.Time
}
