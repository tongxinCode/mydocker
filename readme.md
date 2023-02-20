# mydocker

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
