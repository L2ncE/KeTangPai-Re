package model

import "time"

type Homework struct {
	Id          int
	ClassRoomId int
	Name        string
	Context     string
	PosterName  string
	PostTime    time.Time
}
