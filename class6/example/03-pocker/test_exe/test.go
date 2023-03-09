package main

import (
	"fmt"
	"net/http"
)

var (
	badSlice = make([]int, 0)
)

func testHelloWorld() {
	fmt.Println("hello world!")
}

func testHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	panic(http.ListenAndServe(":8080", http.DefaultServeMux))
}

func testSlice() {
	i := 0
	for {
		badSlice = append(badSlice, 0)
		i++
		if i%10_000 == 0 {
			fmt.Println("length of badSlice:", len(badSlice))
		}
	}
}

func main() {
	testSlice()
}
