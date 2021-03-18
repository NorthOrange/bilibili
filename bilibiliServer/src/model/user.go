package model

import "gorm.io/gorm"

// 用户基本信息表
type User struct {
	gorm.Model
	Name         string `gorm:"type:varchar(7);not null" json:"name"`
	Account      string `gorm:"type:varchar(11);not null;unique" json:"account"`
	Password     string `gorm:"type:varchar(11);not null" json:"password"`
	Avatar       string `gorm:"type:varchar(20);" json:"avatar"` // 用户头像地址
	Sex          string `gorm:"type:varchar(1)" json:"sex"`
	Followers    uint   `json:"followers"`                            // 用户粉丝数
	Following    uint   `json:"following"`                            // 用户关注数
	Likes        uint   `json:"likes"`                                // 用户视频获赞数
	Introduction string `gorm:"type:varchar(30)" json:"introduction"` // 个人简介
}

func (User) TableName() string {
	return "users"
}

// 用户之间关系表
type UserFollow struct {
	gorm.Model
	FromId       uint `json:"fromid"`
	ToId         uint `json:"toid"`
	FollowStatus int  `json:"followstatus"`
}

func (UserFollow) TableName() string {
	return "userfollows"
}
