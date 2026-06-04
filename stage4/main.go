package main

import (
	"stage4/config"
	"stage4/controller"
	"stage4/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 路由
	r := gin.Default()

	// 公共接口
	public := r.Group("/api")
	{
		public.POST("/register", controller.Register)
		public.POST("/login", controller.Login)
		public.GET("/posts", controller.GetPosts)
		public.GET("/posts/:id", controller.GetPost)
		public.GET("/comments/post/:post_id", controller.GetComments)
	}

	// 需要登录的接口
	private := r.Group("/api")
	private.Use(middleware.AuthMiddleware())
	{
		// 文章
		private.POST("/posts", controller.CreatePost)
		private.PUT("/posts/:id", controller.UpdatePost)
		private.DELETE("/posts/:id", controller.DeletePost)

		// 评论
		private.POST("/comments/post/:post_id", controller.CreateComment)
	}

	// 启动服务
	r.Run(":8080")
}
