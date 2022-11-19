package main

import (
	"fmt"
	"time"
)

func main() {
	// 好康的 := 欢迎来我家玩()
	// 打电动() // 阿伟想在去杰哥的路上打电动
	// fmt.Println(好康的)

	// // 方式一：
	// go 打电动()
	// 好康的 := 欢迎来我家玩()
	// fmt.Println(好康的)

	// 方式二：
	var 杰哥答应的好康的 = make(chan string)
	go func() {
		杰哥答应的好康的 <- 欢迎来我家玩()
	}()
	打电动()
	好康的 := <-杰哥答应的好康的 // 等待杰哥的好康的
	fmt.Println(好康的)
}

func 欢迎来我家玩() string {
	// 花费 5s 前往杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郎"
}

func 打电动() {
	fmt.Println("输了啦，都是你害的")
}
