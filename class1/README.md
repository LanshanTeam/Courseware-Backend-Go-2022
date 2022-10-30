# 蓝山工作室——Golang第一节课

代码地址：https://github.com/LanshanTeam/Courseware-Backend-Go-2022

> 以后每次更新代码之后同学们可以直接通过 git 进行更新

## 前言

可能有些同学已经有基础了，但是没有基础也没有关系，我们会从最基础的语法讲起，并在本节课学习到 Golang 的一些库函数操作，在最后会让大家上手一个简单的小项目。

## 基础语法

### 你好，Lanshan!

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Lanshan!")
}
```

### var

#### 变量

##### 变量类型
变量（Variable）的功能是存储数据。不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见变量的数据类型有：整型、浮点型、布尔型等。

Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。

##### 变量声明
Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用。

##### 变量的初始化
Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如： 整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。 布尔型变量默认为false。 切片、函数、指针变量的默认为nil。

```go
var 变量名 类型 = 表达式

var a = "initial" // 类型推导，不指定类型自动判断

var b, c int = 1, 2 // 一次初始化多个变量

var d = true

var e float64 // 普通声明未赋值

f := float32(e) // 短声明

g := a + "apple"
fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
fmt.Println(g)                // initialapple
```

#### 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。

```go
const s string = "constant"
const h = 500000000
const i = 3e20 / h
fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
```

#### 练习 

https://www.acwing.com/problem/content/619/

### for

```go
for init statement; condition expression; post statement {
    // 这里是中间循环体
}
```

`statement`是单次表达式，循环开始时会执行一次这里

`expression`是条件表达式，即循环条件，只要满足循环条件就会执行中间循环体。

`statement`是末尾循环体，每次执行完一遍中间循环体之后会执行一次末尾循环体

执行末尾循环体后将再次进行条件判断，若条件还成立，则继续重复上述循环，当条件不成立时则跳出当下for循环

```go
package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println("loop")
		break // 跳出循环
	}
	
	// 打印7、8
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
			// 当n模2为0时不打印，进到下一次的循环
		}
		fmt.Println(n)
	}
	// 直到i>3
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
  // for 循环嵌套
  for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
```

#### 练习 

https://www.acwing.com/problem/content/723/

### if

```go
if 条件表达式 {
	//当条件表达式结果为true时，执行此处代码   
}

if 条件表达式 {
    //当条件表达式结果为true时，执行此处代码  
} else {
    //当条件表达式结果为false时，执行此处代码  
}
```

```go
package main

import "fmt"

func main() {
	// 条件表达式为false，打印出"7 is odd"
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 条件表达式为ture，打印出"8 is divisible by 4"
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 短声明，效果等效于
	//num := 9
	//if num < 0{
	//	...
	//}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```

#### 练习

https://www.acwing.com/problem/content/664/

### switch

当分支过多的时候，使用if-else语句会降低代码的可阅读性，这个时候，我们就可以考虑使用switch语句

- switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
- switch 语句在默认情况下 case 相当于自带 break 语句，匹配一种情况成功之后就不会执行其它的case，这一点和 c/c++ 不同
- 如果我们希望在匹配一条 case 之后，继续执行后面的 case ，可以使用 fallthrough

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		// 在此打印"two"并跳出
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	// 根据现在的时间判断是上午还是下午
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
```

#### 练习

https://www.acwing.com/problem/content/661/

### array

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

```go
package main

import "fmt"

func main() {
	// 声明了长度为5的数组，数组中的每一个元素都是int类型
	var a [5]int
	// 给数组a的第4位元素赋值为100
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	// 在给数组声明的同时赋值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 声明二位数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
```

#### 练习

https://www.acwing.com/problem/content/739/ 

https://www.acwing.com/problem/content/743/

### slice

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

```go
var s []int
```
类似与声明一个数组，只不过不用填写它的长度

值得一提的是，切片必须先初始化才能使用！

```go
package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	// 使用append在尾部添加元素
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	c := make([]string, len(s))
	// 将s复制给c
	copy(c, s)
	fmt.Println(c) // [a b c d e f]
	
	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]
}
```

#### 练习 

https://www.acwing.com/problem/content/767/

### map

在 Go 语言中，map 是散列表的引用，map 的类型是 map[K]V ，其中 K 和 V 是字典的键和值对应的数据类型。map 中所有的键都拥有相同的数据类型，同时所有的值也都拥有相同的数据类型，但是键的类型和值的类型不一定相同。键的类型 K ，必须是可以通过操作符 == 来进行比较的数据类型，所以 map 可以检测某一个键是否存在。

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)           // map[one:1 two:2]
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    // 1
	fmt.Println(m["unknown"]) // 0

	r, ok := m["unknown"]
	fmt.Println(r, ok) // 0 false

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)
}
```

