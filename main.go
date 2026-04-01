package main

import (
	"MetaNode/stage3"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GLOBAL_DB *gorm.DB

func main() {

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
	GLOBAL_DB = db

	//stage1.Stage1Valid()
	//stage2.Stage2Valid()

	//stage3.Run()
	//stage3.Run1()
	stage3.Run2()
}
