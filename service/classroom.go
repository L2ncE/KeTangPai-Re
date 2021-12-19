package service

import (
	"ketangpai/dao"
	"ketangpai/model"
)

// AddClassRoom 添加教室
func AddClassRoom(classroom model.ClassRoom) error {
	err := dao.InsertClassRoom(classroom)
	return err
}

// DeleteClassRoom 删除教室
func DeleteClassRoom(classroomId int) error {
	err := dao.DeleteClassRoom(classroomId)
	return err
}

// GetNameById3 通过id拿到用户名
func GetNameById3(classroomId int) (string, error) {
	return dao.SelectNameById3(classroomId)
}

// OpenClassRoom 打开教室
func OpenClassRoom(status bool, classroomId int) error {
	err := dao.OpenClassRoom(status, classroomId)
	return err
}

// CloseClassRoom 关闭教室
func CloseClassRoom(status bool, classroomId int) error {
	err := dao.CloseClassRoom(status, classroomId)
	return err
}
