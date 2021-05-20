package model

import (
	"time"

	"gorm.io/gorm"
)

// 视频基本信息表
type Sms_record struct {
	gorm.Model
	Mobile     string    `json:"mobile"`     // 手机号
	Expiration time.Time `json:"expiration"` // 过期时间
	Status     int       `json:"status"`     // 验证码状态，0为未使用， 1为已使用
	Code       string    `json:"code"`
}

func (s Sms_record) TableName() string {
	return "sms_record"
}
