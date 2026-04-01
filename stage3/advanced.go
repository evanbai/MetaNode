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
	Name       string
	Email      string
	Mobile     string
	Posts      []Post
	PostsCount int64
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint
	Comments      []Comment
	CommentsCount int64
	Status        string
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

	//var user User

	//	查询某用户所有的文章及评论
	//user.ID = 2
	//db.Model(&User{}).Preload("Posts.Comments").Preload("Posts").Find(&user)
	//fmt.Println(user)

	//查询评论数量最多的文章信息
	//var maxCommentPostID uint
	//err := db.Model(&Comment{}).
	//	Select("post_id").
	//	Group("post_id").
	//	Order("count(*) DESC").
	//	Limit(1).
	//	Find(&maxCommentPostID).Error
	//if err != nil {
	//	fmt.Println("查找评论数最多的文章失败：", err)
	//	return
	//}
	//fmt.Println(maxCommentPostID)
	//err = db.Model(&User{}).Preload("Posts.Comments").Preload("Posts", "ID=?", maxCommentPostID).Find(&user).Error
	//if err != nil {
	//	fmt.Println("查找评论数最多的文章的用户失败：", err)
	//	return
	//}
	//fmt.Println(user)

	//p := Post{
	//	Model: gorm.Model{
	//		ID: 100,
	//	},
	//	Title:         "post 1",
	//	Content:       "p1 content",
	//	UserID:        1,
	//	CommentsCount: 0,
	//	Status:        "无评论",
	//}
	//db.Create(&p)
	//db.Model(&p).Select("id", "title", "content", "user_id", "comments_count", "status").Save(&p)
	//db.Delete(&p)

	c := Comment{
		Model: gorm.Model{
			ID: 2000,
		},
		Content: "Comment 1 +",
		PostID:  100,
	}
	//db.Create(&c)
	//db.Model(&c).Select("id", "content", "post_id").Save(&c)
	db.Delete(&c)
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	err = syncPostsCount(tx, p)
	if err != nil {
		return err
	}
	return nil
}
func (p *Post) AfterUpdate(tx *gorm.DB) (err error) {
	err = syncPostsCount(tx, p)
	if err != nil {
		return err
	}
	return nil
}
func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	err = syncPostsCount(tx, p)
	if err != nil {
		return err
	}
	return nil
}

func syncPostsCount(tx *gorm.DB, p *Post) (err error) {

	//查询某用户的文章数
	var maxPostsCount int64
	err1 := tx.Model(&Post{}).Where("user_id=?", p.UserID).Count(&maxPostsCount).Error
	if err1 != nil {
		fmt.Println("err1:", err1)
		return err1
	}
	fmt.Println("maxPostsCount:", maxPostsCount)

	// 更新用户表文章总数字段
	err2 := tx.Model(&User{}).Where("  id=?", p.UserID).Update("posts_count", maxPostsCount).Error
	if err2 != nil {
		fmt.Println("err2:", err2)
		return err2
	}

	return nil
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	err = syncCommentsCount(tx, c)
	if err != nil {
		return err
	}
	return nil
}
func (c *Comment) AfterUpdate(tx *gorm.DB) (err error) {
	err = syncCommentsCount(tx, c)
	if err != nil {
		return err
	}
	return nil
}
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	err = syncCommentsCount(tx, c)
	if err != nil {
		return err
	}
	return nil
}

func syncCommentsCount(tx *gorm.DB, c *Comment) (err error) {
	//查询某文章最大评论数
	var maxCommentsCount int64
	err3 := tx.Model(&Comment{}).Where("  post_id=?", c.PostID).Count(&maxCommentsCount).Error
	if err3 != nil {
		fmt.Println("err3:", err3)
		return err3
	}
	fmt.Println("maxCommentsCount:", maxCommentsCount)
	//更新文章表评论总数字段以及状态字段
	var statusStr = "无评论"
	if maxCommentsCount > 0 {
		statusStr = "有评论"
	}
	err4 := tx.Model(&Post{}).Where("  id=?", c.PostID).Updates(map[string]interface{}{"comments_count": maxCommentsCount, "status": statusStr}).Error
	if err4 != nil {
		fmt.Println("err4:", err4)
		return err4
	}

	return nil
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
		Title:         "hello world",
		Content:       "hello world,hahahahaha",
		Comments:      []Comment{c1, c2},
		CommentsCount: 2,
		Status:        "有评论",
	}
	p2 := Post{
		Model: gorm.Model{
			ID: 101,
		},
		Title:         "hello Beijing",
		Content:       "hello Beijing,hahahahaha",
		Comments:      []Comment{c3, c4},
		CommentsCount: 2,
		Status:        "有评论",
	}
	p3 := Post{
		Model: gorm.Model{
			ID: 102,
		},
		Title:         "hello Shanghai",
		Content:       "hello Shanghai,hahahaha",
		CommentsCount: 0,
		Status:        "无评论",
	}

	u1 := User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:       "Alice",
		Posts:      []Post{p1, p2},
		PostsCount: 2,
	}
	u2 := User{
		Model: gorm.Model{
			ID: 2,
		},
		Name:       "Bob",
		Posts:      []Post{p3},
		PostsCount: 1,
	}
	db.Create(&u1)
	db.Create(&u2)
}
