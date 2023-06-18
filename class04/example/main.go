package main

import (
	"Courseware-Backend-Go-2022/class4/example/dsn"
	"Courseware-Backend-Go-2022/class4/example/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 定义一个全局对象db
var db *sql.DB

func initDB() {
	var err error
	// 设置一下dns
	//dsn := "username:password@tcp(127.0.0.1:3306)/database_name"
	dsn := dsn.DSN
	// 打开mysql驱动
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}

// QueryRowById 查询单条数据
func QueryRowById(id int) {
	sqlStr := "select id, name, age from student where id=?"
	var st model.Student
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, id).Scan(&st.Id, &st.Name, &st.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	log.Printf("QueryRowById:   id:%d name:%s age:%d\n", st.Id, st.Name, st.Age)
}

// MultiQueryById 查询多行数据
func MultiQueryById(id int) {
	sqlStr := "select id, name, age from student where id > ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var st model.Student
		err := rows.Scan(&st.Id, &st.Name, &st.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		log.Printf("MultiQueryById: id:%d name:%s age:%d\n", st.Id, st.Name, st.Age)
	}
}

// InsertStudent 插入数据
func InsertStudent(st model.Student) {
	sqlStr := "insert into student(name,age) values (?,?)"
	_, err := db.Exec(sqlStr, st.Name, st.Age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}

// UpdateStudentAgeById 更新数据
func UpdateStudentAgeById(st model.Student) {
	sqlStr := "update student set age=? where id=?"
	_, err := db.Exec(sqlStr, st.Age, st.Id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	log.Println("update success")
}

// DeleteStudentById 删除数据
func DeleteStudentById(id int) {
	sqlStr := "delete from student where id=?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	log.Println("delete success")
}

func main() {
	//初始化连接
	initDB()
	//查询id为1和2的学生
	QueryRowById(1)
	QueryRowById(2)
	//查询所有id大于1的学生
	MultiQueryById(1)

	//插入数据
	st := model.Student{
		Name: "小王",
		Sex:  "男",
		Age:  98,
	}
	InsertStudent(st)
	//更新数据
	st.Age = 100
	UpdateStudentAgeById(st)
	//删除数据
	DeleteStudentById(4)
}
