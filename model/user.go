package model

type User struct {
	Id              int
	Name            string
	Password        string
	Question        string
	Answer          string
	ClassroomIdSign int
	SpeechNum       int
	Status          string
}

type UserRank struct {
	Name      string `json:"name"`
	SpeechNum int    `json:"speechNum"`
}
