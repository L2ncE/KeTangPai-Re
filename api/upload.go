package api

import (
	"github.com/gin-gonic/gin"
	"ketangpai/tool"
)

// upload 上传文件
func upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		tool.RespErrorWithDate(ctx, "上传文件出错")
		return
	}
	err = ctx.SaveUploadedFile(file, file.Filename) //保存
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
