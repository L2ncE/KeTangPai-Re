package main

import (
	"fmt"
	"ketangpai/api"
	"ketangpai/dao"
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
