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

// addComment 添加评论
func addComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	Name := iUsername.(string)

	context := ctx.PostForm("context")
	topicIdString := ctx.PostForm("topic_id") //评论的话题id
	topicId, err := strconv.Atoi(topicIdString)
	if err != nil {
		fmt.Println("topic id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "话题id有误")
		return
	}

	comment := model.Comment{
		TopicId:     topicId,
		Context:     context,
		Name:        Name,
		CommentTime: time.Now(),
	}
	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// addCommentAnonymity 匿名评论
func addCommentAnonymity(ctx *gin.Context) {
	Name := "Anonymity" //匿名评论用户名统一为Anonymity
	context := ctx.PostForm("context")
	topicIdString := ctx.PostForm("topic_id")
	topicId, err := strconv.Atoi(topicIdString)
	if err != nil {
		fmt.Println("topic id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "话题id有误")
		return
	}

	comment := model.Comment{
		TopicId:     topicId,
		Context:     context,
		Name:        Name,
		CommentTime: time.Now(),
	}
	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// deleteComment 删除评论
func deleteComment(ctx *gin.Context) {
	commentIdString := ctx.Param("comment_id") //输入评论id
	commentId, err := strconv.Atoi(commentIdString)
	commentNameString, _ := ctx.Get("username") //取用户名
	nameString, _ := service.GetNameById2(commentId)
	//不能删除他人的评论,将用户名进行判断
	if commentNameString == nameString {
		if err != nil {
			fmt.Println("comment id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "comment_id格式有误")
			return
		}
		err = service.DeleteComment(commentId)
		if err != nil {
			fmt.Println("delete comment err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "不能删除他人评论")
	}
}

// commentLikes 评论点赞
func commentLikes(ctx *gin.Context) {
	commentIdString := ctx.Param("comment_id")
	commentId, err := strconv.Atoi(commentIdString)
	if err != nil {
		fmt.Println("comment id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "comment_id格式有误")
		return
	}
	err = service.CommentLikes(commentId)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// deleteComment0 管理员删除评论 不必再进行用户名判断
func deleteComment0(ctx *gin.Context) {
	commentIdString := ctx.Param("comment_id")
	commentId, err := strconv.Atoi(commentIdString)
	if err != nil {
		fmt.Println("comment id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "comment_id格式有误")
		return
	}
	err = service.DeleteComment(commentId)
	if err != nil {
		fmt.Println("delete comment err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
