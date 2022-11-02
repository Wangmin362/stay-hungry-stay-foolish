### ENV指令

- 可以使用`\`对`$`符号进行转义，从而使得特殊字符串不被`docker build`替换为`ENV`环境变量
- 环境变量可以使用`$ENV`，或者`${ENV}`的方式引用，后者主要用于解决环境变量有空格的情况
- `ENV`指令可以被使用再如下指令中：
  - `ADD`
  - `COPY`
  - `ENV`
  - `EXPOSE`
  - `FROM`
  - `LABEL`
  - `STOPSIGNAL`
  - `USER`
  - `VOLUME`
  - `WORKDIR`
  - `ONBUILD`

#### 转义对于ENV变量的引用，使其为一个普通字符串

```dockerfile
FROM busybox
ENV FOO=/bar
WORKDIR ${FOO}   # WORKDIR /bar
ADD . $FOO       # ADD . /bar
COPY \$FOO /quux # COPY $FOO /quux 这里进行了转义
```

#### ENV的默认值功能

```dockerfile
FROM busybox
ENV FOO=/bar
# 如果定义了FOO环境变量，就使用mmmd替换FOO环境变量
RUN echo 111${FOO:+mmmd}111 # RUN echo 111mmmd111
# 如果没有定义FOO环境变量，就使用sdjj替换FOO环境变量
RUN echo 111${FOO:-sdjj}222 # RUN echo 111/bar222
```
