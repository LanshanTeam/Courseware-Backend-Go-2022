# Linux基础与部署

### Linux基础

#### 啥是linux?

> Linux是一套免费使用和自由传播的类Unix操作系统，是一个基于POSIX和UNIX的多用户、多任务、支持多线程和多CPU的操作系统。它能运行主要的UNIX工具软件、应用程序和网络协议。它支持32位和64位硬件。Linux继承了Unix以网络为核心的设计思想，是一个性能稳定的多用户网络操作系统。 Linux操作系统诞生于1991 年10 月5 日（这是第一次正式向外公布时间）。Linux存在着许多不同的Linux版本，但它们都使用了Linux内核。Linux可安装在各种计算机硬件设备中，比如手机、平板电脑、路由器、视频游戏控制台、台式计算机、大型机和超级计算机。 严格来讲，Linux这个词本身只表示**Linux内核**，但实际上人们已经习惯了用Linux来形容整个基于Linux内核，并且使用GNU 工程各种工具和数据库的操作系统。

简单来说，linux就是一种开源的、强大的、被使用最多的操作系统

**Web服务器通常也是linux**



#### linux常用命令

###### 查看当前目录下所有文件

```shell
ls
```

######  切换目录

```shell
#绝对路径
cd /home/
#相对路径
cd GoProJ
```

###### 查看端口号占用情况

```shell
lsof -i:8080
```

###### 查看所有用户所有进程

```shell
ps -a
```

###### 强制杀死指定进程

```shell
kill -9 12345
```

###### 查看文本文件内容

```shell
cat 1.txt
cat main.go
```

###### 修改文件权限

```shell
chmod 733 main
chmod +x main
```

######  创建文件夹

```shell
mkdir goProJ
```

###### 创建文件

```shell
#touch命令
touch main.go
#vim工具
vi main.go
```

[vim入门到精通]:https://zhuanlan.zhihu.com/p/68111471	.

###### 删除文件

```shell
#有文件夹会一起删除
rm -rf main.go
```

###### 常用快捷键

```shell
 #停止进程
 ctrl + c
 #清屏
 ctrl + l
```

[进阶]:https://www.lanqiao.cn/courses/1	.



### 部署

写好一个web程序，我们如何把它部署到服务器呢？

大致分为两种

1.编译为二进制文件，上传到服务器，服务器直接运行二进制文件

2.将源代码上传到服务器，服务器进行编译(前提是服务器配置好go环境)



#### 上传二进制文件到服务器

首先，在goland终端更改环境变量

```shell
go env -w GOOS=linux
```

然后，来到项目根目录，输入命令编译源代码

```shell
go build main.go
```

等待几秒，你会发现项目中多了一个main文件(注意不是main.exe)

这就是编译好的二进制文件，可以在任何linux上运行

我们可以通过以下命令，将此文件上传到服务器

```powershell
#将当前目录下main文件拷贝到127.0.0.1的/home/路径下
scp main root@127.0.0.1:/home/
```

如果想要免密登录
配置SSH即可

1.在主机A创建密钥对

```shell
ssh-keygen #创建证书
```

然后均回车（选择默认）



2、将文件上传至免登录主机B的authorized_keys

```shell
/root/.ssh/authorized_keys
```



####  上传源代码到服务器

通常会先将源码上传到github再拉下来，或者使用goland一键上传

[服务器go安装]:https://blog.csdn.net/qq_43098070/article/details/126075629	.

然后进入项目目录直接编译即可

```shell
go build main.go
```



#### 运行程序

现在服务器上有编译好的二进制文件了，那我们如何运行它？

只需要

```shell
./main
```

但是我们会发现无法运行，因为文件没有权限

这时候只需要

```shell
chmod +x main
chmod 733 main
```

给它高权限再次运行即可



这时候，使用云服务器的同学们需要去控制台找到安全组的配置规则，在入方向添加相应端口

![uTools_1669869715218](http://typora.fengxiangrui.top/1669869734.png)

![uTools_1669869756091](http://typora.fengxiangrui.top/1669869785.png)



然后，就可以打开浏览器，输入ip:port访问你的服务了



##### 后台运行

如果按以上方式运行，当我们关闭窗口时会发现程序被终止了

那有什么办法可以让程序在后台运行呢？

我们可以使用tmux或者nohup+&命令

###### nohup+&

```shell
nohup ./main &
```

###### tmux

[tmux]:https://www.ruanyifeng.com/blog/2019/10/tmux.html  .


### 作业

发送到fengxiangrui@lanshan.email

#### lv1

部署课件中样例到服务器，浏览器访问页面并截图

#### lv2

熟悉linux指令