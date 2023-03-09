package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/websocket", websocketFc)
	r.Run(":8080")
}

func websocketFc(c *gin.Context) {
	//设置Upgrader
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//升级协议，返回ws连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"status": 500,
			"info":   "failed",
		})
		return
	}

	//接受数据
	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
			} else {
				log.Println(string(p))
			}
		}
	}()

	//发送数据 这里使用WriteJSON，如果对websocket熟悉 ，可以自行构造消息
	go func() {
		for i := 0; ; i++ {
			conn.WriteJSON(gin.H{
				"time": time.Now().Format(time.RFC3339),
				"No":   i,
			})
			time.Sleep(time.Second * 3)
		}
	}()
}
