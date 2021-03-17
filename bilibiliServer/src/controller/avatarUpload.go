// 用户头像上传控制
package controller

import (
	"bilibiliServer/src/model"
	"bilibiliServer/src/tools"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadAvatar(c *gin.Context) { // 用户头像上传
	id := c.PostForm("id")
	log.Println("用户: ", id, " 正在上传头像")

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(400, gin.H{"msg": "头像上传失败!"})
		log.Panicln("头像读取错误: ", err.Error())
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".png" // 创建唯一文件名

	db := tools.GetDb()
	var user model.User

	// 用户有头像则先删除再保存
	db.Model(&user).Where("id = ?", id).First(&user)
	if user.Avatar != "" {
		deleteAvatar(user.Avatar)
	}

	filePath := "./static/avatar/" + fileName // 文件存储路径
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(400, gin.H{"msg": "头像上传失败!"})
		log.Panicln("头像存储错误: ", err.Error())
	}
	db.Model(&user).Where("id = ?", id).Update("avatar", fileName)

	c.JSON(200, gin.H{"avatar": tools.AvatarPath(fileName)})
}

func deleteAvatar(filename string) {
	filePath := "./static/avatar/" + filename
	err := os.Remove(filePath)
	if err != nil {
		panic(err.Error())
	}
}
