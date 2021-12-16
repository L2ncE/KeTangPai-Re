package main

import (
	"fmt"
	"message-board/api"
	"message-board/dao"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	} else {
		fmt.Println("连接数据库成功!")
	}
	api.InitEngine()
}
