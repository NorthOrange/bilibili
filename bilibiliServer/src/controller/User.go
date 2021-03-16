package controller

import (
	"bilibiliServer/src/model"
	"bilibiliServer/src/tools"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) { // 用户注册控制
	db := tools.GetDb()

	var user model.User
	c.BindJSON(&user)
	if b, _ := regexp.MatchString("^[^ ]{2,7}$", user.Name); !b {
		c.JSON(400, gin.H{"msg": "昵称不符合规范!"})
		return
	}
	if b, _ := regexp.MatchString("^[0-9]{6,11}$", user.Account); !b {
		c.JSON(400, gin.H{"msg": "用户名不符合规范!"})
		return
	}
	if b, _ := regexp.MatchString("^[a-zA-Z0-9_-]{6,11}$", user.Password); !b {
		c.JSON(400, gin.H{"msg": "密码不符合规范!"})
		return
	}
	if isUserExits(db, user.Account) {
		c.JSON(400, gin.H{"msg": "该用户帐号已存在!"})
	} else {
		db.Create(&user)
		token := tools.GetToken(strconv.Itoa(int(user.ID))) //  给注册成功的用户返回 token 和 id
		c.JSON(200, gin.H{"msg": "注册成功~", "id": user.ID, "token": token})
	}

}
func Login(c *gin.Context) { // 用户登录
	db := tools.GetDb()
	var remoteUser model.User
	c.BindJSON(&remoteUser)
	if b, _ := regexp.MatchString("^[0-9]{6,11}$", remoteUser.Account); !b {
		c.JSON(400, gin.H{"msg": "用户名不规范"})
		return
	}
	if b, _ := regexp.MatchString("^[a-zA-Z0-9_-]{6,11}$", remoteUser.Password); !b {
		c.JSON(400, gin.H{"msg": "密码不符合规范!"})
		return
	}

	if isUserExits(db, remoteUser.Account) {
		var localUser model.User
		db.Where("Account = ?", remoteUser.Account).First(&localUser)
		if remoteUser.Password != localUser.Password {
			c.JSON(401, gin.H{"msg": "密码错误!"})
		} else {
			token := tools.GetToken(strconv.Itoa(int(localUser.ID)))
			c.JSON(200, gin.H{"msg": "登陆成功~", "id": localUser.ID, "token": token})
		}
	} else {
		c.JSON(400, gin.H{"msg": "帐号不存在!"})
	}
}

func isUserExits(db *gorm.DB, account string) bool {
	var user model.User
	db.Where("account = ?", account).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}
