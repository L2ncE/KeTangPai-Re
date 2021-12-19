package service

import (
	"ketangpai/dao"
	"ketangpai/model"
)

// AddHomework 添加作业
func AddHomework(homework model.Homework) error {
	return dao.InsertHomework(homework)
}

// DeleteHomework 删除作业
func DeleteHomework(Id int) error {
	err := dao.DeleteHomework(Id)
	return err
}

// GetNameById4 通过id得到用户名
func GetNameById4(homeworkId int) (string, error) {
	return dao.SelectNameById4(homeworkId)
}
