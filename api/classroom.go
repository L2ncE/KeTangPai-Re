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

// addClassRoom 新增课堂
func addClassRoom(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //得到用户名
	creatorName := iUsername.(string)

	classname := ctx.PostForm("classname")

	classroom := model.ClassRoom{
		ClassName:    classname,
		CreatorName:  creatorName,
		CreateTime:   time.Now(),
		LastOpenTime: time.Now(),
		Status:       true,
	}

	err := service.AddClassRoom(classroom)
	if err != nil {
		fmt.Println("add classroom err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// deleteClassRoom 删除课堂
func deleteClassRoom(ctx *gin.Context) {
	classroomIdString := ctx.Param("classroom_id")
	classroomId, err := strconv.Atoi(classroomIdString)
	classroomNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameById3(classroomId)
	//必须用户名相同,无法删除他人课堂
	if classroomNameString == nameString {
		if err != nil {
			fmt.Println("classroom id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "classroom_id格式有误")
			return
		}
		err = service.DeleteClassRoom(classroomId)
		if err != nil {
			fmt.Println("delete classroom err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "不能删除他人课堂")
	}
}

// deleteClassRoom0 管理员删除课堂
func deleteClassRoom0(ctx *gin.Context) {
	classroomIdString := ctx.Param("classroom_id")
	classroomId, err := strconv.Atoi(classroomIdString)
	//无需判断用户名
	if err != nil {
		fmt.Println("classroom id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "classroom_id格式有误")
		return
	}
	err = service.DeleteClassRoom(classroomId)
	if err != nil {
		fmt.Println("delete classroom err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

// openClassRoom 开启课堂
func openClassRoom(ctx *gin.Context) {
	classroomIdString := ctx.Param("classroom_id")
	classroomId, err := strconv.Atoi(classroomIdString)
	classroomNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameById(classroomId)
	//必须用户名相同,无法开启他人课堂
	if classroomNameString == nameString {
		if err != nil {
			fmt.Println("classroom id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "classroom_id格式有误")
			return
		}
		err = service.OpenClassRoom(true, classroomId)
		if err != nil {
			fmt.Println("open classroom err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "不能开启他人课堂")
	}
}

// closeClassRoom 关闭课堂
func closeClassRoom(ctx *gin.Context) {
	classroomIdString := ctx.Param("classroom_id")
	classroomId, err := strconv.Atoi(classroomIdString)
	classroomNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameById(classroomId)
	//必须用户名相同,无法关闭他人课堂
	if classroomNameString == nameString {
		if err != nil {
			fmt.Println("classroom id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "classroom_id格式有误")
			return
		}
		err = service.CloseClassRoom(false, classroomId)
		if err != nil {
			fmt.Println("close classroom err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "不能关闭他人课堂")
	}
}

// signInClassRoom 课堂签到
func signInClassRoom(ctx *gin.Context) {
	classroomIdString := ctx.Param("classroom_id")
	classroomId, err := strconv.Atoi(classroomIdString)
	if err != nil {
		fmt.Println("classroom id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "classroom_id格式有误")
		return
	}
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	err = service.SignInClassroom(classroomId, username)
	if err != nil {
		fmt.Println("sign classroom err: ", err)
		tool.RespInternalError(ctx)
		return
	}
}
