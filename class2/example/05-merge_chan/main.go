package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		ch1 = make(chan struct{})
		ch2 = make(chan struct{})
		ch3 = make(chan struct{})
	)
	go handleCh1(ch1)
	go handleCh2(ch2)
	go handleCh3(ch3)
	for {
		select {
		case _ = <-ch1:
			fmt.Println("get from ch1")
		case _ = <-ch2:
			fmt.Println("get from ch2")
		case _ = <-ch3:
			fmt.Println("get from ch3")
		}
	}
}

func handleCh1(ch1 chan struct{}) {
	for {
		time.Sleep(3 * time.Second)
		ch1 <- struct{}{}
	}
}

func handleCh2(ch2 chan struct{}) {
	for {
		time.Sleep(4 * time.Second)
		ch2 <- struct{}{}
	}
}

func handleCh3(ch3 chan struct{}) {
	for {
		time.Sleep(2 * time.Second)
		ch3 <- struct{}{}
	}
}
