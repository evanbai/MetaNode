package controller

import (
	"net/http"
	"stage4/config"
	"stage4/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateComment 发布评论
func CreateComment(c *gin.Context) {
	userID, _ := c.Get("userID")
	postID := c.Param("post_id")

	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	comment.UserID = userID.(uint)

	pid, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文章ID格式错误"})
		return
	}
	comment.PostID = uint(pid)

	config.DB.Create(&comment)
	c.JSON(http.StatusOK, gin.H{"msg": "评论成功"})
}

// GetComments 获取文章的所有评论
func GetComments(c *gin.Context) {
	postID := c.Param("post_id")
	var comments []model.Comment
	config.DB.Preload("User").Where("post_id = ?", postID).Find(&comments)
	c.JSON(http.StatusOK, gin.H{"data": comments})
}
