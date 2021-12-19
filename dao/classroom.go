package dao

import (
	"fmt"
	"ketangpai/model"
)

// InsertClassRoom 创建教室
func InsertClassRoom(classroom model.ClassRoom) error {
	_, err := dB.Exec("INSERT INTO ClassRoom(ClassName, CreatorName, CreateTime, LastOpenTime,Status) "+"values(?, ?, ?, ?, ?);", classroom.ClassName, classroom.CreatorName, classroom.CreateTime, classroom.LastOpenTime, classroom.Status)
	return err
}

// SelectNameById3 通过id查找创建用户
func SelectNameById3(classroomId int) (string, error) {
	var classroom model.ClassRoom

	row := dB.QueryRow("SELECT CreatorName FROM classroom WHERE id = ? ", classroomId)
	if row.Err() != nil {
		return classroom.CreatorName, row.Err()
	}

	err := row.Scan(&classroom.CreatorName)
	if err != nil {
		return classroom.CreatorName, err
	}

	return classroom.CreatorName, nil
}

// DeleteClassRoom 删除教室
func DeleteClassRoom(classroomId int) error {

	sqlStr := `delete from classroom where Id=?`
	_, err := dB.Exec(sqlStr, classroomId)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}

// OpenClassRoom 打开教室
func OpenClassRoom(status bool, id int) error {
	sqlStr := `update classroom set Status=? where id = ?`
	_, err := dB.Exec(sqlStr, status, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

// CloseClassRoom 关闭教室
func CloseClassRoom(status bool, id int) error {
	sqlStr := `update classroom set Status=? where id = ?`
	_, err := dB.Exec(sqlStr, status, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
