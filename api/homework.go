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

// addHomework 添加作业
func addHomework(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	PosterName := iUsername.(string)

	context := ctx.PostForm("context")
	name := ctx.PostForm("name")
	classroomIdString := ctx.PostForm("classroom_id") //布置作业的课堂id
	classroomId, err := strconv.Atoi(classroomIdString)
	if err != nil {
		fmt.Println("classroom id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "课堂id有误")
		return
	}

	homework := model.Homework{
		ClassRoomId: classroomId,
		Context:     context,
		Name:        name,
		PosterName:  PosterName,
		PostTime:    time.Now(),
	}
	err = service.AddHomework(homework)
	if err != nil {
		fmt.Println("add homework err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// deleteHomework 删除作业
func deleteHomework(ctx *gin.Context) {
	homeworkIdString := ctx.Param("homework_id") //输入评论id
	homeworkId, err := strconv.Atoi(homeworkIdString)
	homeworkNameString, _ := ctx.Get("username") //取用户名
	nameString, _ := service.GetNameById2(homeworkId)
	//不能删除他人的评论,将用户名进行判断
	if homeworkNameString == nameString {
		if err != nil {
			fmt.Println("homework id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "homework_id格式有误")
			return
		}
		err = service.DeleteHomework(homeworkId)
		if err != nil {
			fmt.Println("delete homework err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "不能删除他人布置作业")
	}
}

// deleteHomework0 管理员删除作业 不必再进行用户名判断
func deleteHomework0(ctx *gin.Context) {
	homeworkIdString := ctx.Param("homework_id") //输入评论id
	homeworkId, err := strconv.Atoi(homeworkIdString)
	if err != nil {
		fmt.Println("homework id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "homework_id格式有误")
		return
	}
	err = service.DeleteHomework(homeworkId)
	if err != nil {
		fmt.Println("delete homework err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
