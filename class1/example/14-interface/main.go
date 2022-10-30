package main

import "fmt"

type Phone interface {
	speak()
	getPrice()
}

type IPhone struct {
	name  string
	price int
}

type Oppo struct {
	name  string
	price int
	color string
}

type Mi struct {
	f bool
}

func (P IPhone) speak() {
	fmt.Println("Hi，I'm Siri！")
}

func (P Oppo) speak() {
	fmt.Println("I'm", P.name)
}

func (P Mi) speak() {
	fmt.Println("Hi! I'm XiaoAi")
}

func (P IPhone) getPrice() {
	fmt.Println("My price is", P.price)
}

func show(myPhone Phone) {
	myPhone.speak()
	myPhone.getPrice()
}

func judgeType(q interface{}) {
	temp, ok := q.(string)
	if ok {
		fmt.Println("类型转换成功!", temp)
	} else {
		fmt.Println("类型转换失败!", temp)
	}
}

func main() {
	myPhone := IPhone{
		name:  "华为13远峰蓝",
		price: 8848,
	}
	show(myPhone)

	var i interface{}
	i = 10
	v, ok := i.(int)
	fmt.Println(v, ok)
	v2 := i.(int)
	fmt.Println(v2)
}
