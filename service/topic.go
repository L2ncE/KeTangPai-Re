package service

import (
	"ketangpai/dao"
	"ketangpai/model"
	"time"
)

func AddTopic(topic model.Topic) error {
	err := dao.InsertTopic(topic)
	return err
}

func DeleteTopic(topicId int) error {
	err := dao.DeleteTopic(topicId)
	return err
}

func GetTopics() ([]model.Topic, error) {
	return dao.SelectTopic()
}

func GetTopicById(topicId int) (model.Topic, error) {
	return dao.SelectTopicById(topicId)
}

func GetNameById(topicId int) (string, error) {
	return dao.SelectNameById(topicId)
}

func ChangeTopic(id int, context string, UpdateTime time.Time) error {
	err := dao.ChangeTopic(id, context, UpdateTime)
	return err
}

func TopicLikes(topicId int) error {
	err := dao.TopicLikes(topicId)
	return err
}
