package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
)

func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	iUsername, _ := ctx.Get("username")
	l1 := len([]rune(newPassword))
	if l1 <= 16 {
		username := iUsername.(string)

		//检验旧密码是否正确
		flag, err := service.IsPasswordCorrect(username, oldPassword)
		if err != nil {
			fmt.Println("judge password correct err: ", err)
			tool.RespInternalError(ctx)
			return
		}

		if !flag {
			tool.RespErrorWithDate(ctx, "旧密码错误")
			return
		}

		//修改新密码
		err = service.ChangePassword(username, newPassword)
		if err != nil {
			fmt.Println("change password err: ", err)
			tool.RespInternalError(ctx)
			return
		}

		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "密码请在16位之内")
		return
	}
}

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithDate(ctx, "登陆失败,请输入密保")
		answer := ctx.PostForm("answer")
		if answer == service.SelectAnswerByUsername(username) {
			ctx.JSON(200, gin.H{
				"msg": "密保正确,请重新输入密码",
			})
			newPassword := ctx.PostForm("new_password")
			err = service.ChangePassword(username, newPassword)
			if err != nil {
				fmt.Println("change password err: ", err)
				tool.RespInternalError(ctx)
				return
			}
		} else {
			tool.RespErrorWithDate(ctx, "密保错误")
			return
		}
		return
	}

	ctx.SetCookie("username", username, 600, "/", "", false, false)
	tool.RespSuccessful(ctx)
}

func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	question := ctx.PostForm("question")
	answer := ctx.PostForm("answer")

	l1 := len([]rune(username))
	l2 := len([]rune(password))
	if l1 <= 8 {
		if l2 <= 16 {
			user := model.User{
				Name:     username,
				Password: password,
				Question: question,
				Answer:   answer,
			}

			flag, err := service.IsRepeatUsername(username)
			if err != nil {
				fmt.Println("judge repeat username err: ", err)
				tool.RespInternalError(ctx)
				return
			}

			if flag {
				tool.RespErrorWithDate(ctx, "用户名已经存在")
				return
			}

			err = service.Register(user)
			if err != nil {
				fmt.Println("register err: ", err)
				tool.RespInternalError(ctx)
				return
			}

			tool.RespSuccessful(ctx)
		} else {
			tool.RespErrorWithDate(ctx, "密码请在16位之内")
			return
		}
	} else {
		tool.RespErrorWithDate(ctx, "用户名请在8位之内")
		return
	}
}
