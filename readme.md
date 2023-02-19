# mydocker

## 3.2

cpu.shares意味着默认cpu的分片权重，当仅有一个cgroup任务时无论设置为多少，该任务会占用相应物理cpu的100%

memory.limit_in_bytes意味着最大使用内容，当stress程序试图使用更大内存时会异常退出，下方是错误输出

```bash
stress: FAIL: [1] (416) <-- worker 5 got signal 9
```

