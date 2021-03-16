# <img src="./bilibili_logo.jpg" width="60px" height="60px"> BiliBili

仿 [哔哩哔哩移动端](https://m.bilibili.com/) 制作的动漫小视频网站。搭配[bilibiliHelp](http://github.com/NorthOrange/bilibiliHelp)项目食用效果更佳.

> 此项目仅供学习和交流

## 用到的框架、库和工具
### [前端](./bilibiliServer)
###### html+css+js
[vue](https://cn.vuejs.org/):一套构建用户界面的渐进式框架。 <br>
[vant](https://youzan.github.io/vant/#/zh-CN/): 一个开源的移动端组件库。 <br>
[axios](http://www.axios-js.com/): 一个易用、简洁且高效的http库。 <br>
### 后端
###### golang
[gin](https://github.com/gin-gonic/gin): 一个 go 语言编写的 web 框架。 <br>
[gorm](https://gorm.io/zh_CN/): go 语言的 ORM 类库。 <br>
[go-jwt](https://github.com/dgrijalva/jwt-go): json web token 的 go 语言实现。<br>

### 其它
ffmpeg: Fmpeg 是领先的多媒体框架，能够解码、编码、转码、混合、解密、流媒体、过滤和播放人类和机器创造的几乎所有东西。用来实现对用户上传的视频进行转码和封面抽取。<br>

MySQL: 是最流行的关系型数据库管理系统，在 WEB 应用方面 MySQL 是最好的 RDBMS(Relational Database Management System：关系数据库管理系统)应用软件之一。 <br>


## 项目部署小提示
你应该添加两个这样的文件: <br>
`./bilibiliServer//config/app.json`:

```
{
    "DbUsername": "user name",
    "DbPassword": "user password",
    "DbName": "database name"
    "socket: "http://127.0.0.1:11117"
}
```

`./bilibili/src/config.js`:

```
const socket = "http://127.0.0.1:11117";
export default socket;
```

## 功能实现

- [x] 注册、登录
