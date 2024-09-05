package main

import (
	"GenerateWallet/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//注册路由
	r := gin.Default()
	r.POST("/create_wallet", func(ctx *gin.Context) {
		controller.CreateWallet(ctx)
	})
	err := r.Run()
	if err != nil {
		fmt.Println("启动失败")
		return
	} // By default, it listens on :8080

}
