package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

func ChangePassword(username, newPassword string) error {
	err := dao.UpdatePassword(username, newPassword)
	return err
}

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

func Register(user model.User) error {
	err := dao.Insert(user)
	return err
}

func SelectAnswerByUsername(username string) string {
	answer := dao.SelectAnswerByUsername(username)
	return answer
}
