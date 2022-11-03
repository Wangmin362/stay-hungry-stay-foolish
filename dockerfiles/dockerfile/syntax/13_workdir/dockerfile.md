
### 作用

- 用于指定`WORKDIR`之后所有命令的工作目录
- `WORKDIR`可以使用绝对路径，也可以使用相对路径，如果是相对路径，那么会相对于像一个`WORKDIR`的相对路径

```dockerfile
FROM busybox

WORKDIR /opt
RUN pwd

# 实际的工作目录为：/opt/abc
WORKDIR abc
RUN pwd

# 实际的工作目录为：/opt/abc/def
WORKDIR def
RUN pwd

# 实际的工作目录为：/root/k8s
WORKDIR /root/k8s
RUN pwd
```
