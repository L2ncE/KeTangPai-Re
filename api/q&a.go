package api

import (
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"ketangpai/tool"
	"time"
)

func question(ctx *gin.Context) {
	question := ctx.PostForm("question")
	IUsername, _ := ctx.Get("username")
	username := IUsername.(string)
	QuestionDetail := model.QuestionAndAnswer{
		Name:     username,
		Context:  question,
		PostTime: time.Now(),
	}
	tool.RespSuccessfulWithDate(ctx, QuestionDetail)
}

func answer(ctx *gin.Context) {
	answer := ctx.PostForm("answer")
	tool.RespSuccessful(ctx)
	IUsername, _ := ctx.Get("username")
	username := IUsername.(string)
	AnswerDetail := model.QuestionAndAnswer{
		Name:     username,
		Context:  answer,
		PostTime: time.Now(),
	}
	tool.RespSuccessfulWithDate(ctx, AnswerDetail)
}
