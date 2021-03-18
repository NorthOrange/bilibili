// 控制用户动态
package controller

import (
	"bilibiliServer/src/model"
	"bilibiliServer/src/tools"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FollowQuery(c *gin.Context) { // 两个用户之间关系查询

	db := tools.GetDb()
	var localFollow model.UserFollow
	var remoteFollow model.UserFollow

	err := c.BindJSON(&remoteFollow)
	if err != nil {
		log.Println("数据解析失败: ", err)
	}

	db.Model(&localFollow).Where("from_id = ? AND to_id = ?", remoteFollow.FromId, remoteFollow.ToId).Find(&localFollow)
	log.Println(remoteFollow.FromId, "和", remoteFollow.ToId, "的关系被请求了")
	if localFollow.ID != 0 { // 如果关系存在, 返回两者关系
		c.JSON(200, gin.H{
			"followstatus": localFollow.FollowStatus,
		})
	} else {
		c.JSON(200, gin.H{
			"followstatus": 0,
		})
	}
}

func Follow(c *gin.Context) { // 用户之间关系修改

	db := tools.GetDb()
	var remoteFollow model.UserFollow
	err := c.BindJSON(&remoteFollow)
	if err != nil {
		log.Panicln("数据解析错误", err)
	}
	log.Println("用户: ", remoteFollow.FromId, " 和用户: ", remoteFollow.ToId, "建立关系: ", remoteFollow.FollowStatus)
	if tokenid, b := c.Get("tokenid"); b {
		if id := strconv.Itoa(int(remoteFollow.FromId)); id != tokenid {
			c.JSON(403, gin.H{"msg": "你没有权限!"})
			log.Println("行为发起用户的id为: ", id, " 但是对应 token : ", tokenid)
			return
		}
	}
	var localFollow model.UserFollow
	db.Model(&localFollow).Where("from_id = ? AND to_id = ?", remoteFollow.FromId, remoteFollow.ToId).Find(&localFollow)
	if localFollow.ID == 0 { // 如果关系不存在, 创建关系
		db.Create(&remoteFollow)
	} else {
		localFollow.FollowStatus = remoteFollow.FollowStatus
		db.Save(&localFollow)
	}
	// 更新用户状态
	var fromuser, touser model.User
	db.Model(&fromuser).Where("ID = ?", remoteFollow.FromId).Find(&fromuser)
	db.Model(&touser).Where("ID = ?", remoteFollow.ToId).Find(&touser)
	if remoteFollow.FollowStatus == 0 { // 取消关注
		db.Model(&fromuser).Update("following", fromuser.Following-1)
		db.Model(&touser).Update("followers", touser.Followers-1)

	} else { // 关注
		db.Model(&fromuser).Update("following", fromuser.Following+1)
		db.Model(&touser).Update("followers", touser.Followers+1)
	}

	c.JSON(200, gin.H{
		"msg": "用户关系修改成功",
	})
}

func LikeQuery(c *gin.Context) { //视频和用户之间关系查询

	db := tools.GetDb()
	var localLike model.VideoLike
	var remoteLike model.VideoLike

	err := c.BindJSON(&remoteLike)
	if err != nil {
		log.Println("数据解析失败: ", err)
	}

	db.Model(&localLike).Where("video_id = ? AND user_id = ?", remoteLike.VideoId, remoteLike.UserId).Find(&localLike)
	log.Println(remoteLike.VideoId, "和", remoteLike.UserId, "的关系被请求了")
	if localLike.ID != 0 { // 如果关系存在, 返回两者关系
		c.JSON(200, gin.H{
			"likestatus": localLike.LikeStatus,
		})
	} else {
		c.JSON(200, gin.H{
			"likestatus": 0,
		})
	}
}

func Like(c *gin.Context) { // 用户和视频之间关系修改
	var originStatus int

	db := tools.GetDb()
	var remoteLike model.VideoLike
	err := c.BindJSON(&remoteLike)
	if err != nil {
		log.Panicln("数据解析错误", err)
	}
	log.Println("视频: ", remoteLike.VideoId, " 和用户: ", remoteLike.UserId, "建立关系: ", remoteLike.LikeStatus)
	if tokenid, b := c.Get("tokenid"); b {
		if id := strconv.Itoa(int(remoteLike.UserId)); id != tokenid {
			c.JSON(403, gin.H{"msg": "你没有权限!"})
			log.Println("行为发起用户的id为: ", id, " 但是对应 token : ", tokenid)
			return
		}
	}
	var localLike model.VideoLike
	db.Model(&localLike).Where("video_id = ? AND user_id = ?", remoteLike.VideoId, remoteLike.UserId).Find(&localLike)
	if localLike.ID == 0 { // 如果关系不存在, 创建关系
		db.Create(&remoteLike)
	} else { // 如果关系存在, 记录原来关系, 更改关系
		originStatus = localLike.LikeStatus
		localLike.LikeStatus = remoteLike.LikeStatus
		db.Save(&localLike)
	}
	// 更新用户状态和视频
	// 获取视频和用户
	var video model.Video
	var user model.User
	db.Model(&video).Where("ID = ?", remoteLike.VideoId).Find(&video)
	db.Model(&user).Where("ID = ?", video.FromId).Find(&user)
	// 更新状态
	if originStatus == 0 { // 原来是没有关系
		if remoteLike.LikeStatus == 1 { // 新增一个点赞
			db.Model(&video).Update("likes", video.Likes+1) // 更新视频获得赞数
			db.Model(&user).Update("likes", user.Likes+1)   // 更新用户总获赞数
		} else { // 新增一个踩
			db.Model(&video).Update("dislikes", video.Dislikes+1) // 更新视频被踩数
		}
	} else if originStatus == 1 { // 原来是点赞
		if remoteLike.LikeStatus == 0 { // 一个点赞被取消了
			db.Model(&video).Update("likes", video.Likes-1) // 更新视频获得赞数
			db.Model(&user).Update("likes", user.Likes-1)   // 更新用户总获赞数
		} else { // 从点赞变成点踩
			db.Model(&user).Update("likes", user.Likes-1)         // 更新用户总获赞数
			db.Model(&video).Update("likes", video.Likes-1)       // 更新视频获得赞数
			db.Model(&video).Update("dislikes", video.Dislikes+1) // 更新视频被踩数
		}
	} else { // 原来是点踩
		if remoteLike.LikeStatus == 0 { // 一个踩被取消了
			db.Model(&video).Update("dislikes", video.Dislikes-1) // 更新视频被踩数
		} else { // 一个点踩变成了点赞
			db.Model(&video).Update("dislikes", video.Dislikes-1) // 更新视频被踩数
			db.Model(&video).Update("likes", video.Likes+1)       // 更新视频获得赞数
			db.Model(&user).Update("likes", user.Likes+1)         // 更新用户总获赞数
		}
	}

	c.JSON(200, gin.H{
		"msg":             "用户和视频关系修改成功",
		"videolikestatus": remoteLike.LikeStatus,
	})
}
