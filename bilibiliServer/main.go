package main

import "github.com/gin-gonic/gin"

func main() {
	// 初始化框架
	r := gin.Default()

	r.Run(":11117")
}
