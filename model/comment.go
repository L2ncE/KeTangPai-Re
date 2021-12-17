package model

import "time"

type Comment struct {
	Id          int
	TopicId     int
	Context     string
	Name        string
	CommentTime time.Time
}
