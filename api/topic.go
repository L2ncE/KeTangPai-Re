package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"ketangpai/service"
	"ketangpai/tool"
	"strconv"
	"time"
)

// topicDetail 话题详细信息和其下属评论
func topicDetail(ctx *gin.Context) {
	topicIdString := ctx.Param("topic_id") //输入话题id
	topicId, err := strconv.Atoi(topicIdString)
	if err != nil {
		fmt.Println("topic id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "topic_id格式有误")
		return
	}

	//根据topicId拿到topic
	topic, err := service.GetTopicById(topicId)
	if err != nil {
		fmt.Println("get topic by id err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	//找到它的评论
	comments, err := service.GetTopicComments(topicId)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("get topic comments err: ", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	var topicDetail model.TopicDetail
	topicDetail.Topic = topic
	topicDetail.Comments = comments

	tool.RespSuccessfulWithDate(ctx, topicDetail)
}

func briefTopics(ctx *gin.Context) {
	topics, err := service.GetTopics()
	if err != nil {
		fmt.Println("get topics err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, topics)
}

// addTopic 添加评论
func addTopic(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //得到用户名
	name := iUsername.(string)

	context := ctx.PostForm("context")

	topic := model.Topic{
		Context:    context,
		Name:       name,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}

	err := service.AddTopic(topic)
	if err != nil {
		fmt.Println("add topic err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// deleteTopic 删除话题
func deleteTopic(ctx *gin.Context) {
	topicIdString := ctx.Param("topic_id")
	topicId, err := strconv.Atoi(topicIdString)
	topicNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameById(topicId)
	//必须用户名相同,无法删除他人话题
	if topicNameString == nameString {
		if err != nil {
			fmt.Println("topic id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "topic_id格式有误")
			return
		}
		err = service.DeleteTopic(topicId)
		if err != nil {
			fmt.Println("delete topic err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "不能删除他人评论")
	}
}

// changeTopic 更改话题
func changeTopic(ctx *gin.Context) {
	newTopic := ctx.PostForm("newTopic")
	iTopicId := ctx.Param("topic_id")
	topicId, err := strconv.Atoi(iTopicId)
	UpdateTime := time.Now()
	topicNameString, _ := ctx.Get("username")
	//只能更改自己的话题
	nameString, _ := service.GetNameById(topicId)
	if topicNameString == nameString {
		if err != nil {
			fmt.Println("topic id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "topic_id格式有误")
			return
		} else {
			//更改后会更新时间
			err := service.ChangeTopic(topicId, newTopic, UpdateTime)
			if err != nil {
				fmt.Println("change topic err: ", err)
				tool.RespInternalError(ctx)
				return
			}
			tool.RespSuccessful(ctx)
		}
	} else {
		tool.RespErrorWithDate(ctx, "无法更改他人留言")
	}
}

// topicLikes 话题点赞
func topicLikes(ctx *gin.Context) {
	topicIdString := ctx.Param("topic_id")
	topicId, err := strconv.Atoi(topicIdString)
	if err != nil {
		fmt.Println("topic id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "topic_id格式有误")
		return
	}
	err = service.TopicLikes(topicId)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// deleteTopic0 管理员删除话题
func deleteTopic0(ctx *gin.Context) {
	topicIdString := ctx.Param("topic_id")
	topicId, err := strconv.Atoi(topicIdString)
	//无需判断用户名
	if err != nil {
		fmt.Println("topic id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "topic_id格式有误")
		return
	}
	err = service.DeleteTopic(topicId)
	if err != nil {
		fmt.Println("delete topic err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
