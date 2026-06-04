package controller

import (
	"net/http"
	"stage4/config"
	"stage4/model"
	"stage4/utils"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	//密码加密
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "密码加密失败"})
		return
	}
	user.Password = hash

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名不能为空！"})
		return
	}
	//创建用户
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名/邮箱已存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})

}

// Login 用户登录
func Login(c *gin.Context) {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	// 查询用户
	var user model.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "用户名或密码错误"})
		return
	}

	//验证密码
	if !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "密码验证失败"})
		return
	}

	//生成token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "token生成失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "登录成功",
		"token": token,
	})

}
