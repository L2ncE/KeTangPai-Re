package model

import "time"

type Topic struct {
	Id         int       `json:"id"`
	Context    string    `json:"context"`
	Name       string    `json:"name"`
	PostTime   time.Time `json:"post_time"`
	UpdateTime time.Time `json:"update_time"`
	Likes      int       `json:"likes"`
}

type TopicDetail struct {
	Topic
	Comments []Comment
}
