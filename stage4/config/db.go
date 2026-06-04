package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"stage4/model"
	"time"
)

// 全局数据库对象
var DB *gorm.DB

func InitDB() {
	// 配置日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
		},
	)

	// 连接Sqlite数据库
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		log.Fatal("表结构迁移失败：", err)
	}

	DB = db
	log.Println("数据库初始化成功！")
}
