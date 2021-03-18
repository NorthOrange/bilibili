// token 验证中间件
package middleware

import (
	"bilibiliServer/src/tools"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			log.Println("一个未携带 token 的非法请求")
			c.JSON(401, gin.H{
				"msg": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		token = token[7:]
		tokenId := tools.ParseToken(token)

		c.Set("tokenid", tokenId)
		c.Next()

	}
}
