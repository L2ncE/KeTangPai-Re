package service

import (
	"ketangpai/dao"
	"ketangpai/model"
)

// AddComment 添加评论
func AddComment(comment model.Comment) error {
	return dao.InsertComment(comment)
}

// GetTopicComments 得到话题评论
func GetTopicComments(topicId int) ([]model.Comment, error) {
	return dao.SelectCommentByTopicId(topicId)
}

// DeleteComment 删除评论
func DeleteComment(Id int) error {
	err := dao.DeleteComment(Id)
	return err
}

// GetNameById2 通过id得到用户名
func GetNameById2(commentId int) (string, error) {
	return dao.SelectNameById2(commentId)
}

// CommentLikes 评论点赞
func CommentLikes(commentId int) error {
	err := dao.CommentLikes(commentId)
	return err
}