### range

用于遍历，注意以下两点

1. range在map中遍历顺序是随机的，多次遍历的结果可能不同

2. range在数组中是从下标0开始递增遍历的，多次遍历的结果是相同的

```go
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9

	m := map[string]string{"a": "A", "b": "B", "c": "C"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println("key", k)
	}
}
```

#### 练习

https://www.acwing.com/problem/content/724/（使用range）

### func

函数是指一段可以直接被另一段程序或代码引用的程序或代码，一个较大的程序一般应分为若干个程序块，每一个模块用来实现一个特定的功能。

```go
package main

import "fmt"

func add(a int, b int) int {
	// 返回a+b的和
	return a + b
}

// 若类型相同，允许这样写
func add2(a, b int) int {
	return a + b
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
	v, ok = exists(map[string]string{"a": "A"}, "b")
	fmt.Println(v, ok) //   false
}
```

#### 练习

https://www.acwing.com/problem/content/821/

https://www.acwing.com/problem/content/811/

### point

指针也是一个变量，但它是一种特殊的变量，因为它存储的数据不仅仅是一个普通的值，如简单的整数或字符串，而是另一个变量的内存地址。

```go
package main

import "fmt"

func add2(n int) {
	n += 2
}

func add2pt(n *int) {
	*n += 2
}

func swap(a int, b int) {
	a, b = b, a
}

//函数参数为指针类型
func swapWithPt(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2pt(&n)
	fmt.Println(n) // 7

	a, b := 2, 3
	swap(a, b)
	fmt.Println(a, b)
	swapWithPt(&a, &b)
	fmt.Println(a, b)
}
```

### struct

Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。

```go
package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "wang"}
	c.password = "1024"
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d)               // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "1024")) // true
	fmt.Println(a.name) // wang

	fmt.Println(checkPassword2(&a, "1024")) // true
	fmt.Println(a.name) // test
}

// 值传递
func checkPassword(u user, password string) bool {
	u.name = "test"
	return u.password == password
}

// 指针传递
func checkPassword2(u *user, password string) bool {
	u.name = "test"
	return u.password == password
}
```

### struct-method

在 Go 语言中，结构体就像是类的一种简化形式，那么面向对象程序员可能会问：类的方法在哪里呢？在 Go 中有一个概念，它和方法有着同样的名字，并且大体上意思相同：Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。

```go
package main

import "fmt"

type user struct {
	name     string
	password string
}

//func (u user) checkPassword(password string) bool {
//	return u.password == password
//}
//
//func (u user) resetPassword(password string) {
//	u.password = password
//}

func (u *user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user{name: "wang", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048")) // true
}
```

#### 练习

https://www.acwing.com/problem/content/818/（使用方法，不使用函数）

### interface

很多面向对象的语言都有接口这个概念，Go语言的接口的独特之处在于它是隐式实现。换句话说，对于一个具体的类型，无须声明它实现了哪些接口，只要该类型提供了接口所必须的方法即可。这种设计让你无须改变已有类型的实现，就可以为这些类型扩展新的接口，对于那些不能修改包的类型，这一点特别有用。

Go语言中提供了一种类型叫做接口类型。接口是一种抽象类型，它并没有暴露所含数据的布局或内部结构，当然也没有哪些数据的基本操作，它所提供的仅仅是一些方法而已。如果你拿到一个接口类型的值，你无从知道它是什么，你能知道的仅仅是它能做什么，也就是说，仅仅能知道它提供了哪些方法。

一个接口类型定义了一套方法，如果一个具体的类型要实现该接口，那么必须实现该接口类型定义中的所有方法。

