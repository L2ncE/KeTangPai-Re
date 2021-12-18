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

func addComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	Name := iUsername.(string)

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

// addCommentAnonymity 匿名评论
func addCommentAnonymity(ctx *gin.Context) {
	Name := "Anonymity"
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

func deleteComment(ctx *gin.Context) {
	commentIdString := ctx.Param("comment_id")
	commentId, err := strconv.Atoi(commentIdString)
	commentNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameById2(commentId)
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
