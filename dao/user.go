package dao

import (
	"fmt"
	"ketangpai/model"
)

// UpdatePassword 更新密码操作
func UpdatePassword(Name string, newPassword string) error {
	sqlStr := `update user set Password=? where Name = ?`
	_, err := dB.Exec(sqlStr, newPassword, Name)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

// SelectUserByUsername 查找用户
func SelectUserByUsername(Name string) (model.User, error) {
	user := model.User{}

	row := dB.QueryRow("SELECT id, password FROM user WHERE Name = ? ", Name)
	if row.Err() != nil {
		return user, row.Err()
	}

	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Insert 注册时插入数据
func Insert(user model.User) error {

	sqlStr := "insert into user(Name,Password,Question,Answer,Status)values (?,?,?,?,?)"
	_, err := dB.Exec(sqlStr, user.Name, user.Password, user.Question, user.Answer, user.Status)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// SelectAnswerByUsername 通过用户名来找到密保的答案
func SelectAnswerByUsername(Name string) string {
	user := model.User{}
	sqlStr := `select answer from user where name=?;`
	dB.QueryRow(sqlStr, Name).Scan(&user.Answer)
	return user.Answer
}

// SelectStatusByName 通过用户名查身份
func SelectStatusByName(Name string) string {
	user := model.User{}
	sqlStr := `select Status from user where Name=?;`
	dB.QueryRow(sqlStr, Name).Scan(&user.Status)
	return user.Status
}

// SelectIdByName 通过用户名查Id
func SelectIdByName(Name string) int {
	user := model.User{}
	sqlStr := `select Id from user where Name=?;`
	dB.QueryRow(sqlStr, Name).Scan(&user.Id)
	return user.Id
}
