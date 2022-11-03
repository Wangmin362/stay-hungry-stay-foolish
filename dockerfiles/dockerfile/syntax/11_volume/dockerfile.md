
### 作用

- `Dockerfile`中声明的卷有何作用？表达了什么含义？
- 似乎感觉没啥作用，难道和`EXPOSE`指令一样，用于文档说明？？

```dockerfile
FROM busybox

RUN mkdir /myvol
RUN echo "hello world" > /myvol/greeting
# 如果在docker run的时候没有指定 -v映射外部的宿主机的目录，那么就是数据集一个临时目录映射进来，容器删除后，数据仍旧会被删除
VOLUME /myvol

CMD ping baidu.com
```