```go
package main

import "fmt"

// Phone 接口名为Phone，接口内有speak与read方法
type Phone interface {
	speak()
	getPrice()
}

// IPhone 以下结构体可以分别设置自己的属性
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

// 手机讲话区域（对相关的接口进行实现）
func (P IPhone) speak() {
	fmt.Println("Hi，我是 Siri！")
}

func (P Oppo) speak() {
	fmt.Println("我是", P.name)
}

func (P Mi) speak() {
	fmt.Println("大家好,我是小爱童鞋!")
}

func (P IPhone) getPrice() {
	fmt.Println("My price is", P.price)
}

func show(myPhone Phone) {
	myPhone.speak()
	myPhone.getPrice()
}

func main() {
	// 将新建对象传入展示大舞台,大舞台代码不变,展示不同的效果
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
```

没有任何方法的接口就是空接口,实际上每个类型都实现了空接口,所以空接口类型可以接受任何类型的数据

### errors

errors 库函数

```go
package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	u, err := findUser([]user{{"wang", "1024"}, {"yuan", "1212"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name) // wang

	if u, err := findUser([]user{{"wang", "1024"}, {"yuan", "1212"}}, "li"); err != nil {
		fmt.Println(err) // not found
		return
	} else {
		fmt.Println(u.name)
	}
}
```

### string

string 库函数

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))                // true
	fmt.Println(strings.Count(a, "l"))                    // 2
	fmt.Println(strings.HasPrefix(a, "he"))               // true
	fmt.Println(strings.HasSuffix(a, "llo"))              // true
	fmt.Println(strings.Index(a, "ll"))                   // 2
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo
	fmt.Println(strings.Repeat(a, 2))                     // hellohello
	fmt.Println(strings.Replace(a, "e", "E", -1))         // hEllo
	fmt.Println(strings.Split("a-b-c", "-"))              // [a b c]
	fmt.Println(strings.ToLower(a))                       // hello
	fmt.Println(strings.ToUpper(a))                       // HELLO
	fmt.Println(len(a))                                   // 5
}
```

#### 练习

https://www.acwing.com/problem/content/772/

### fmt

fmt 库函数

```go
package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n) // hello 123
	fmt.Println(p)    // {1 2}

	fmt.Printf("s=%v\n", s)  // s=hello
	fmt.Printf("n=%v\n", n)  // n=123
	fmt.Printf("p=%v\n", p)  // p={1 2}
	fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}

	f := 3.141592653
	fmt.Println(f)          // 3.141592653
	fmt.Printf("%.2f\n", f) // 3.14
}
```

### time

time 库函数

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 2022-03-27 18:04:59.433297 +0800 CST m=+0.000087933
	t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
	fmt.Println(t)                                                  // 2022-03-27 01:25:36 +0000 UTC
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()) // 2022 March 27 1 25
	fmt.Println(t.Format("2006-01-02 15:04:05"))                    // 2022-03-27 01:25:36
	diff := t2.Sub(t)
	fmt.Println(diff)                           // 1h5m0s
	fmt.Println(diff.Minutes(), diff.Seconds()) // 65 3900
	t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t)    // true
	fmt.Println(now.Unix()) // 1648738080
}
```

### strconv

strconv 库函数

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 1.234

	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n) // 111

	n, _ = strconv.ParseInt("0x1000", 0, 64)
	fmt.Println(n) // 4096

	n2, _ := strconv.Atoi("123")
	fmt.Println(n2) // 123

	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err) // 0 strconv.Atoi: parsing "AAA": invalid syntax
}
```

## 年轻人的第一个GoProject

### 猜数游戏

#### v1

由于是猜数嘛，我们肯定需要先有可以猜的数。所以我们需要生成一个随机数

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}

```

#### v2

有的同学可能会发现，虽然是产生了随机数，但是每一次生成的数字都是一样的。这是因为我们生成随机数的种子没有改变，需要让种子发生变化才能使每次生成的随机数不同

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	// 使用一直在不断变化的时间作为我们的种子
	rand.Seed(time.Now().UnixNano())
	// 设置种子之后产生一个最大为100的整形
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}
```

#### v3

实现输入我们猜的数字

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	var guess int
	// 输入我们猜的数字
	_, err := fmt.Scanf("%d", &guess)
	// Go语言中处理错误的方法
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}
	fmt.Println("You guess is", guess)
}
```

