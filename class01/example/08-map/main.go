package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)            // map[one:1 two:2]
	fmt.Println(len(m))       // 2
	fmt.Println(m["one"])     // 1
	fmt.Println(m["unknown"]) // 0

	r, ok := m["unknown"]
	fmt.Println(r, ok) // 0 false

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 3}

	fmt.Println(m2, m3)
}
