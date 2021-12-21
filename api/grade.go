package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"ketangpai/service"
	"ketangpai/tool"
	"strconv"
	"time"
)

// addGrade 添加成绩
func addGrade(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //得到用户名
	posterName := iUsername.(string)
	Name := ctx.PostForm("name")
	Subject := ctx.PostForm("subject")
	SGrade := ctx.PostForm("grade")
	Grade, _ := strconv.Atoi(SGrade)
	grade := model.Grade{
		Name:       Name,
		Subject:    Subject,
		Grade:      Grade,
		Poster:     posterName,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}

	err := service.AddGrade(grade)
	if err != nil {
		fmt.Println("add grade err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// changeGrade 更改成绩
func changeGrade(ctx *gin.Context) {
	SNewGrade := ctx.PostForm("newGrade")
	newGrade, _ := strconv.Atoi(SNewGrade)
	iGradeId := ctx.Param("grade_id")
	gradeId, err := strconv.Atoi(iGradeId)
	UpdateTime := time.Now()
	err = service.ChangeGrade(gradeId, newGrade, UpdateTime)
	if err != nil {
		fmt.Println("change grade err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
