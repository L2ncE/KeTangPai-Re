package api

import (
	"github.com/gin-gonic/gin"
	"ketangpai/tool"
	"os"
)

func download(ctx *gin.Context) {
	fileDir := ctx.Query("fileDir")
	fileName := ctx.Query("fileName")
	//打开文件
	_, err := os.Open(fileDir + "/" + fileName)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}
	ctx.Header("Content-Type", "application/octet-stream")              //常见的文件下载格式
	ctx.Header("Content-Disposition", "attachment; filename="+fileName) //下载到本地
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File(fileDir + "/" + fileName)
	tool.RespSuccessful(ctx)
}
