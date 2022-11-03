### 基本使用

- `RUN`命令有两种语法可以使用：
  - 1、`shell`形式: `RUN <command>`
    - 在`linux`中使用`/bin/sh -c`的方式执行命令，可以通过命令指定其它解释器执行命令，譬如`RUN /bin/bash -c echo hello`
    - 在`windows`中使用`cmd /S /C`的方式执行命令
  - 2、`exec`形式: `RUN ["executeable", "param1", "param2"]`
    - 必须使用双引号括住参数
    - 这种形式不会调用shell解释器执行命令
- mount
- network
- secure

#### 指定执行器

```dockerfile
FROM busybox

# 这种方式必须使用双引号
RUN ["/bin/sh", "-c", "echo hello"]

# 这种方式的参数必须使用引号, 单引号和双引号都可以
RUN /bin/sh -c "echo hello"
RUN /bin/sh -c 'echo hello'
```

```dockerfile
FROM busybox

RUN echo $HOME
RUN /bin/sh -c 'echo $HOME'
RUN sh -c  'echo $HOME'

# 由于 exec 格式的RUN命令，并不会调用shell解释执行，因此不会解析HOME环境变量，而是当作一个普通的字符串
RUN ["echo", "$HOME"]

# 可以手动调用shell解释器执行命令
RUN ["/bin/sh","-c","echo $HOME"]
```

#### --mount=type=bind

#### --mount=type=cache

#### --mount=type=tmpfs

#### --mount=type=secret

#### --mount=type=ssh
