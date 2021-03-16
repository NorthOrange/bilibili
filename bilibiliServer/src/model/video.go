package model

import "gorm.io/gorm"

// 视频基本信息表
type Video struct {
	gorm.Model
	Name         string `gorm:"type:varchar(50);not null" json:"name"` // 视频名字
	Introduction string `gorm:"type:varchar(100)" json:"introduction"` // 视频简介
	Path         string `gorm:"type:varchar(20);not null" json:"path"` // 视频访问地址
	Cover        string `gorm:"type:varchar(20)" json:"cover"`         // 视频封面
	Likes        uint   `json:"likes"`                                 // 视频获赞数
	Dislikes     uint   `json:"dislikes"`                              // 视频被踩数
	FromId       uint   `json:"fromid"`                                // 发布者 id
}

func (Video) TableName() string {
	return "videos"
}
