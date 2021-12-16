package service

import (
	"database/sql"
	"ketangpai/dao"
	"ketangpai/model"
)

// ChangePassword 修改密码服务
func ChangePassword(username, newPassword string) error {
	err := dao.UpdatePassword(username, newPassword)
	return err
}

// IsPasswordCorrect 判断密码是否正确服务
func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	return true, nil
}

// IsRepeatUsername 判断用户名是否重复
func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// Register 注册服务
func Register(user model.User) error {
	err := dao.Insert(user)
	return err
}

// SelectAnswerByUsername 通过昵称查找答案服务
func SelectAnswerByUsername(username string) string {
	answer := dao.SelectAnswerByUsername(username)
	return answer
}
