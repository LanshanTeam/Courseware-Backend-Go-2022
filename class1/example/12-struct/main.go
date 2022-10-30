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
	fmt.Println(a.name)                   // wang

	fmt.Println(checkPassword2(&a, "1024")) // true
	fmt.Println(a.name)                     // test
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
