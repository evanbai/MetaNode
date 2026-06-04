package controller

import (
	"net/http"
	"stage4/config"
	"stage4/model"

	"github.com/gin-gonic/gin"
)

// CreatePost 创建文章
func CreatePost(c *gin.Context) {
	userID, _ := c.Get("userID")
	var post model.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}
	post.UserID = userID.(uint)

	config.DB.Create(&post)
	c.JSON(http.StatusCreated, gin.H{"msg": "发布成功", "data": post})
}

// GetPosts 获取所有文章
func GetPosts(c *gin.Context) {
	var posts []model.Post
	config.DB.Preload("User").Preload("Comments").Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// GetPost 获取单篇文章
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	if err := config.DB.Preload("User").Preload("Comments").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "文章不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// UpdatePost 更新文章（仅作者）
func UpdatePost(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var post model.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "文章不存在"})
		return
	}

	// 权限校验
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"msg": "无权限修改"})
		return
	}

	var update model.Post
	c.ShouldBindJSON(&update)
	config.DB.Model(&post).Updates(update)
	c.JSON(http.StatusOK, gin.H{"msg": "更新成功"})
}

// DeletePost 删除文章（仅作者）
func DeletePost(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var post model.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "文章不存在"})
		return
	}

	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"msg": "无权限删除"})
		return
	}

	config.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
}
