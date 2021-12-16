package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	err := engine.Run()
	if err != nil {
		return
	}
}
