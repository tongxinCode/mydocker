# mydocker

项目跟从《自己动手写Docker》/陈显鹭，王炳燊，秦妤嘉著.—北京：电子工业出版社，2017.7

基本实现了容器运行时构建、容器引擎和部分容器管理功能

事实上，引自书中：

    在过去的五年里，Linux已经逐步增加了Cgroups、Namespace、Seccomp、capability、Apparmor等一些功能，这些特性使得容器技术得以在Linux上快速发展。Docker重度使用这些特性，而且目前风靡大江南北。实际上，容器技术是一系列晦涩难懂甚至有些神秘的系统特性的集合，因此Docker公司将这些底层的技术合并在一起，开源出了一个项目runC。

    把containerd从Docker 里彻底剥离出来，作为一个独立的开源项目独立发展，目标是提供一个更加开放、稳定的容器运行基础设施。和原先包含在Docker 里的containerd相比，独立的containerd将具有更多的功能，可以涵盖整个容器运行时管理的所有需求。

Docker容器技术的架构暂时是这样的：

    hardware
    |
    operation system(linux)
    |
    runC--> runtime tool(for creating container process)
    |
    containerd-shim(for configuring the io of container processes)
    |
    containerd--> container engine(for managing the containers, images)
    |
    dockerd-->user-interface server
    |
    docker,docker-compose--> user-interface tool
    |
    swarm--> services manager

凡是涉及容器镜像和运行时创建的，目前既定的标准是OCI（Open Container Initiative）

针对OCI，上层管理工具Kubernetes等使用CRI（Container Runtime Interface）、CNI（Container Network Interface）、CSI（Container Storage Interface），用于对接不同的容器引擎

因此，流行的容器架构往往是这样的：

    any hardware
    |
    any operation system
    |
    any runtime tool
    |
    any container engine
    |
    any services manager
    |
    user-interface tool

## 3.2

cpu.shares意味着默认cpu的分片权重，当仅有一个cgroup任务时无论设置为多少，该任务会占用相应物理cpu的100%

```bash
build/mydocker run -ti -cpushare 256 -m 100m stress --vm-bytes 20m --vm-keep --vm 1
```

memory.limit_in_bytes意味着最大使用内容，当stress程序试图使用更大内存时会异常退出，下方是错误输出

```bash
stress: FAIL: [1] (416) <-- worker 5 got signal 9
```

## 4.1

```bash
build/mydocker run -ti sh
```

## 4.3

```bash
build/mydocker run -ti -v /root/volume:/containerVolume sh
```

## 4.4

```bash
build/mydocker run -ti sh
```

另外打开一个terminal窗口,执行mydocker commit命令

## 5.1

```bash
build/mydocker run -d top
```

## 5.2

```bash
build/mydocker run -d top

build/mydocker run -d --name bird top

build/mydocker ps
```

## 5.3

```bash
build/mydocker logs bird
```

## 5.4

```bash
build/mydocker run -d --name bird top

build/mydocker ps

build/mydocker exec bird sh
```

## 5.5

```bash
build/mydocker run -d --name bird top

build/mydocker ps

build/mydocker stop bird 
```

## 5.6

```bash
build/mydocker rm bird
```

## 5.7 && 5.8

TODO or pass

不同容器的隔离文件系统

容器指定环境变量

## 6.5

```bash
build/mydocker network create --driver bridge --subnet 192.168.10.1/24 testbridge

build/mydocker run -ti -net testbridge sh

build/mydocker run -ti -net testbridge sh
```

```bash
build/mydocker run -ti -p 80:80 -net testbridge sh

nc -lp 80
```

```bash
# host os
sysctl -w net.ipv4.conf.all.forwarding=1
```

```bash
# echo "nameserver 10.143.22.118" > /etc/resolv.conf
# 访问外网服务器地址
ping 110.242.68.66
```
