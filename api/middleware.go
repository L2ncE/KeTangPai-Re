package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"ketangpai/tool"
)

var mySigningKey = []byte("CSAGolang")

////cookie 中间件
//func auth(ctx *gin.Context) {
//	username, err := ctx.Cookie("username")
//	if err != nil {
//		tool.RespErrorWithDate(ctx, "游客你好！没有您的信息,请先登录!")
//		ctx.Abort()
//		return
//	}
//
//	ctx.Set("username", username)
//	ctx.Next()
//}

// JWTAuth JWT登录
func JWTAuth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token == "" {
		tool.RespErrorWithDate(ctx, "游客你好！没有您的信息,请先登录!")
		ctx.Abort()
		return
	}
	ctx.Set("username", ParseToken(token))
	ctx.Next()
}

func ParseToken(s string) string {
	//解析传过来的token
	tokenClaims, err := jwt.ParseWithClaims(s, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return tokenClaims.Claims.(*model.MyClaims).Username
}
