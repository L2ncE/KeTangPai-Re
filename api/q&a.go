package api

import (
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"ketangpai/tool"
	"net/http"
	"time"
)

func question(ctx *gin.Context) {
	question := ctx.PostForm("question")
	IUsername, _ := ctx.Get("username")
	username := IUsername.(string)
	questionAndAnswerDetail := model.QuestionAndAnswer{
		Name:     username,
		Context:  question,
		PostTime: time.Now(),
	}
	tool.RespSuccessfulWithDate(ctx, questionAndAnswerDetail)
}

func answer(ctx *gin.Context) {
	answer := ctx.PostForm("answer")
	tool.RespSuccessful(ctx)
	ctx.JSON(http.StatusOK, gin.H{

		"answer": answer,
	})
}
