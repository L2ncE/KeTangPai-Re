package service

import (
	"ketangpai/dao"
	"ketangpai/model"
	"time"
)

// AddTopic 添加话题
func AddTopic(topic model.Topic) error {
	err := dao.InsertTopic(topic)
	return err
}

// DeleteTopic 删除话题
func DeleteTopic(topicId int) error {
	err := dao.DeleteTopic(topicId)
	return err
}

// GetTopics 得到话题
func GetTopics() ([]model.Topic, error) {
	return dao.SelectTopic()
}

// GetTopicById 通过id得到话题
func GetTopicById(topicId int) (model.Topic, error) {
	return dao.SelectTopicById(topicId)
}

// GetNameById 通过id拿到用户名
func GetNameById(topicId int) (string, error) {
	return dao.SelectNameById(topicId)
}

// ChangeTopic 改变话题
func ChangeTopic(id int, context string, UpdateTime time.Time) error {
	err := dao.ChangeTopic(id, context, UpdateTime)
	return err
}

// TopicLikes 话题点赞
func TopicLikes(topicId int) error {
	err := dao.TopicLikes(topicId)
	return err
}
