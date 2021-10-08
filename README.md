# HDUHelp-Icebreaking

### 前言：

### 由于一些不可控因素，导致时间比较赶，所以部分功能尚未实现

#### 功能实现：

- 实现了用户的注册和登录，同时选择了我个人比较熟悉的jwt来进行鉴权

- 实现了猜人名以及头像上传的操作

#### 尚未实现：

- oauth的鉴权方式
- 隐藏自己的头像（这个实现比较容易，在ddl前可能会补全）

#### 可能存在的漏洞：

- token被截获导致的csrf（考虑上https来防范）
- 可能存在的xss漏洞（我这里交给前端的渲染来解决这个问题）

#### 演示视频正在做（ddl前完工）

[演示视频][http://49.235.65.44:8888/down/LevphCVxKWZl] (需下载)

#### 项目结构
```
.
│  go.mod
│  go.sum
│  main.go
│
├─Controller
│      check.go
│      guess.go
│      Login.go
│      NotFound.go
│      Register.go
│      Upload.go
│
├─Middlewires
│      AuthJWT.go
│      Token.go
│
├─Model
│      check.go
│      fecth_date.go
│      InitSql.go
│      Login.go
│      Register.go
│      Upload.go
│
└─Router
        router.go
```
