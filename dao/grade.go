package dao

import (
	"fmt"
	"ketangpai/model"
	"time"
)

// InsertGrade 插入成绩
func InsertGrade(grade model.Grade) error {

	sqlStr := "insert into grade(Name,Subject,Grade,Poster,PostTime,UpdateTime)values (?,?,?,?,?,?)"
	_, err := dB.Exec(sqlStr, grade.Name, grade.Subject, grade.Grade, grade.Poster, grade.PostTime, grade.UpdateTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// ChangeGrade 修改成绩
func ChangeGrade(id int, newGrade int, time time.Time) error {
	sqlStr := `update grade set Grade=?,UpdateTime=? where id = ?`
	_, err := dB.Exec(sqlStr, newGrade, time, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
