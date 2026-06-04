# Go + Gin + GORM 个人博客系统后端

## 项目介绍
本项目是基于 Go 语言开发的个人博客后端系统，使用 Gin 框架搭建 Web 服务，GORM 实现数据库操作，SQLite 作为数据库，JWT 实现用户认证功能。
实现了用户注册登录、文章 CRUD、评论管理、权限控制等完整功能。

---

## 技术栈
- Go 1.18+
- Gin Web 框架
- GORM ORM 库
- SQLite 数据库
- JWT 身份认证
- bcrypt 密码加密

---

## 项目结构

### stage4
├── config/ # 数据库配置  
│ └── db.go  
├── controller/ # 接口控制器  
│ ├── user.go  
│ ├── post.go  
│ └── comment.go  
├── middleware/ # 认证中间件    
│ └── auth.go  
├── model/ # 数据模型  
│ └── model.go  
├── utils/ # 工具类  
│ ├── jwt.go  
│ └── password.go  
├──  blog.db  
├── main.go # 项目入口  
├── go.mod  
│ └── go.sum  
└── README.md  
## 功能清单
✅ 用户注册（密码加密存储）  
✅ 用户登录（JWT 签发）  
✅ 文章创建、查询、更新、删除  
✅ 仅文章作者可修改 / 删除文章  
✅ 评论发布、按文章查询评论  
✅ 统一错误处理  
✅ 接口权限控制  
✅ 数据库自动迁移  
## 环境要求
安装 Go 语言环境（1.18 及以上版本）  
## 运行步骤
### 1. 安装依赖
   bash
   运行
   go mod tidy
### 2. 启动项目
   bash
   运行
   go run main.go
### 3. 服务地址
http://localhost:8080
启动后自动生成 blog.db 数据库文件，并自动创建数据表。  
## 接口文档
### 公共接口（无需登录）  

|请求方式|接口地址|说明| 
|---|---|---|
|POST|/api/register|用户注册|  
|POST|/api/login|用户登录，返回 token  
|GET|/api/posts|获取所有文章  
|GET|/api/posts/:id|获取单篇文章详情  
|GET|/api/posts/:post_id/comments|获取文章的所有评论  
### 私有接口（需要登录，请求头携带 Token）  
#### 请求头格式：Authorization: Bearer 你的token  
|请求方式|接口地址|说明|
|---|---|---|
|POST|/api/posts|创建文章
|PUT|/api/posts/:id|更新文章（仅作者）
|DELETE|/api/posts/:id|删除文章（仅作者）
|POST|/api/posts/:post_id/comments|发布评论
## 接口请求示例
### 1. 注册
    {
      "username": "test",
      "password": "123456",
      "email": "test@qq.com"
    }
### 2. 登录
    {
      "username": "test",
      "password": "123456"
    }
### 3. 创建文章
    {
      "title": "我的第一篇博客",
      "content": "这是文章内容"
    }
### 4. 发布评论
    {
      "content": "这是一条评论"
    }
## 数据库表说明
+ users：用户信息（id、用户名、密码、邮箱、时间字段）
+ posts：文章信息（id、标题、内容、作者 ID、时间字段）
+ comments：评论信息（id、内容、用户 ID、文章 ID、时间字段）
## 错误码说明
+ 400：参数错误
+ 401：未登录 / Token 无效
+ 403：无权限操作
+ 404：资源不存在
+ 500：服务器错误