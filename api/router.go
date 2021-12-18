package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register) //注册
	engine.POST("/login", login)       //登陆

	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/password", changePassword) //修改密码
	}

	topicGroup := engine.Group("/topic")
	{
		{
			topicGroup.Use(auth)
			topicGroup.POST("/", addTopic)               //发布新话题
			topicGroup.POST("/:topic_id", changeTopic)   //修改话题
			topicGroup.DELETE("/:topic_id", deleteTopic) //删除话题

			topicGroup.GET("/", briefTopics)          //查看全部话题概略
			topicGroup.GET("/:topic_id", topicDetail) //查看一条话题详细信息和其下属评论

			topicGroup.POST("/:topic_id/likes", topicLikes) //给话题点赞
		}
	}

	commentGroup := engine.Group("/comment")
	{
		commentGroup.POST("/anonymity", addCommentAnonymity) //匿名评论
		{
			commentGroup.Use(auth)
			commentGroup.POST("/", addComment)                 //发送评论
			commentGroup.DELETE("/:comment_id", deleteComment) //删除评论

			commentGroup.POST("/:comment_id/likes", commentLikes) //给评论点赞
		}
	}

	err := engine.Run()
	if err != nil {
		return
	}
}
