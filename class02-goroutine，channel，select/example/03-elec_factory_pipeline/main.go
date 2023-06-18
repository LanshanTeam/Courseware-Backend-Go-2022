package main

import (
	"fmt"
	"time"
)

type Item struct {
	Name  string
	Count int
}

func main() {
	pipeline := make(chan Item, 10) // 一条可以放 10 个 item 的流水线
	go func() {
		for {
			time.Sleep(1 * time.Second)
			pipeline <- Item{
				Name:  "螺丝",
				Count: 5,
			}
		}
	}()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			pipeline <- Item{
				Name:  "齿轮",
				Count: 3,
			}
		}
	}()
	for {
		item := <-pipeline
		fmt.Printf("%#v\n", item)
	}
}
