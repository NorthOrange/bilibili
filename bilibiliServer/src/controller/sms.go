package controller

import (
	"bilibiliServer/src/model"
	"bilibiliServer/src/tools"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func SendSMS(c *gin.Context) { //
	db := tools.GetDb()
	var sms_record model.Sms_record
	// 获取手机号
	if err := c.BindJSON(&sms_record); err != nil {
		c.JSON(400, "非法输入")
	}

	sms_record.Code = tools.GetSmsCode() // 生成验证码
	log.Println(sms_record.Mobile, " 在发送短信验证码")
	ststua := tools.SendSMS(sms_record.Mobile, sms_record.Code) // 发送验证码
	if ststua == "000000" {
		fiveMinute, _ := time.ParseDuration("5m") // 5分钟后失效
		sms_record.Expiration = time.Now().Add(fiveMinute)
		sms_record.Status = 0
		db.Create(&sms_record) // 生成记录
		c.JSON(200, gin.H{"msg": "验证码发送成功"})
		return
	} else {
		c.JSON(400, gin.H{"msg": "验证码发送失败"})
		return
	}
}
