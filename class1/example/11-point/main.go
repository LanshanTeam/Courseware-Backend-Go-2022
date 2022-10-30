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
