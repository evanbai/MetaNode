package stage3

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	gorm.Model
	Name   string
	Email  string
	Mobile string
	Posts  []Post
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
}

func Run2() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DriverName:        "mysql",
		DSN:               "root:root@tcp(127.0.0.1:3306)/metanote?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize: 255,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "evan_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	M := db.Migrator()

	if M != nil {
		if !M.HasTable(&User{}) {
			fmt.Println("Create User table")
			db.AutoMigrate(User{})
		}
		if !M.HasTable(&Post{}) {
			fmt.Println("Create Post table")
			db.AutoMigrate(Post{})
		}
		if !M.HasTable(&Comment{}) {
			fmt.Println("Create Comment table")
			db.AutoMigrate(Comment{})
		}
	}

	//初始化表结构及数据
	//initData(db)

	var user User

	//	查询某用户所有的文章及评论
	//user.ID = 2
	//db.Model(&User{}).Preload("Posts.Comments").Preload("Posts").Find(&user)
	//fmt.Println(user)

	//查询评论数量最多的文章信息
	var maxCommentPostID uint
	err := db.Model(&Comment{}).
		Select("post_id").
		Group("post_id").
		Order("count(*) DESC").
		Limit(1).
		Find(&maxCommentPostID).Error
	if err != nil {
		fmt.Println("查找评论数最多的文章失败：", err)
		return
	}
	fmt.Println(maxCommentPostID)
	err = db.Model(&User{}).Preload("Posts.Comments").Preload("Posts", "ID=?", maxCommentPostID).Find(&user).Error
	if err != nil {
		fmt.Println("查找评论数最多的文章的用户失败：", err)
		return
	}
	fmt.Println(user)

}

func initData(db *gorm.DB) {
	c1 := Comment{
		Model: gorm.Model{
			ID: 2001,
		},
		Content: "Very Good!",
	}
	c2 := Comment{
		Model: gorm.Model{
			ID: 2002,
		},
		Content: "Very Nice!",
	}
	c3 := Comment{
		Model: gorm.Model{
			ID: 2003,
		},
		Content: "Very Bad!",
	}
	c4 := Comment{
		Model: gorm.Model{
			ID: 2004,
		},
		Content: "Very Beautiful!",
	}

	p1 := Post{
		Model: gorm.Model{
			ID: 100,
		},
		Title:    "hello world",
		Content:  "hello world,hahahahaha",
		Comments: []Comment{c1, c2},
	}
	p2 := Post{
		Model: gorm.Model{
			ID: 101,
		},
		Title:    "hello Beijing",
		Content:  "hello Beijing,hahahahaha",
		Comments: []Comment{c3},
	}
	p3 := Post{
		Model: gorm.Model{
			ID: 102,
		},
		Title:    "hello Shanghai",
		Content:  "hello Shanghai,hahahaha",
		Comments: []Comment{c4},
	}

	u1 := User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:  "Alice",
		Posts: []Post{p1, p2},
	}
	u2 := User{
		Model: gorm.Model{
			ID: 2,
		},
		Name:  "Bob",
		Posts: []Post{p3},
	}
	db.Create(&u1)
	db.Create(&u2)
}
