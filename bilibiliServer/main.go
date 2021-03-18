package main

import (
	"bilibiliServer/src/middleware"
	"bilibiliServer/src/tools"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 配置文件解析
	err := tools.ParseConfig("./config/config.json")
	if err != nil {
		log.Panicln(err)
	}
	// 资源存储文件夹创建
	err = os.MkdirAll("./static/avatar", os.ModePerm)
	if err != nil {
		log.Panicln("avatar 文件夹创建失败")
	}
	err = os.MkdirAll("./static/video", os.ModePerm)
	if err != nil {
		log.Panicln("video 文件夹创建失败")
	}
	// 日志文件创建
	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// 数据库初始化
	err = tools.InitDb()
	if err != nil {
		log.Panicln(err)
	}
	// 框架初始化
	r := gin.Default()
	r.Use(middleware.Cors()) // 跨域解决
	// 路由注册
	r = CollectRoute(r)

	go tools.TransCodingVideo() // 视频转码服务
	r.Run(":11117")
}
