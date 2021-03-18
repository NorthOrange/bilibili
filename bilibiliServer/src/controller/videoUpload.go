// 视频上传控制
package controller

import (
	"bilibiliServer/src/model"
	"bilibiliServer/src/tools"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func VideoUpload(c *gin.Context) {
	// 获取上传信息
	var video model.Video

	if tokenid, b := c.Get("tokenid"); b {
		if id := c.PostForm("id"); id != tokenid {
			c.JSON(403, gin.H{"msg": "你没有权限!"})
			log.Println("行为发起用户的id为: ", id, " 但是对应 token : ", tokenid)
			return
		}
	}
	// 发布者 id
	fromid, _ := strconv.Atoi(c.PostForm("id"))
	video.FromId = uint(fromid)

	video.Name = c.PostForm("name")
	log.Println("用户: ", fromid, " 正在上传视频: ", video.Name)
	video.Introduction = c.PostForm("introduction")
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(400, gin.H{"msg": "视频上传失败!"})
		log.Panicln("视频读取失败: ", err.Error())
		return
	}

	// 查询数据库是否有该视频
	db := tools.GetDb()
	var qVideo model.Video
	db.Model(&qVideo).Where("name = ?", video.Name).First(&qVideo)
	if qVideo.Name != "" {
		c.JSON(400, gin.H{"msg": "视频已经存在, 请换个名字!"})
		log.Println("视频名已存在")
		return
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".mp4" // 创建唯一文件名

	// 视频存储
	filePath := "./static/video/" + fileName
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(400, gin.H{"msg": "视频上传失败!"})
		log.Panicln("视频存储失败: ", err.Error())
	}

	// 查询是否上传视频封面
	if c.PostForm("nocover") == "true" { //
		log.Println("用户没有上传封面")
		tools.VideoCoverGet(filePath) // 没有上传封面, 用视频第一帧当封面
	} else {
		coverFile, err := c.FormFile("cover")
		if err != nil {
			log.Panicln("视频封面获取失败: ", err.Error())
		}

		coverPath := strings.Replace(filePath, ".mp4", ".jpg", -1) // 用户上传了封面, 跟视频同名存储
		err = c.SaveUploadedFile(coverFile, coverPath)
		if err != nil {
			log.Panicln("视频封面存储失败: ", err.Error())
		}
	}
	video.Cover = strings.Replace(fileName, ".mp4", ".jpg", -1)

	// 视频入库
	video.Path = fileName
	db.Create(&video)

	// 转码存储
	tools.VideoNameChan <- filePath
	c.JSON(200, gin.H{"msg": "视频上传成功~"})

}
