package main

import (
	"context"
	"fmt"
	"github.com/eiannone/keyboard"
	"runtime"
	"time"
)

const (
	bulletNum = 10 // 装弹人数
	aimNum    = 5  // 瞄准人数
	shotNum   = 3  // 发射人数
)

// 思考：如果这里不设置长度(长度为0)，打印出最后的 goroutine 数量有3个，为什么设置成 1 之后只会有 2 个
var (
	BulletArtillery = make(chan struct{}, 1)
	AimArtillery    = make(chan struct{}, 1)
	ShotArtillery   = make(chan struct{}, 1)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fire(ctx)
	var (
		input rune
		err   error
	)
	for input != 'q' {
		input, _, err = keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("停止前 Goroutine 数量：%d\n", runtime.NumGoroutine())
	cancel()
	time.Sleep(time.Second * 2) // 等待手下停止打炮
	fmt.Println("打炮结束.")
	// 会打印出 2 个，一个是 main goroutine，另一个是 keyboard 监听事件的 goroutine
	fmt.Printf("停止后 Goroutine 数量：%d\n", runtime.NumGoroutine())
}

func fire(ctx context.Context) {
	for i := 0; i < bulletNum; i++ {
		go func(pos int) { // 思考，为什么要定义个 pos? 而不是直接使用 i
			for {
				select {
				case <-BulletArtillery:
					time.Sleep(time.Second)
					fmt.Print("装弹 ->")
					AimArtillery <- struct{}{}
				case <-ctx.Done():
					fmt.Printf("装弹手%d：回家了\n", pos)
					return
				}
			}
		}(i)
	}
	for i := 0; i < aimNum; i++ {
		go func(pos int) {
			for {
				select {
				case <-AimArtillery:
					time.Sleep(time.Second)
					fmt.Print(" 瞄准 ->")
					ShotArtillery <- struct{}{}
				case <-ctx.Done():
					fmt.Printf("瞄准手%d：回家了\n", pos)
					return
				}
			}
		}(i)
	}
	for i := 0; i < shotNum; i++ {
		go func(pos int) {
			for {
				select {
				case <-ShotArtillery:
					time.Sleep(time.Second)
					fmt.Println(" 发射！")
					BulletArtillery <- struct{}{}
				case <-ctx.Done():
					fmt.Printf("发射手%d：回家了\n", pos)
					return
				}
			}
		}(i)
	}
	// 开始装弹
	BulletArtillery <- struct{}{}
}
