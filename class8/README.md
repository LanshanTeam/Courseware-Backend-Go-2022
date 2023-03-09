# 网络协议进阶

### Websocket

![uTools_1678116912475](http://typora.fengxiangrui.top/1678116918.png)

#### 什么是websocket

- WebSocket是HTML5下一种新的协议（websocket协议本质上是一个**基于tcp的协议**）
- 它实现了浏览器与服务器**全双工通信**，能更好的节省服务器资源和带宽并达到实时通讯的目的
- Websocket是一个**持久化**的协议



#### 为什么需要websocket

##### HTTP

这需要从HTTP说起

众所周知， HTTP协议是**无状态协议**，HTTP协议自身不对请求和响应之间的通信状态进行保存

也就是说在HTTP协议对于发送过的请求或响应都**不做持久化处理**

并且，HTTP**协议是无连接的**

无连接的含义是限制每次连接只处理一个请求

服务器处理完客户的请求，并收到客户的应答后，即断开连接

然而，HTTP协议存在以下问题：

- **无状态：**因此每次会话完成后，服务端都不知道下一次客户端身份
- **无连接：**HTPP协议采用一次请求，一次响应，每次请求和响应就携带有大量的header，因此，效率也更低下
- **客户端不能主动发送数据**

举个例子，如果我们想实现QQ消息及时推送该怎么办?

##### 轮询

对于以上问题，可以使用**轮询**解决

但轮询解决方案通常对服务端压力较大，延迟也较高

于是就该websocket上场了

##### websocket

> 一旦WebSocket连接建立后，后续数据都以**帧序列**的形式传输。
>
> 在客户端断开WebSocket连接或Server端中断连接前，**不需要**客户端和服务端**重新发起连接请求**。
>
> 在海量并发及客户端与服务器交互负载流量大的情况下，极大的**节省了网络带宽资源**的消耗，有**明显的性能优势**，且客户端发送和接受消息是在同一个持久连接上发起，实现了“真·长链接”，**实时性优势明显**。

![uTools_1678117644702](http://typora.fengxiangrui.top/1678117650.png)

#### 如何使用websocket

在go中，有现成的SDK可以使用

- [gorilla](https://github.com/gorilla)/**[websocket](https://github.com/gorilla/websocket)**

只需要在终端输入命令：

```shell
go get github.com/gorilla/websocket
```

通过库中的example，我们使用gin框架来实操一下

##### 升级协议

首先，要使用websocket，我们需要使用http协议来交换信息，让客户端与服务端达成一致，然后升级协议

我们需要用到这个`websocket.Upgrader`

```go
type Upgrader struct {
// HandshakeTimeout 指定握手完成的持续时间。
HandshakeTimeout time.Duration

//ReadBufferSize和WriteBufferSize以字节为单位指定I/O缓冲区大小。
ReadBufferSize, WriteBufferSize int

//WriteBufferPool是用于写入操作的缓冲区池
WriteBufferPool BufferPool

//Subprotocols 包含需要用到的子协议
Subprotocols []string

// Error 指定用于生成HTTP错误响应的函数。如果出现错误为nil，则使用http.Error生成http响应。
Error func(w http.ResponseWriter, r *http.Request, status int, reason error)

//CheckOrigin 防止跨站点请求伪造
CheckOrigin func(r *http.Request) bool

// EnableCompression 指定服务器是否应尝试根据消息压缩（RFC 7692）
EnableCompression bool
}
```

一般无需特别设置

```go
var upgrader = websocket.Upgrader{
CheckOrigin: func(r *http.Request) bool {
return true
},
}
```

然后使用这个方法，即可完成协议升级

```go
func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*Conn, error)
```

代码如下

```go
package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	//......
}
```

##### 收发数据

```go
	//接受数据
go func() {
for ; ; {
_, p, err := conn.ReadMessage()
if err != nil {
log.Println(err)
} else {
log.Println(string(p))
}
}
}()

//发送数据
go func() {
for i := 0; ; i++ {
conn.WriteJSON(gin.H{
"time": time.Now().Format(time.RFC3339),
"No":   i,
})
}
}()
```



##### 完整代码

服务端完整代码如下

```go
package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
		for ; ; {
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

```

运行代码后，我们就有了一个ws接口

```shll
ws://localhost:8080/websocket
```

可以找个在线测试网站进行测试

![uTools_1678119664395](http://typora.fengxiangrui.top/1678119669.png)



#### websocet协议数据包介绍

![img](https://camo.githubusercontent.com/8fb41bdc8f77228c84ca7c7f754f88b3a55ec6d2a41b905652c0de439c5d51a3/68747470733a2f2f696d672d626c6f672e6373646e2e6e65742f32303134303330363233333530313834333f77617465726d61726b2f322f746578742f6148523063446f764c324a736232637559334e6b626935755a585176645441784d4451344e7a55324f413d3d2f666f6e742f3561364c354c32542f666f6e7473697a652f3430302f66696c6c2f49304a42516b46434d413d3d2f646973736f6c76652f37302f677261766974792f536f75746845617374)

##### 第一个字节

![uTools_1678333250420](http://typora.fengxiangrui.top/1678333253.png)

- **FIN:1位，用于描述消息是否结束，**如果为1则该消息为消息尾部,如果为零则还有后续数据包;
- **RSV1,RSV2,RSV3**：**各1位**，用于扩展定义的,如果没有扩展约定的情况则必须为0
- **OPCODE:4位，用于表示消息接收类型，**如果接收到未知的opcode，接收端必须关闭连接。

OPCODE定义的范围：

**0x0表示附加数据帧**
　　　　**0x1表示文本数据帧**
　　　　**0x2表示二进制数据帧**
　　　　0x3-7暂时无定义，为以后的非控制帧保留
　　　　**0x8表示连接关闭**
　　　　**0x9表示ping**
　　　　**0xA表示pong**
　　　　0xB-F暂时无定义，为以后的控制帧保留



##### 第二个字节及后续

- **MASK:1位**，用于标识PayloadData是否经过**掩码处理**，客户端发出的数据帧需要进行掩码处理，所以此位是1。数据需要解码。

- **Payload length** === **x**，如果
    - 如果 **x**值在0-125，则是payload的真实长度。
    - 如果 **x**值是126，则后面2个字节形成的16位无符号整型数的值是payload的真实长度。
    - 如果 **x**值是127，则后面8个字节形成的64位无符号整型数的值是payload的真实长度。

此外，如果payload length占用了多个字节的话，payload length的二进制表达采用网络序（big endian，重要的位在前）。

- **Masking-Key**：掩码key，用于数据解码
- **Playload Data**：负载数据



了解完这些，就可以自己动手搓一个websocket包了

[参考demo]：https://github.com/lluckyboi/websocket



### HTTPS

#### 为什么要HTTPS?

因为我们在把服务暴露到外网时，如果使用的是**HTTP**，那信息的**收发全是明文**，只要有人抓个包，包里的信息就全泄露了。

因此我们在HTTP层之上再加一层TLS层来加密。

![uTools_1678344424790](http://typora.fengxiangrui.top/1678344432.png)

#### HTTPS的握手过程

![uTools_1678344452005](http://typora.fengxiangrui.top/1678344458.png)

首先，建立TCP连接，因为HTTP是基于TCP的

在TCP建立完协议后，就开始加密流程

分为两个阶段：

1. **TLS四次握手**
2. **加密通信**

第一阶段是利用非对称加密的特性交换信息，最后得到一个**会话密钥**

第二阶段则是在**会话密钥**的基础上，进行**对称加密**通信



##### 第一次握手

1. `Client Hello` **客户端**发出：告诉服务端它加密协议版本(TLS1.2)和加密套件信息(RSA)，和一个**客户端随机数**

##### 第二次握手

1. `Server Hello ` **服务端**发出：告诉客户端**服务端随机数**+**服务器证书**+确定的加密协议版本

##### 第三次握手

1. `Client Key Exchange`**客户端**发出：告诉服务端一个它**新生成的随机数**(`pre_master_key`)，并用**服务器证书加密**(公钥)

   2.`Change Cipher Spec`   **客户端**发出：用三个随机数进行计算得到一个"**会话秘钥**"，发给服务端

   3.`Encrypted Handshake Message ` **客户端**发出：客户端把迄今为止的通信数据，做成摘要，并用"**会话秘钥**"加密，发给服务端

##### 第四次握手

1. `Change Cipher Spec` 由**服务端**发出：由于服务器收到了第三个随机数，因此也可以生成**会话密钥**，**后续可以加密传输了**
2. `Encrypted Handshake Message`：跟客户端的操作一样，把迄今为止的通信数据，做成摘要，用"**会话秘钥**"加密一下，发给客户端做校验，到这里，服务端的握手流程也结束了，因此这也叫**Finished报文**。



##### **小结**

- 可以看到，四次握手的目的就是为了生成"**会话密钥**"，后续用**对称加密**的方式进行通信，因为**对称加密更快一点**
- 前两个随机数是明文传输，第三个则是经过服务器公钥加密的
- 用三个随机数是为了增加"会话密钥"的随机性



#### 服务器证书

- 服务器证书是：被**权威数字证书机构（CA）的私钥**加密过的**服务器公钥**。可以通过公开的CA公钥解密。

这样确保了**服务器公钥**的真实性



- 一般CA的公钥**被直接放到了浏览器或是操作系统中**



#### 如何在项目中使用HTTPS

![uTools_1678344544315](http://typora.fengxiangrui.top/1678344547.png)

见:https://juejin.cn/post/7065101056073531428





### RPC

#### 啥是RPC

RPC是指远程过程调用,是Remote Procedure Call三个单词的缩写，功能就是**像本地的函数一样去调远程函数**，远程一般是指通过网络从远程计算机程序上请求服务，也可以在在宿主机下通过网络进行不同架构下的互相请求服务。

在分布式或者微服务架构上非常常用，RPC让不同服务之间的服务调用像本地调用一样简单高效，RPC是一种**网络协议**，是一种规范，每个大厂几乎都有自己研发的RPC协议。

RPC是一种服务器-客户端（Client/Server）模式，经典实现是一个通过**发送请求-接受回应**进行信息交互的系统。

【简单理解】：两台不同计算机（程序），`计算机A`有一个**约定协议**，`计算机B`想调用`计算机A`需要通过**约定协议**来进行通讯调用。

#### RPC与HTTP

##### 先讲讲HTTP

HTTP：**Hypertext Transfer Protocol**即超文本传输协议。

HTTP协议在1990年才开始作为主流协议出现；之所以被我们所熟知，是因为通常HTTP用于web端，也就是web浏览器和web服务器交互。当ajax和json在前端大行其道的时候，json也开始发挥其自身能力，简洁易用的特性让json成为前后端数据传输主流选择。HTTP协议中以Restful规范为代表，其优势很大。它**可读性好**，且**可以得到防火墙的支持、跨语言的支持**。

HTTP的缺点也很快暴露：

1. **有用信息占比少**，HTTP在OSI的第七层，包含了大量的HTTP头等信息
2. **效率低**，因为第七层的缘故，中间有很多层传递
3. HTTP协议**调用远程方法复杂**，需要封装各种参数名和参数值以及加密通讯等

##### 所以RPC好在哪？

1. **都是有用信息**
2. **效率高**
3. **调用简单**
4. **无需关心网络传输或者通讯问题**

##### RPC一般用于什么地方？

在**微服务、分布式**已经成为日常的今天，服务通常都部署在不同的服务器，服务器也在不同地区，这时候就存在跨地域跨服务器调用问题，**RPC即用于这样类似的情况**。

RPC适用于公司内部使用，性能消耗低，传输效率高，服务治理方便，但是不建议传输较大的文本、视频等。



#### GRPC

李文周博客，肯定比我讲的好

https://www.liwenzhou.com/posts/Go/gRPC/



### 作业

1.跑一下weboscket示例代码，尝试配置https(选做) ，学习grpc

2.尝试用websocket实现在线聊天室

