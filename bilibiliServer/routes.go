package main

import (
	"bilibiliServer/src/controller"
	"bilibiliServer/src/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/user/register", controller.Register)                           // 注册控制
	r.POST("/api/user/login", controller.Login)                                 // 登录控制
	r.GET("/api/user/info/:id", controller.UserInfo)                            // 用户信息查询
	r.POST("/api/user/modify/:id", middleware.JWTAuth(), controller.UserModify) // 用户信息修改
	r.POST("/api/avatar/upload", middleware.JWTAuth(), controller.UploadAvatar) // 用户头像上传

	r.POST("/api/video/upload", middleware.JWTAuth(), controller.VideoUpload) // 视频上传
	r.GET("/api/video/get", controller.VideoGet)                              // 视频访问
	r.GET("/api/user/video/list/:id", controller.VideoListGet)                // 用户发布的视频列表

	r.POST("/api/live/follow/query", controller.FollowQuery)            // 两个用户之间关系查询
	r.POST("/api/live/follow", middleware.JWTAuth(), controller.Follow) // 用户之间关系修改

	r.POST("/api/live/like/query", controller.LikeQuery)            //视频和用户之间关系查询
	r.POST("/api/live/like", middleware.JWTAuth(), controller.Like) // 视频和用户之间关系修改

	r.Static("/avatar", "./static/avatar/")
	r.Static("/video", "./static/video/")

	return r
}
