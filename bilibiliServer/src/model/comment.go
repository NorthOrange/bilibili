package model

import "gorm.io/gorm"

// 用户基本信息表
type Comment struct {
	gorm.Model
	UserId  uint   `json:"userid"`  // 用户id
	VideoId uint   `json:"videoid"` // 被评论视频的id
	Content string `json:"content"` // 评论内容
}

func (c Comment) TableName() string {
	return "commnets"
}
