// token 的生成和解析

package tools

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTSecret = []byte("bilibiliServer!")

type MyClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GetToken(id string) string {
	expireTime := time.Now().Add(140 * 24 * time.Hour).Unix()
	claims := &MyClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime, // 过期时间
			Issuer:    "bilibili", // 颁发者
		},
	}
	// 特定算法生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		panic(err.Error())
	}
	return tokenString
}

// 2、解析JWT
func ParseToken(tokenString string) string {
	// 后面是一个匿名函数
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JWTSecret, nil
	})
	if err != nil {
		return ""
	}
	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims.ID
	}
	return ""
}
