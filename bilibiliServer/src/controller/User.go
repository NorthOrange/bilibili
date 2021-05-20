package controller

import (
	"bilibiliServer/src/model"
	"bilibiliServer/src/tools"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	SmsCode  string `json:"smsCode"`
}

func Register(c *gin.Context) { // 用户注册控制
	db := tools.GetDb()

	var user model.User
	var input RegisterInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, "非法输入")
	}
	user.Name = input.Name
	user.Mobile = input.Mobile
	user.Password = input.Password
	smsCode := input.SmsCode

	if b, _ := regexp.MatchString("^[^ ]{2,7}$", user.Name); !b {
		c.JSON(400, gin.H{"msg": "昵称不符合规范!"})
		return
	}
	if b, _ := regexp.MatchString("^1[0-9]{10}$", user.Mobile); !b {
		c.JSON(400, gin.H{"msg": "手机号不符合规范!"})
		return
	}
	if b, _ := regexp.MatchString("^[a-zA-Z0-9_-]{6,11}$", user.Password); !b {
		c.JSON(400, gin.H{"msg": "密码不符合规范!"})
		return
	}
	if isUserExits(db, user.Mobile) {
		c.JSON(400, gin.H{"msg": "该用户帐号已存在!"})
	} else {
		var sms_record model.Sms_record
		db.Where("mobile = ? and status =0", user.Mobile).Order("id desc").First(&sms_record) // 找到手机号对应的验证码
		if time.Now().After(sms_record.Expiration) {
			c.JSON(400, gin.H{"msg": "验证码已过期，请重新发送!"})
			return
		}
		if smsCode != sms_record.Code {
			c.JSON(400, gin.H{"msg": "验证码错误!"})
			return
		}
		db.Create(&user)
		token := tools.GetToken(strconv.Itoa(int(user.ID))) //  给注册成功的用户返回 token 和 id
		c.JSON(200, gin.H{"msg": "注册成功~", "id": user.ID, "token": token})
	}

}
func Login(c *gin.Context) { // 用户登录
	db := tools.GetDb()
	var remoteUser model.User
	err := c.BindJSON(&remoteUser)
	if err != nil {
		log.Panicln("数据解析错误: ", err)
	}
	if b, _ := regexp.MatchString("^1[0-9]{10}$", remoteUser.Mobile); !b {
		c.JSON(400, gin.H{"msg": "手机号不规范"})
		return
	}
	if b, _ := regexp.MatchString("^[a-zA-Z0-9_-]{6,11}$", remoteUser.Password); !b {
		c.JSON(400, gin.H{"msg": "密码不符合规范!"})
		return
	}

	if isUserExits(db, remoteUser.Mobile) {
		var localUser model.User
		db.Where("mobile = ?", remoteUser.Mobile).First(&localUser)
		if remoteUser.Password != localUser.Password {
			c.JSON(401, gin.H{"msg": "密码错误!"})
		} else {
			token := tools.GetToken(strconv.Itoa(int(localUser.ID))) // 给登录成功的用户返回 id 和 token
			c.JSON(200, gin.H{"msg": "登陆成功~", "id": localUser.ID, "token": token})
		}
	} else {
		c.JSON(400, gin.H{"msg": "帐号不存在!"})
	}
}

func UserInfo(c *gin.Context) { // 用户信息查询
	id := c.Param("id")

	log.Println("用户 " + id + " 进行了一次数据查询")
	db := tools.GetDb()
	var user model.User
	db.Where("id = ?", id).First(&user)
	c.JSON(200, gin.H{"id": id, "name": user.Name, "following": user.Following, "followers": user.Followers, "likes": user.Likes, "introduction": user.Introduction, "mobile": user.Mobile, "avatar": tools.AvatarPath(user.Avatar), "sex": user.Sex})
}

func UserModify(c *gin.Context) { // 用户资料修改
	id := c.Param("id")
	if tokenid, b := c.Get("tokenid"); b {
		if id != tokenid {
			c.JSON(403, gin.H{"msg": "你没有权限!"})
			log.Println("行为发起用户的id为: ", id, " 但是对应 token : ", tokenid)
			return
		}
	}
	log.Println("用户: ", id, " 正在修改资料")
	var user model.User
	c.BindJSON(&user)

	db := tools.GetDb()
	if user.Name != "" {
		if b, _ := regexp.MatchString("^[^ ]{2,7}$", user.Name); !b {
			c.JSON(400, gin.H{"msg": "昵称不符合规范!"})
			return
		} // 有用户名且符合规范, 修改用户名
		db.Model(&user).Where("id = ?", id).Update("name", user.Name)
		log.Println("用户", id, "修改用户名成功~")
		c.JSON(200, gin.H{"msg": "用户名修改成功~"})
	}
	if user.Introduction != "" {
		// 个性签名不为空, 修改个性签名
		db.Model(user).Where("id = ?", id).Update("introduction", user.Introduction)
		log.Println("用户", id, "修改签名成功~")

		c.JSON(200, gin.H{"msg": "个性签名修改成功~"})
	}
	if user.Sex != "" {
		if user.Sex != "男" && user.Sex != "女" {
			c.JSON(400, gin.H{"msg": "性别不符合规范!"})
			return
		}
		// 性别不为空且符合规范, 修改性别
		db.Model(user).Where("id = ?", id).Update("sex", user.Sex)
		log.Println("用户", id, "修改性别成功~")

		c.JSON(200, gin.H{"msg": "性别修改成功~"})
	}
}

func isUserExits(db *gorm.DB, mobile string) bool {
	var user model.User
	db.Where("mobile = ?", mobile).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}
