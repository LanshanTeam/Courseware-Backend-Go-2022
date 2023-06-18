## **什么是服务注册发现？**

对于搞微服务的同学来说，服务注册、服务发现的概念应该不会太陌生。

简单来说，当服务A需要依赖服务B时，我们就需要告诉服务A，哪里可以调用到服务B，这就是服务注册发现要解决的问题。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=OGZkZDEwNjU1ODMyNDAzM2E1MTEyY2JhMDZiZDA3Yzdfd0JKWGJxNklUbk4xRmRxQTJDRDludmY1bThyVUpBQmZfVG9rZW46Ym94Y25JRWFnbzc0QkVIMHpVOHZNeFdBR2FiXzE2ODcwNzA5MDg6MTY4NzA3NDUwOF9WNA)

- `Service B` 把自己注册到 `Service Registry` 叫做 **服务注册**
- `Service A` 从 `Service Registry` 发现 `Service B` 的节点信息叫做 **服务发现**

## **服务注册**

服务注册是针对服务端的，服务启动后需要注册，分为几个部分：

- 启动注册
- 定时续期
- 退出撤销

### **启动注册**

当一个服务节点起来之后，需要把自己注册到 `Service Registry` 上，便于其它节点来发现自己。注册需要在服务启动完成并可以接受请求时才会去注册自己，并且会设置有效期，防止进程异常退出后依然被访问。

### **定时续期**

定时续期相当于 `keep alive`，定期告诉 `Service Registry` 自己还在，能够继续服务。

### **退出****撤销**

当进程退出时，我们应该主动去撤销注册信息，便于调用方及时将请求分发到别的节点。同时，go-zero 通过自适应的负载均衡来保证即使节点退出没有主动注销，也能及时摘除该节点。

## **服务发现**

服务发现是针对调用端的，一般分为两类问题：

- 存量获取
- 增量侦听

还有一个常见的工程问题是

- 应对服务发现故障

当服务发现服务（比如 etcd、consul、nacos等）出现问题的时候，我们不要去修改已经获取到的 endpoints 列表，从而可以更好的确保 etcd 等宕机后所依赖的服务依然可以正常交互。

### **存量获取**

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=NDNkZTUyNmM4MWFkYWYxNDBmMDk4ZThkOTE2YTliYWZfdFlYVHZIZmwxeEJvRG84WENmMlZJTWtHYkJxQTlUb1ZfVG9rZW46Ym94Y25teDB1bnM3NE04VnBEUUx4a2F5OFFkXzE2ODcwNzA5MDg6MTY4NzA3NDUwOF9WNA)

当 `Service A` 启动时，需要从 `Service Registry` 获取 `Service B` 的已有节点列表：`Service B1`, `Service B2`, `Service B3`，然后根据自己的负载均衡算法来选择合适的节点发送请求。

### **增量侦听**

上图已经有了 `Service B1`, `Service B2`, `Service B3`，如果此时又启动了 `Service B4`，那么我们就需要通知 `Service A` 有个新增的节点。如图：

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=ZWQ4MGQ3YzI5MDI5M2M4YjZlYjQwZDUwNzM0YjBiOTZfb2J1OEI1NGliMGpOT0F0YXdvSGNJQ25jNXY2TzF1U3VfVG9rZW46Ym94Y25HWFlwRVdtbnE1VHRtajdSeFFLekVKXzE2ODcwNzA5MDg6MTY4NzA3NDUwOF9WNA)

### **应对服务发现故障**

对于服务调用方来说，我们都会在内存里缓存一个可用节点列表。不管是使用 `etcd`，`consul` 或者 `nacos` 等，我们都可能面临服务发现集群故障，以 `etcd` 为例，当遇到 `etcd` 故障时，我们就需要冻结 `Service B` 的节点信息而不去变更，此时一定不能去清空节点信息，一旦清空就无法获取了，而此时 `Service B` 的节点很可能都是正常的，并且 `go-zero` 会自动隔离和恢复故障节点。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=MzJlOTI5NGU5MzZkYzAxNzE3MDg3M2YxNzg1NWNiMjNfRXV2eHRvNHFsYjJtckx4eDEzUWx0OWxEUjBYUWF4S2lfVG9rZW46Ym94Y25CM2kxMzVMS09jbEZFUXpmTmdjc1RjXzE2ODcwNzA5MDg6MTY4NzA3NDUwOF9WNA)

