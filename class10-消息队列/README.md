# ✨消息队列

## 🤔什么是消息队列

来我们先来看一下消息队列的定义。

**消息队列（****Message Q****ueue****） 简称****MQ**是一种先进先出的队列型数据结构，实际上是系统内核中的一个内部链表。拥有权限的进程可以向消息队列中写入或读取消息，消息被顺序插入队列中，其中发送进程将消息添加到队列末尾，接受进程从队列头读取消息。

看完定义，我们现在大概知道了什么是MQ。我们可以简单的把MQ比喻成一个蓝山工作室的公共盒子，不同成员可以把信息存放到盒子中，可以让工作室的成员消费。

## 为什么要使用MQ

消息队列的优点

1. 屏蔽异构平台的细节，发送方和接收方不需要了解对方的系统，只需要认识信息就够了;
2. 异步：消息可以堆积在队列中，发送方和接收方不需要同时在线，也不需要同时扩容。
3. 削峰：在系统请求量/并发量高的情况下，如果我们直接把这些请求打到列如MySQL上，很有可能直接就打死。
4. 解耦：由于发送方和接收方没有直接依赖，可以直接灵活增加or修改模块，提高系统的可扩展性，这点在分布式系统中尤为重要。

总结一下

1. 在特殊场景下，比如高并发、高可用、分布式等情况下，使用消息队列可以提高系统性能和稳定性。
2. 在事件驱动架构中，使用消息队列可以实现发布-订阅模式，让不同的模块之间可以相互通信和协作。

这里我们拿我们的产品we重邮来说。

例如： 阅读数量功能

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=MzY2YTRhOWNlY2ZjZmZhMmJjYTA5OTZhYzgwYTRkMGRfVnhJU1RvRUpFVWwxQ0FaZzVqeXduRzhSekwyWGFLSDZfVG9rZW46Ym94Y25EeW81SnJMaDNYSDFYSXFTbVlzUDNNXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

欸嘿：🙌 结合到上面刚刚所学到内容。我们先来分析一下场景。如果一个一般来说一篇咨询可能同时有很多浏览，一股脑的放到redis上面，这在请求量高的情况下是肯定不行。那怎么办捏，我们就把消息放到MQ中，等服务器空闲再拉取NSQ。

## 📡消息协议

消息队列要在网络中进行通信，那就需要采用一种通信协议。

**AMQP**、**MQTT**和**STOMP**是三种最常见、最流行的基于TCP/IP的消息传递协议。

### AMQP

高级消息队列协议，即Advanced Message Queuing Protocol（[AMQP](https://www.amqp.org/about/what)）,一个提供统一消息服务的应用层标准 高级消息队列协议（二进制应用层协议），是应用层协议的一个开放标准,为面向消息的中间件设计，兼容 JMS。基于此协议的客户端与消息中间件可传递消息，并不受客户端/中间件同产品，不同的开发语言等条件的限制。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=MjUxZWE0NjhjZWJlMTcyYTU4NmU1MDU0NjBmYzJlYzJfRWx6WFp4bW9EdEJqWThLbjRoaTVSdjM2OXdqc0NsUXlfVG9rZW46Ym94Y25IQVVoTUo1Y2FIYU1DN25sb1FpNkpGXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

### MQTT

消息队列遥测传输协议，即Message Queuing Telemetry Transport（MQTT），是一种基于`发布/订阅`（`publish/subscribe`）模式的“轻量级”通讯协议，由IBM在1999年发布。议由于其轻量、简单、开放和易于实现的，这些特点使它适用范围非常广泛。包括受限的环境中，如：机器与机器（M2M）通信和物联网（IoT）。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=NGNhODMzZmM2NGU3NjY0NzNjZTM1NzQ1OGVkNWJkMjZfVVYzcDBLV2VvdEdyRWk4bjgwaXB4WGxRUmJwYmRvUEhfVG9rZW46Ym94Y25DNEVRcDZLdDdKZjlwOEpETzlIQ0doXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)



### STOMP

