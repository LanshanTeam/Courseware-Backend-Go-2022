package main

import (
	"Courseware-Backend-Go-2022/class4/gorm_example/dsn"
	"Courseware-Backend-Go-2022/class4/gorm_example/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func initDB() {
	var err error
	//dsn := "user:pass@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := dsn.DSN
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("connect success:")
}

func insert(st model.Student) {
	res := db.Table("student").Create(&st)
	if res.Error != nil {
		log.Println("insert err:", res.Error)
		return
	}
	log.Println("insert success")
}

func selectStudent() {
	st := model.Student{}
	//查询第一条记录
	db.First(&st)
	log.Println("success find:", st)

	//如果主键是数字类型，可以使用内联条件检索
	db.First(&st, 10)
	// SELECT * FROM users WHERE id = 10;

	db.First(&st, "10")
	// SELECT * FROM users WHERE id = 10;

	db.Find(&st, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);

	//如果主键是字符串（例如像 uuid），查询将被写成这样：
	//db.First(&st, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	//当目标对象有一个主要值时，将使用主键构建条件，例如：
	st.Id = 1
	db.First(&st)
	st = model.Student{}
	// SELECT * FROM users WHERE id = 10;

	var result model.Student
	db.Model(model.Student{Id: 10}).First(&result)
	// SELECT * FROM users WHERE id = 10;

	//检索全部记录
	var sts []model.Student
	db.Find(&sts)

	//随机获取一条记录
	db.Take(&st)
	log.Println("success take:", st)

	//获取最后一条记录
	db.Last(&st)
	log.Println("success last:", st)

}

func main() {
	initDB()

	st := model.Student{
		Name: "老王",
		Sex:  "男",
		Age:  "8",
	}
	insert(st)

}
