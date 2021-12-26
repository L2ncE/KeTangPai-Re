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
		userGroup.Use(JWTAuth)                      //需要token
		userGroup.POST("/password", changePassword) //修改密码
	}

	topicGroup := engine.Group("/topic")
	{
		{
			topicGroup.Use(JWTAuth)                             //需要token
			topicGroup.POST("/", addTopic)                      //发布新话题
			topicGroup.POST("/:topic_id", changeTopic)          //修改话题
			topicGroup.DELETE("/:topic_id", deleteTopic)        //删除话题
			topicGroup.POST("/:topic_id/likes", topicLikes)     //给话题点赞
			topicGroup.DELETE("/:topic_id/admin", deleteTopic0) //管理员删除话题
		}
		topicGroup.GET("/", briefTopics)          //查看全部话题概略
		topicGroup.GET("/:topic_id", topicDetail) //查看一条话题详细信息和其下属评论
	}

	commentGroup := engine.Group("/comment")
	{
		commentGroup.POST("/anonymity", addCommentAnonymity) //匿名评论
		{
			commentGroup.Use(JWTAuth)                                 //需要token
			commentGroup.POST("/", addComment)                        //发送评论
			commentGroup.DELETE("/:comment_id", deleteComment)        //删除评论
			commentGroup.DELETE("/:comment_id/admin", deleteComment0) //管理员删除评论
			commentGroup.POST("/:comment_id/likes", commentLikes)     //给评论点赞
		}
	}

	classroomGroup := engine.Group("/classroom")
	{
		classroomGroup.Use(JWTAuth)                                 //需要token
		classroomGroup.POST("/", addClassRoom)                      //开启新课堂
		classroomGroup.DELETE("/:classroom_id", deleteClassRoom)    //删除课堂
		classroomGroup.POST("/:classroom_id/open", openClassRoom)   //开启课堂
		classroomGroup.POST("/:classroom_id/close", closeClassRoom) //关闭课堂
		classroomGroup.POST("/:classroom_id/sign", signInClassRoom) //课堂签到
	}

	homeworkGroup := engine.Group("/homework")
	{
		homeworkGroup.Use(JWTAuth)                            //需要token
		homeworkGroup.POST("/", addHomework)                  //布置作业
		homeworkGroup.DELETE("/:homework_id", deleteHomework) //删除作业
	}

	uploadGroup := engine.Group("/upload")
	{
		homeworkGroup.Use(JWTAuth) //需要token
		//限制上传最大尺寸
		engine.MaxMultipartMemory = 8 << 20
		uploadGroup.POST("/", upload) //上传文件
	}

	downloadGroup := engine.Group("/download")
	{
		downloadGroup.Use(JWTAuth)       //需要token
		downloadGroup.GET("/", download) //下载文件
	}

	questionAnswerGroup := engine.Group("/q&a")
	{
		questionAnswerGroup.Use(JWTAuth)                //需要token
		questionAnswerGroup.POST("/question", question) //课中提问
		questionAnswerGroup.POST("/answer", answer)     //课中回答
	}

	gradeGroup := engine.Group("/grade")
	{
		gradeGroup.Use(JWTAuth)                    //需要token
		gradeGroup.POST("/", addGrade)             //新增成绩
		gradeGroup.POST("/:grade_id", changeGrade) //更新成绩
	}

	err := engine.Run()
	if err != nil {
		return
	}
}