服务注册、服务发现的基本原理大致如此，当然实现起来还是比较复杂的，接下来我们一起看看 `go-zero` 里支持哪些服务发现的方式。

## 服务发现组件的部署

### Consul

```Shell
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp  consul consul agent -dev -client=0.0.0.0
```

部署成功后打开 `127.0.0.1:8500` 即可打开管理界面。

### Etcd

在任意文件夹下新建 Makefile 文件，并填写一下内容。

```Makefile
NODE1=172.18.31.33
REGISTRY=bitnami/etcd
ETCD_VERSION=latest
# run cluster
prepare-cluster:
        cd example
        docker-compose up -d

# run single node in docker 
prepare:
        docker pull ${REGISTRY}:${ETCD_VERSION}
        docker run -d --name Etcd-server \
    --network app-tier \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
```

在文件夹的路径下运行 `make prepare` 初始化 `Docker` ，再使用 `make prepare-cluster` 运行 Etcd 集群。

## **go-zero 之内置服务发现（这部分是抄的，建议大家还是网上找找比较详细的）**

`go-zero` 默认支持三种服务发现方式：

- 直连
- 基于 etcd 的服务发现
- 基于 kubernetes endpoints 的服务发现

### **直连**

直连是最简单的方式，当我们的服务足够简单时，比如单机即可承载我们的业务，我们可以直接只用这种方式。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=ZTM1ODhhM2Q4MTRiOTRlYjMzMGU3ZDVkZjdjMWE5MTJfSHM4VFRuRUdWcGpYekVSQjhXaWpjN0dGa0pjZmpGeWhfVG9rZW46Ym94Y25xRTdoV3VIcExXbXZYcFcxanpXSmpoXzE2ODcwNzA5MDg6MTY4NzA3NDUwOF9WNA)

这个方法的缺点是不能动态增加节点，每次新增节点都需要修改调用方配置并重启。

### **基于 etcd 的服务发现**

当我们的服务有一定规模之后，因为一个服务可能会被很多个服务依赖，我们就需要能够动态增减节点，而无需修改很多的调用方配置并重启。

常见的服务发现方案有 `etcd`, `consul`, `nacos` 等。

![img](https://lanshanteam.feishu.cn/space/api/box/stream/download/asynccode/?code=YjM5N2E4ZDRiMTY0Mjk3ZTNjMGQ2NjBlMmVkMjliYmRfRXpFVUlPWTJFRjVQWW5RbjhiS1ZYUlU0ZnhZMFQ1QzRfVG9rZW46Ym94Y25nYWlYbUJ4a0diUTFJNlIyU0RoTFJkXzE2ODcwNzA5MDg6MTY4NzA3NDUwOF9WNA)

go-zero内置集成了基于 `etcd` 的服务发现方案，具体使用方法如下：

```Makefile
Rpc:
  Etcd:
     Hosts:
     - 192.168.0.111:2379
     - 192.168.0.112:2379
     - 192.168.0.113:2379
     Key: user.rpc
1234567
```

- `Hosts` 是 `etcd` 集群地址
- `Key` 是服务注册上去的 `key`

### **基于 Kubernetes Endpoints 的服务发现**

这个有点高级了，你们现在不会。

## CloudWeGo 的服务发现组件

***前置内容***

> 需要通过 WSL2 或者 Linux 或者 MacOS 进行开发

https://juejin.cn/post/7216255521372110909

### Kitex 中使用

https://github.com/kitex-contrib/registry-etcd

https://github.com/kitex-contrib/registry-consul

### Hertz 中使用

https://github.com/hertz-contrib/registry