#### v4

实现完整的猜数逻辑

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	var guess int
	_, err := fmt.Scanf("%d", &guess)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}
	fmt.Println("You guess is", guess)
	// 通过 if-else 来看你是否猜对，若没猜对告诉你猜大了还是猜小了
	if guess > secretNumber {
		fmt.Println("Your guess is bigger than the secret number. Please try again")
	} else if guess < secretNumber {
		fmt.Println("Your guess is smaller than the secret number. Please try again")
	} else {
		fmt.Println("Correct, you Legend!")
	}
}
```

#### v5

加入for循环，项目实现完成

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// 作弊模式
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	// 通过一个 for 循环实现一直猜数，直到猜中
	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
```

## 作业

### LV0 

将今天讲的知识巩固一遍并自己敲一遍示例代码以及练习题。

### LV1 

给定一个长度为 n 的数组 a 以及两个整数 l 和 r，请你编写一个**函数**，将 a[l] ∼ a[r] 从小到大排序。

输出排好序的数组 a 。

#### 输入格式

第一行包含三个整数 n，l，r。

第二行包含 n 个整数，表示数组 a。

#### 输出格式

共一行，包含 n 个整数，表示排序完成后的数组 a。

#### 数据范围

0 ≤ l ≤ r < n ≤ 1000

#### 输入样例

```
5 2 4
4 5 1 3 2
```

#### 输出样例

```
4 5 1 2 3
```

### LV2 

使用结构体实现一个豆瓣电影详细页，并能够通过控制台获得内容，可以自定义一些花里胡哨的功能，实现越多功能越好。

#### 样例图片

![](https://picture.lanlance.cn/i/2022/10/30/635e6206f12ab.png)

#### 样例代码

```go
package main

import "fmt"

type Movie struct {
	Name string
}

func main() {
	m := Movie{Name: "西线无战事"}
	fmt.Printf("请输入你的命令\n1.获得名字\n2.退出程序\n")
	var option int
	for {
		fmt.Scanf("%d", &option)
		if option == 1 {
			fmt.Println(m.Name)
		} else if option == 2 {
			return
		}
	}
}
```

#### 样例输出

```shell
请输入你的命令
1.获得名字
2.退出程序
1
西线无战事
2

进程 已完成，退出代码为 0
```

### LV3

使用接口实现一个简单的英雄联盟

#### 基础要求

- 英雄可以进行攻击
- 英雄可以释放技能
- 英雄可以使用道具

#### 进阶要求

比如说释放技能会减少自己的蓝量且减少攻击目标的血量，使用道具会有攻击力加强等，如何实现可以自行思考。

#### 样例代码

```go
package main

import "fmt"

type Hero interface {
	Say()
}

type Yasuo struct {
	HP int
}

type Vayne struct {
	HP int
}

func (Y Yasuo) Say() {
	fmt.Println("I'm Yasuo")
}

func (V Vayne) Say() {
	fmt.Println("I'm Vayne")
}

func main() {
	fmt.Printf("请输入你的命令\n1.选择你的英雄\n2.获取英雄列表\n3.退出程序\n")
	var option int
	if option == ...
  ...
}
```

#### 样例输出

```shell
请输入你的命令
1.选择你的英雄
2.获取英雄列表
3.退出程序
2
Yasuo
Vayne
3

进程 已完成，退出代码为 0
```

> 当项目设计的较为复杂是，菜单可以不止一级

### LV4 
按照自己的想法将猜数游戏升级到v6并说明你的升级内容
>可以是任何的升级，请发挥你们的想象力。如果可以甚至可以将你的猜数游戏部署到网页上(如果真的实现了，第一个人奖励一杯奶茶哦)


作业完成后将作业 GitHub 地址发送至 **yuanxinhao@lanshan.email** ，若对 GitHub 的使用有问题，可以先网上寻找解决方法，实在不行可以私信学长。

**截止时间**：下一次上课之前
