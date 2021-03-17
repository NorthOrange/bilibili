package main

import (
	"bilibiliServer/middleware"
	"bilibiliServer/src/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/user/register", controller.Register)                           // 注册控制
	r.POST("/api/user/login", controller.Login)                                 // 登录控制
	r.GET("/api/user/info/:id", controller.UserInfo)                            // 用户信息查询
	r.POST("/api/avatar/upload", middleware.JWTAuth(), controller.UploadAvatar) // 用户头像上传
	r.POST("/api/user/modify/:id", middleware.JWTAuth(), controller.UserModify) // 用户信息修改

	r.Static("/avatar", "./static/avatar/")

	return r
}
