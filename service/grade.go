package service

import (
	"ketangpai/dao"
	"ketangpai/model"
	"time"
)

// AddGrade 添加成绩
func AddGrade(grade model.Grade) error {
	err := dao.InsertGrade(grade)
	return err
}

// ChangeGrade 修改成绩
func ChangeGrade(id int, newGrade int, time time.Time) error {
	err := dao.ChangeGrade(id, newGrade, time)
	return err
}
