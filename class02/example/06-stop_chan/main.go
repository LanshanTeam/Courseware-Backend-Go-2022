package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		ch1  = make(chan struct{})
		stop = make(chan struct{})
	)
	go handleCh1(ch1)
	go func() {
		time.Sleep(10 * time.Second)
		stop <- struct{}{}
	}()
LOOP:
	for {
		select {
		case _ = <-ch1:
			fmt.Println("get from ch1")
		case _ = <-stop:
			break LOOP
		}
	}
}

func handleCh1(ch1 chan struct{}) {
	for {
		time.Sleep(3 * time.Second)
		ch1 <- struct{}{}
	}
}
