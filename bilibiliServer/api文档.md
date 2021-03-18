# api 文档
### 响应状态码
| 状态码 | 说明                     |
| ------ | ------------------------ |
| 200    | 请求成功                 |
| 400    | 参数校验失败             |
| 401    | 验证失败                 |
| 403    | token 验证失败, 没有权限 |
| 404    | 资源不存在               |
请求失败,响应数据为一个提示信息 `msg`.


### 注册
###### 请求地址
`POST /api/user/register`
###### 请求参数: application/json
| 字段     | 类型   | 说明                       | 是否必选 |
| -------- | ------ | -------------------------- | :------: |
| name     | string | 用户名, 2~7位字符          |    y     |
| account  | string | 帐号, 6~11位纯数字         |    y     |
| password | string | 密码, 6~11位字母数字下划线 |    y     |

###### 成功返回: application/json
| 字段  | 类型   | 说明                   |
| ----- | ------ | ---------------------- |
| msg   | string | 提示信息               |
| id    | string | 后端生成的用户唯一标识 |
| token | string | 根据 id 生成的令牌     |


### 登录
###### 请求地址
`POST /api/user/login`
###### 请求参数: application/json
| 字段     | 类型   | 说明                       | 是否必选 |
| -------- | ------ | -------------------------- | :------: |
| account  | string | 帐号, 6~11位纯数字         |    y     |
| password | string | 密码, 6~11位字母数字下划线 |    y     |
|          |
###### 成功返回: application/json
| 字段  | 类型   | 说明                   |
| ----- | ------ | ---------------------- |
| msg   | string | 提示信息               |
| id    | string | 后端生成的用户唯一标识 |
| token | string | 根据 id 生成的令牌     |


### 用户资料查询
###### 请求地址
`get /api/user/info/:id`

###### 成功返回: application/json
| 字段         | 类型   | 说明           |
| ------------ | ------ | -------------- |
| id           | uint   | 用户id         |
| account      | string | 用户帐号       |
| name         | string | 用户昵称       |
| introduction | string | 用户个人简介   |
| avatar       | string | 用户头像       |
| sex          | string | 用户性别       |
| likes        | uint   | 用户视频获赞数 |
| following    | uint   | 用户关注数     |
| follower     | uint   | 用户粉丝数     |


### 用户头像上传
###### 请求地址
`post /api/avatar/upload`
###### 请求参数: multipart/form-data
###### 需在请求头中添加上: `"Authorization": "Bearer token"`

| 字段   | 类型 | 说明             | 是否必选 |
| ------ | ---- | ---------------- | :------: |
| id     | uint | 用户id           |    y     |
| avatar | file | 要上传的头像图片 |    y     |

###### 成功返回: application/json
| 字段   | 类型   | 说明     |
| ------ | ------ | -------- |
| avatar | string | 用户头像 |

### 用户资料修改
###### 请求地址
`post /api/user/mpdify/:id`
###### 请求参数: application/json
###### 需在请求头中添加上: `"Authorization": "Bearer token"`
| 字段         | 类型   | 说明                       | 是否必选 |
| ------------ | ------ | -------------------------- | :------: |
| name         | string | 用户要修改的昵称,可选      |    y     |
| introduction | string | 用户要修改的个性签名, 可选 |    y     |
| sex          | string | 用户要修改的性别, 可选     |    y     |
###### 成功返回: application/json
| 字段 | 类型   | 说明                       |
| ---- | ------ | -------------------------- |
| msg  | string | 根据修改结果返回的提示消息 |


### 视频上传
###### 请求地址
`post /api/video/upload`
###### 请求参数: multipart/form-data
###### 需在请求头中添加上: `"Authorization": "Bearer token"`

