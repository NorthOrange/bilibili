package tools

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/cloopen/go-sms-sdk/cloopen"
)

func SendSMS(mobile string, code string) string {
	cfg := cloopen.DefaultConfig().
		// 开发者主账号,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIAccount("8a216da87979249801798322587e0445").
		// 主账号令牌 TOKEN,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIToken("76aa35457b8a41628c475eae0d9527ef")
	sms := cloopen.NewJsonClient(cfg).SMS()
	// 下发包体参数
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: "8a216da87979249801798322596d044c",
		// 手机号码
		To: mobile,
		// 模版ID
		TemplateId: "1",
		// 模版变量内容 非必填
		Datas: []string{code, "5"},
	}
	// 下发
	resp, err := sms.Send(input)
	if err != nil {
		log.Fatal(err)
		return "404"
	}
	return resp.StatusCode

}

func GetSmsCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
