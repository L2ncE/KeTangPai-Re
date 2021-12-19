package model

import "time"

type Homework struct {
	Id            int
	Name          string
	Context       string
	PublisherName string
	PublishTime   time.Time
}
