package stage3

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var myschema = `
CREATE TABLE employees  (
    id  		bigint unsigned auto_increment primary key,
    name 		longtext 	null,
    department  longtext 	null,
    salary     	float   	null
);

CREATE TABLE books  (
    id  	bigint unsigned auto_increment primary key,
    title 	longtext	null,
    author  longtext	null,
    price  	float		null
)`

type Employees struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float32 `db:"salary"`
}

type Books struct {
	Id     int
	Title  string
	Author string
	Price  float32
}

func Run1() {
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/metanote?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true")
	if err != nil {
		log.Fatalln(err)
	}

	// 初始化数据
	//db.MustExec(myschema)
	//tx := db.MustBegin()
	//tx.MustExec("INSERT INTO employees (name, department,salary) VALUES (?, ?, ?)", "张三", "技术部", "7000")
	//tx.NamedExec("INSERT INTO employees (name, department,salary) VALUES (:name, :department, :salary)", &Employees{Name: "李四", Department: "技术部", Salary: 11000})
	//tx.Exec("INSERT INTO employees (name, department,salary) VALUES (?,?,?)", "王五", "技术部", 9000)
	//tx.Exec("INSERT INTO employees (name, department,salary) VALUES (?,?,?)", "麻六", "销售部", 5000)
	//tx.MustExec("INSERT INTO books (title, author,price) VALUES (?, ?, ?)", "故事1", "张三", "35")
	//tx.MustExec("INSERT INTO books (title, author,price) VALUES (?, ?, ?)", "故事2", "张三", "18")
	//tx.MustExec("INSERT INTO books (title, author,price) VALUES (?, ?, ?)", "故事3", "李四", "55")
	//tx.MustExec("INSERT INTO books (title, author,price) VALUES (?, ?, ?)", "故事4", "王五", "63")
	//tx.Commit()

	// 查询 employees 表中所有部门为 "技术部" 的员工信息
	//var employees []Employees
	//sqlStr := "select * from employees where department = ?"
	//err2 := db.Select(&employees, sqlStr, "技术部")
	//if err2 != nil {
	//	fmt.Printf("query failed, err:%v\n", err)
	//	return
	//}
	//fmt.Printf("employees:%+v", employees)

	//查询 employees 表中工资最高的员工信
	//var employee Employees
	//sqlStr := "SELECT * FROM employees ORDER BY salary DESC LIMIT 1"
	//err2 := db.Get(&employee, sqlStr)
	//if err2 != nil {
	//	fmt.Printf("query failed, err:%v\n", err)
	//	return
	//}
	//fmt.Printf("employee:%+v", employee)

	//查询价格大于50元的书籍
	var books []Books
	sqlStr := "select * from books where price > ?"
	err2 := db.Select(&books, sqlStr, 50)
	if err2 != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("books:%+v", books)

}