| 字段         | 类型   | 说明                            | 是否必选 |
| ------------ | ------ | ------------------------------- | :------: |
| video        | file   | 视频文件                        |    y     |
| name         | string | 视频标题, 2~50个字符            |    y     |
| introduction | string | 视频简介, 0~100个字符           |    y     |
| cover        | file   | 视频封面,和 `nocover`字段取一个 |    n     |
| nocover      | string | 没有封面上传时, 值应该为 "true" |    n     |
###### 成功返回: application/json
| 字段 | 类型   | 说明     |
| ---- | ------ | -------- |
| msg  | string | 提示消息 |


### 视频信息查询
###### 请求地址
`get /api/video/get`

###### 成功返回: application/json
| 字段             | 类型   | 说明             |
| ---------------- | ------ | ---------------- |
| video            | string | 视频访问地址     |
| videoid          | uint   | 视频id           |
| introduction     | string | 视频简介         |
| cover            | string | 视频封面访问地址 |
| name             | string | 视频标题         |
| likes            | uint   | 视频获赞数       |
| dislikes         | uint   | 视频被踩数       |
| fromid           | uint   | 视频发布者id     |
| username         | string | 视频发布者昵称   |
| avatar           | string | 用户头像         |
| userintroduction | string | 用户简介         |


### 用户视频列表信息查询
###### 请求地址
`get /api/user/video/list/:id`

###### 成功返回: application/json
| 字段      | 类型       | 说明         |
| --------- | ---------- | ------------ |
| namelist  | string数组 | 视频名字列表 |
| coverlist | string数组 | 视频封面列表 |
| idlist    | int数组    | 视频id列表   |

### 用户之间关系查询
###### 请求地址
`get /api/live/follow/query`
###### 请求参数: application/json
| 字段   | 类型   | 说明         | 是否必选 |
| ------ | ------ | ------------ | :------: |
| fromid | number | 关系发起者id |    y     |
| toid   | number | 关系接收者id |    y     |
###### 成功返回: application/json
| 字段         | 类型 | 说明                             |
| ------------ | ---- | -------------------------------- |
| followstatus | int  | 用户之间关系, 0为没关注, 1为关注 |


### 用户之间关系修改
###### 请求地址
`post /api/live/follow`
###### 请求参数: application/json
###### 需在请求头中添加上: `"Authorization": "Bearer token"`

| 字段          | 类型   | 说明                                 | 是否必选 |
| ------------- | ------ | ------------------------------------ | :------: |
| fromid        | number | 关注发起者id                         |    y     |
| toid          | number | 被关注者id                           |    y     |
| folloewstatus | number | 想要设置的关系, 1为关注, 0为取消关注 |    y     |
###### 成功返回: application/json
| 字段 | 类型   | 说明     |
| ---- | ------ | -------- |
| msg  | string | 提示信息 |


### 用户和视频之间关系查询
###### 请求地址
`get /api/live/like/query`
###### 请求参数: application/json
| 字段    | 类型   | 说明   | 是否必选 |
| ------- | ------ | ------ | :------: |
| videoid | number | 视频id |    y     |
| userid  | number | 用户id |    y     |
###### 成功返回: application/json
| 字段       | 类型 | 说明                                            |
| ---------- | ---- | ----------------------------------------------- |
| likestatus | int  | 用户和视频之间关系, 0 为无关系,1为点赞,-1为点踩 |


### 用户和视频之间关系修改
###### 请求地址
`get /api/live/like`
###### 请求参数: application/json
| 字段       | 类型   | 说明                                             | 是否必选 |
| ---------- | ------ | ------------------------------------------------ | :------: |
| videoid    | number | 视频id                                           |    y     |
| userid     | number | 用户id                                           |    y     |
| likestatus | number | 用户和视频之间的关系,0 为无关系,1为点赞,-1为点踩 |    y     |
###### 成功返回: application/json
| 字段       | 类型   | 说明                                                    |
| ---------- | ------ | ------------------------------------------------------- |
| msg        | string | 提示信息                                                |
| likestatus | int    | 更新后的用户和视频之间关系, 0 为无关系,1为点赞,-1为点踩 |