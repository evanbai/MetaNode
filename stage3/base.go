package stage3

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/metanote?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(dst...)

	return db
}

func Run() {
	db := InitDB(&Student{}, &Accounts{}, &Transactions{})

	//插入一条记录
	//student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	//result := db.Create(&student)
	//fmt.Println(student)
	//fmt.Println(result.RowsAffected)

	//批量插入记录
	//var students = []Student{{Name: "张三", Age: 20, Grade: "三年级"}, {Name: "李四", Age: 12, Grade: "一年级"}, {Name: "王五", Age: 22, Grade: "四年级"}}
	//db.Create(students)

	// 条件查询
	//students := []Student{}
	//queryResult := db.Where("age > ?", 18).Find(&students)
	//fmt.Println(students)
	//fmt.Println(queryResult.RowsAffected)

	//更新
	//db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	//删除
	//db.Unscoped().Where("age < ?", 15).Delete(&Student{})

	// 事务 - 初始化
	//var accounts = []Accounts{{Name: "A", Balance: 514}, {Name: "B", Balance: 877}, {Name: "C", Balance: 323}}
	//db.Create(&accounts)

	// 事务 - 交易
	tx := db.Begin()
	var accountA Accounts
	var accountB Accounts
	tx.Where("name = ?", "A").Find(&accountA)
	tx.Where("name = ?", "B").Find(&accountB)
	// A 向 B 支付 100元
	var amount float32 = 100
	if accountA.Balance < amount {
		tx.Rollback()
	}
	accountA.Balance = accountA.Balance - amount
	tx.Save(&accountA)
	accountB.Balance = accountB.Balance + amount
	tx.Save(&accountB)
	transaction := Transactions{FromAccountId: accountA.ID, ToAccountId: accountB.ID, Amount: amount}
	tx.Save(&transaction)
	tx.Commit()

}

type Student struct {
	gorm.Model
	Name  string
	Age   uint
	Grade string
}

type Accounts struct {
	gorm.Model
	Name    string
	Balance float32
}

type Transactions struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        float32
}
