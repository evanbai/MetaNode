package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
	Posts    []Post `json:"posts"`
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"title"`
	Content  string    `gorm:"not null" json:"content"`
	UserID   uint      `json:"user_id"`
	User     User      `json:"user"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
	PostID  uint   `json:"post_id"`
}
