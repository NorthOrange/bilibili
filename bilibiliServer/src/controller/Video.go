package controller

import (
	"bilibiliServer/src/model"
	"bilibiliServer/src/tools"
	"log"

	"github.com/gin-gonic/gin"
)

func GetVideo(c *gin.Context) { // 随机返回一个视频
	db := tools.GetDb()
	var video model.Video
	var user model.User
	db.Raw("select * from videos order by rand() limit 1").Find(&video)
	log.Println(video.Name + "被请求了")
	db.Model(&user).Where("id = ?", video.FromId).Find(&user)
	c.JSON(200, gin.H{
		"videoid":          video.ID,
		"name":             video.Name,
		"introduction":     video.Introduction,
		"video":            tools.VideoPath(video.Path),
		"likes":            video.Likes,
		"dislike":          video.Dislikes,
		"cover":            tools.VideoPath(video.Cover),
		"fromid":           video.FromId,
		"username":         user.Name,
		"avatar":           tools.AvatarPath(user.Avatar),
		"userintroduction": user.Introduction,
	})
}
