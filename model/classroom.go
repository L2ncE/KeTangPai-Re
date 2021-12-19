package model

import "time"

type ClassRoom struct {
	Id           int
	ClassName    int
	CreatorName  string
	CreateTime   time.Time
	LastOpenTime time.Time
	Status       bool
}