简单（或流式）文本消息协议，Simple (or Streaming) Text Oriented Messaging Protoco（[STOMP](https://juejin.cn/post/7028140479514411038)） 协议是一种基于帧的协议，它可以让不同的客户端通过一个中间服务器进行异步消息传递。它是一个简单的可互操作的协议，它的设计灵感来源于 HTTP 的简单性。STOMP 协议支持发布-订阅和点对点两种消息模式。STOMP 协议可以用于多种语言和平台，例如 Java、Python、Ruby、JavaScript 等。

## 消息队列模型

MQ最简单的模型就三个角色。

消息队列：存储消息。生产者：发送消息到消息队列。消费者：从消息队列获取消息并处理消息。

### 队列模型

它允许多个生产者往同一个队列发送消息。但**多个消费者之间是竞争的关系**，也就是说**一条消息只能被其中一个消费者接收到，读完即被删除。**

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=NjMwZjEzMDc2ZWVlYWVmZjI0MWUzYjBlN2Q5MDFjNWZfMFh2bHBBWHNnODZrVzBBMHo4UjF4cUtsYVFIVXEyM0RfVG9rZW46Ym94Y25uRVVOdVRhczVCVGhTRVRCdHdCZTJlXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

### 发布订阅模型

那如果我们要将消息发送给多个消费者，或者是被多个消费者消费，那简单的队列模型显然无法做到。

如果我们硬要用队列模型也不是不可以。我们可以为每一个消费者单独创建一个队列，让生产者发送多份信息给消费者的消费队列。但是这样也会存在很多问题。生产者必须事先知道那些消费者需要这个信息。而且在创建了许多相同的信息，造成了资源的浪费。

为了解决这个问题，就演化出了另外一种消息模型：发布-订阅模型。

在发布-订阅模型中，消息的发送方称为发布者(`Publisher)`，消息的接收方称为订阅者`(Subscriber`)， 服务端存放消息的容器称为主题(`Topic`)。发布者将消息发送到主题中，订阅者在接收消息之前需要先「订阅主题」。「订阅」在这里既是一个动作，同时还可以认为是主题在消费时的一个逻辑副本，每份订阅中，订阅者都可以接收到主题的所有消息。

实际上，在这种发布-订阅模型中，如果只有一个订阅者，那它和队列模型就基本是一样的了。也就是说，发布-订阅模型在功能层面上是可以兼容队列模型的。（这两种消息模型其实并没有本质上的区别，都可以通过一些扩展或者变化来互相替代

发布-订阅模型和队列模型的唯一不同点在于：一份消息数据是否可以**被多次消费**。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=ZmE5YzY3YWZkY2M5ZTQ3MmE4OTJhZDI3MzkyN2FlZGNfNDNOZTgzbjhZcUJOMXhvOTFkYU5KdG9GR2dscTFSRkdfVG9rZW46Ym94Y245ZGFDNUpZbUdlem1vUUZRdXJ5WXdlXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

## 常见的MQ

### RabbitMQ

RabbitMQ是由erlang语言开发，基于AMQP协议，可复用的企业消息系统，是当前最主流的消息中间件之。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=NzM3YzUwYzdmNGFmMGUwZjAwY2VmNjFjNDVhMDU4NDNfSHZnSUVud0pxWm5kUmRLTTBEM2ZVWXRnZVE0azM5SmZfVG9rZW46Ym94Y25EUks3bHk5eldjcTNxU0w3UXdLNGtlXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

组成部分说明：

`Broker`：消息队列服务进程，此进程包括两个部分：Exchange和Queue

`Exchange`：消息队列交换机，按一定的规则将消息路由转发到某个队列，对消息进行过虑。

Queue：消息队列，存储消息的队列，消息到达队列并转发给指定的

Producer：消息生产者，即生产方客户端，生产方客户端将消息发送

Consumer：消息消费者，即消费方客户端，接收MQ转发的消息。

生产者发送消息流程：

1. 生产者和Broker建立TCP连接。
2. 生产者和Broker建立通道。
3. 生产者通过通道消息发送给Broker，由Exchange将消息进行转发。
4. Exchange将消息转发到指定的Queue（队列）

消费者接收消息流程：

1. 消费者和Broker建立TCP连接
2. 消费者和Broker建立通道
3. 消费者监听指定的Queue（队列）
4. 当有消息到达Queue时Broker默认将消息推送给消费者。
5. 消费者接收到消息。
6. ack回复

#### 基本工作模式

我们可以从他的官网上看到这些[RabbitMQ官网](https://www.rabbitmq.com/)

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=NTFlZGEwZDMyYmU2OWU2MTkyZWJiMTJjNjRiMWFiZGZfV3BkeU5YdUpzVThWamxzdnBIaDF2ZEhOV0d2Q1JhMlZfVG9rZW46Ym94Y251ZG83THlzQ3BCc1k4dXRRRmRFNjBiXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=YzVjOTk2ZTk3NjhlOTQxNmY5MWViY2ZjNjI0YzVhMTNfa3pTZEhWYjE2SWd4WmpXSzl1OHQyenh4bHpsNUVldGpfVG9rZW46Ym94Y25sb0l3cVhFaWNScGU0MHJZY3N3THVoXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

总共七种工作模式。代码看官网，有很多语言的示例。我们重点讲解一下，消费模型的区别。

##### 基本消费模型

#### 简单模式

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=OGZiNTAzMTZmNmVjOTQ1NWIxNjE1NDUyOWY0MzBjODdfd3puWGtVbTdqVGVabjhxck1vSXR0V1BUSGNpTE9BdElfVG9rZW46Ym94Y25lQXlkYnFtWjFURDJrVlZ4SzFwcWFmXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

在上图的模型中，有以下概念：

- P：生产者，也就是要发送消息的程序
- C：消费者：消息的接受者，会一直等待消息到来。
- queue：消息队列，图中红色部分。可以缓存消息；生产者向其中投递消息，消费者从其中取出消息。

##### 工作模式

工作队列或者竞争消费者模式

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=MjJlMTQ4NjVlM2Q5NmJmZWQ4YzcxMzExMTE3MjlmNDBfekZjeVVjMldydmpseGw4OXVseEpUM1NzcHNkeDB0cm9fVG9rZW46Ym94Y25raHRBb1lLbWVPTnJ5VHprUXFNVm1lXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

工作队列和普通的队列相比，多了一个消费端，两个消费端共同消费同一个队列中的消息，**但是一个消息只能被一个消费者获取**。

这个消息模型在Web应用程序中特别有用，可以处理短的HTTP请求窗口中无法处理复杂的任务。

我们可以简单的模拟一下这个过程：

P：生产者：任务的发布者

C1：消费者1：领取任务并且完成任务，假设完成速度较慢（模拟耗时）

C2：消费者2：领取任务并且完成任务，假设完成速度较快

##### **Publish/subscribe发布订阅 模式**

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=OGQwODM3ZWYwMjBjNmIwMWFmZTAxZmIzYzJhNjViOWRfbjVPUDkzU3hiWTRGSTk1b29lNzRNaENqOGV2bmhISGpfVG9rZW46Ym94Y25KRXByVE5WWURVd3ZFbzZqR21PMVRnXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

从示意图可以看到 生产者和前面两种模式不同：

- 1） 声明Exchange，不再声明Queue
- 2） 发送消息到Exchange，不再发送到Queue

思考：

1. publish/subscribe与work queues有什么区别。

区别：

1. work queues不用定义交换机，而publish/subscribe需要定义交换机。
2. publish/subscribe的生产方是面向交换机发送消息，work queues的生产方是面向队列发送消息(底层使用默认交换机)。
3. publish/subscribe需要设置队列和交换机的绑定，work queues不需要设置，实际上work queues会将队列绑定到默认的交换机 。

相同点：

1. 两者实现的发布/订阅的效果是一样的，多个消费端监听同一个队列不会重复消费消息。
2. 实际工作用 publish/subscribe还是work queues。

建议使用 publish/subscribe，发布订阅模式比工作队列模式更强大（也可以做到同一队列竞争），并且发布订阅模式可以指定自己专用的交换机。

##### **Routing 路由模式**

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=MTE1MGUyOGVkYzA2YjA3MTU4Mjc0NWM1ZWRjOTE4NzNfTmhYTkJEM1VKbU1CeFZtNFFIU1pwS215cFVLZGZtQUVfVG9rZW46Ym94Y25CYm1LTFVUMUJBYzBtcnlIODJNRnpiXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

P：生产者，向Exchange发送消息，发送消息时，会指定一个routing key。

X：Exchange（交换机），接收生产者的消息，然后把消息递交给 与routing key完全匹配的队列

C1：消费者，其所在队列指定了需要routing key 为 error 的消息

C2：消费者，其所在队列指定了需要routing key 为 info、error、warning 的消息

###### 

##### **Topics 主题模式**

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=OTBlOTY5YTNjMzQ3NDNhODZiODhhMWE5NTQ3NjdkYTVfaEQyenllNnNkcEpuVU4xU1JadWZGMjNZbUN2eGZVUkRfVG9rZW46Ym94Y25pQm0zNjZTa1BLdnFFMWhPc1ZHMUdlXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

生产者生产消息投递到topic交换机中，上面是完全匹配路由键，而主题模式是模糊匹配，只要有合适规则的路由就会投递给消费者

#### 优缺点

优点：

- 高性能，能够处理大量的并发消息
- 易于配置和使用，提供了丰富的文档和管理界面
- 灵活的路由，能够根据不同的规则将消息分发到不同的队列
- 可靠性，支持持久化，确认机制，高可用等特性
- 多语言支持，提供了多种语言的客户端库

缺点：

- Erlang 依赖，需要安装 Erlang 环境才能运行 RabbitMQ
- 内存占用较高，当消息堆积时可能导致内存溢出或性能下降
- 集群管理复杂，需要手动配置集群节点和镜像队列
- 不支持延迟消息，需要借助死信队列或插件实现

### RocketMQ

RocketMQ是一个纯Java、分布式、队列模型的开源消息中间件，前身是MetaQ，是阿里参考Kafka特点研发的一个队列模型的消息中间件，后开源给apache基金会成为了apache的顶级开源项目，具有高性能、高可靠、高实时、分布式特点。[RocketMQ · 官方网站 | RocketMQ](https://rocketmq.apache.org/)

他主要有四大核心组成部分：**NameServer**、**Broker**、**Producer**以及**Consumer**四部分。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=NjQyZDhlOWZmNjAzNWE1ZmFiZmQ1NDRkOGVhNmZmNDNfQklCWVl2SnpYYlRtTXk2dUhnMzVyRkFSalEwZ01sYzBfVG9rZW46Ym94Y25OZnJDQVVhTmYybU1QT3B2MG9NR1FnXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

**Tip**：我们可以看到**RocketMQ**啥都是**集群**部署的，这是他**吞吐量****大**，**高可用**的原因之一，集群的模式也很花哨，可以支持多master 模式、多master多slave异步复制模式、多 master多slave同步双写模式。

由于其基于Kafaka 所以其模式和Kafaka差不多。

#### 优缺点

优点：

- 支持事务消息，能够保证消息的最终一致性
- 支持顺序消息，能够按照指定的顺序消费消息
- 支持延迟消息，能够在指定的时间后发送或消费消息
- 支持批量发送和消费消息，提高吞吐量
- 支持多种部署模式，包括单机模式，集群模式和云原生模式

缺点：

- 学习成本较高，需要掌握多种概念和组件
- 配置参数较多，需要根据业务场景进行调优
- 依赖外部存储系统（如MySQL）来存储元数据信息

### Kafaka

Kafka 是一个分布式的消息发布订阅系统，最初由 LinkedIn 公司开发，后来成为 Apache 项目的一部分。Kafka 可以处理高性能，高可靠性，高容量和高扩展性的数据流。支持多种消息模式，如发布/订阅，请求/响应和流式处理。提供了丰富的管理控制台来配置，监控和度量。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=Y2ZkNTU2ZjkwZWJmZGM3YWFlYzg3OTYxZjY1MDQ4MjFfSWh5a0lUaklZQjFpZHJrTEdrZ0J3SXpmOXE5ZDBSYk9fVG9rZW46Ym94Y25CWWZKc0lNb0FJS0w0ZEw2YjBzckxiXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

#### 优缺点：

- 低延迟：Kafka 可以在 10 毫秒内处理消息，因为它解耦了消息，让消费者可以随时消费消息
- 高吞吐量：由于低延迟，Kafka 能够处理高速度和高容量的消息，每秒可以支持数千条消息
- 持久化：Kafka 可以将数据流持久化到磁盘中，并且可以根据需要保留数据一段时间
- 容错性：Kafka 可以通过分区和复制机制来实现容错性，并且可以在节点故障时自动恢复
- 零停机：Kafka 可以在不影响服务的情况下进行升级或扩展

缺点：

- 学习成本较高：Kafka 需要掌握多种概念和组件，如生产者、消费者、主题、分区、偏移量、代理等
- 配置参数较多：Kafka 需要根据业务场景进行调优，并且需要考虑数据一致性、可用性和延迟之间的平衡
- 依赖外部系统（如 ZooKeeper）来管理集群状态和元数据信息

### NSQ

NSQ 是一种基于 Go 语言的分布式实时消息平台，它具有分布式、去中心化的拓扑结构，支持无限水平扩展。无单点故障、故障容错、高可用性以及能够保证消息的可靠传递的特征

#### NSQ 组件

1. nsqd：接受、排队、传递消息的守护进程，消息队列中的核心。
2. nsqlookupd：管理拓扑信息，其实就是围绕 nsqd 的发现服务，因为其存储了 nsqd 节点的注册信息，所以通过它就可以查询到指定 topic 主题的 nsqd 节点。
3. nsqadmin：一套封装好的 WEB UI ，可以看到各种统计数据并进行管理操作。
4. utilities：封装好的一些简单的工具。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=Y2ZmZDk1YjQ0ZWZkZGFiODM1YjBiYWMxYWZmYmU4YjJfWkF6am5xUzJJQ08xV2FaMXhnWXpIRkE2cHBhV1NrRTFfVG9rZW46Ym94Y25DQUNxbWVIaTRPdTNFT3k5cnl5WlliXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

从上图我们可以看到

1. 生产者 producer 将消息投递到指定的 nsqd 中指定的 topic 主题。
2. nsqd 可以有多个 topic 主题，一旦其接受到消息，将会把消息广播到所有与这个 topic 相连的 channel 队列中。
3. channel 队列接收到消息则会以负载均衡的方式随机的将消息传递到与其连接的所有 consumer 消费者中的某一个。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=M2IzMWExMjZhZTFkMTZjY2U3YzMyMTg2ZjU3ODNmOGJfZTVnT3dlTmJqVjZRUkpXM2ZDM1A3SjJoTEtXN1Nsb21fVG9rZW46Ym94Y25JcTU1YUt2MnByNFpneHYwS1FjT1RnXzE2ODcwNzA4NjU6MTY4NzA3NDQ2NV9WNA)

通过 nsqadmin 可以看到整个集群的统计信息并进行管理，多个 nsqd 节点组成集群并将其基本信息注册到 nsqlookupd 中，通过 nsqlookupd 可以寻址到具体的 nsqd 节点，而不论是消息的生产者还是消费者，其本质上都是与 nsqd 进行通信。

#### 优缺点

优点：

- 部署简单，配置极简
- 支持无限水平扩展，无单点故障
- 支持多种消息协议，如 HTTP、TCP、WebSocket 等
- 支持消息重试和延迟投递
- 支持数据备份和恢复

缺点：

- 消费者端不好做流控，很难做批量推送
- 消费者可能经常有空 pull，即 pull 不到消息，造成浪费
- 不支持事务和顺序消费

## 作业

### Lv1

部署上述任意MQ，实现信息传输。

### Lv2

部署上述任意一个消息队列，结合 Redis 。实现缓冲，解耦化。

### Lv3

用Go实现简单的队列模型 。

**参考链接以及课外阅读**

- [什么是消息队列](https://zhuanlan.zhihu.com/p/52773169)
- [AMQP协议](https://zhuanlan.zhihu.com/p/147675691)
- [STOMP协议](https://juejin.cn/post/7028140479514411038)
- [MQTT协议](https://mcxiaoke.gitbooks.io/mqtt-cn/content/mqtt/01-Introduction.html)
- [RabbitMQ快速入门](https://blog.csdn.net/kavito/article/details/91403659)
- [浅入浅出RocketMQ ](https://juejin.cn/post/6844904008629354504)
- [Kafka 核心机制与架构](https://juejin.cn/post/7176576097205616700)
