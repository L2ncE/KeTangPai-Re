package dao

import (
	"fmt"
	"ketangpai/model"
)

// InsertHomework 插入作业
func InsertHomework(homework model.Homework) error {

	sqlStr := "insert into homework(Name,ClassRoomId,Context,PostTime,PosterName)values (? ,? ,? ,? ,? )"
	_, err := dB.Exec(sqlStr, homework.Name, homework.ClassRoomId, homework.Context, homework.PostTime, homework.PosterName)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// DeleteHomework 删除作业
func DeleteHomework(Id int) error {
	sqlStr := `delete from homework where Id=?`
	_, err := dB.Exec(sqlStr, Id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}

// SelectNameById4 通过id找到用户名
func SelectNameById4(homeworkId int) (string, error) {
	var homework model.Homework

	row := dB.QueryRow("SELECT PosterName FROM homework WHERE id = ? ", homeworkId)
	if row.Err() != nil {
		return homework.PosterName, row.Err()
	}

	err := row.Scan(&homework.PosterName)
	if err != nil {
		return homework.PosterName, err
	}

	return homework.PosterName, nil
}
