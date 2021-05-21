package controller

import (
	"bilibiliServer/src/model"
	"bilibiliServer/src/tools"
	"log"

	"github.com/gin-gonic/gin"
)

func VideoGet(c *gin.Context) { // 随机返回一个视频
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
		"comments":         video.Comments,
	})
}

func VideoGetById(c *gin.Context) { // 根据 id 获取视频
	db := tools.GetDb()
	var video model.Video
	var user model.User
	db.Where("id = ?", c.Param("id")).Find(&video)
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
		"comments":         video.Comments,
	})
}

func VideoListGet(c *gin.Context) { // 查询用户发布的视频列表, 返回视频名数组和视频封面数组
	db := tools.GetDb()
	id := c.Param("id")
	log.Println("用户: ", id, " 查询视频列表")
	var videoList []model.Video
	db.Raw("select name,cover from videos where from_id = ?", id).Find(&videoList)
	var videoName, videoCover, idlist []interface{}
	for _, v := range videoList {
		videoName = append(videoName, v.Name)
		videoCover = append(videoCover, tools.VideoPath(v.Cover))
		idlist = append(idlist, v.ID)
	}
	c.JSON(200, gin.H{"namelist": videoName, "coverlist": videoCover, "idlist": idlist})
}

func VideoComment(c *gin.Context) { // 评论视频
	var comment model.Comment
	c.BindJSON(&comment)
	db := tools.GetDb()
	db.Create(comment)
	c.JSON(200, gin.H{"msg": "评论成功"})
}
