package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = &sync.Mutex{}
	lock2 = &sync.Mutex{}
)

func main() {
	go repeater1()
	go repeater1()
	go repeater1()
	go repeater2()
	go repeater1()
	go repeater1()
	go repeater1()
	go repeater2()
	go repeater2()
	go repeater2()
	repeater1()
}

func repeater1() {
	for {
		time.Sleep(time.Second)
		// fmt.Println("over.")
		lock1.Lock()
		lock2.Lock()
		fmt.Print("o")
		fmt.Print("v")
		fmt.Print("e")
		fmt.Print("r")
		fmt.Println(".")
		lock1.Unlock()
		lock2.Unlock()
	}
}

func repeater2() {
	for {
		time.Sleep(time.Second)
		// fmt.Println("over.")
		lock1.Lock()
		lock2.Lock()
		fmt.Print("o")
		fmt.Print("v")
		fmt.Print("e")
		fmt.Print("r")
		fmt.Println(".")
		lock2.Unlock()
		lock1.Unlock()
	}
}
