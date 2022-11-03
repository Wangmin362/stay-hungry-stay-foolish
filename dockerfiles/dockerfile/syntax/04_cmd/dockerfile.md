
### CMD
- 语法：
  - `CMD ["executable","param1","param2"]`: exec格式，推荐使用这种格式
  - `CMD ["param1","param2"]`: 作为`ENTRYPOINT`的默认参数
  - `CMD command param1 param2`: shell格式
- 如果在一个`Dockerfile`中有多个`CMD`指令，那么只有最后一条`CMD`指令会被执行，其余的`CMD`指令会被忽略
- 如果在启动容器的时候传递了`command`，那么`Dockerfile`中的`CMD`指令会被忽略
  - `docker run -d -p 8080:8080 mybox echo run_command`：此时容器中的`CMD`就会被这里的`echo run_command`代替
- `CMD`的本质作用是为容器的运行提供一个默认的执行命令，所以`docker run`可以覆盖里面的容器启动命令
- 如果`Dockerfile`中没有指定任何的`CMD`，并且再`docker run`的时候也不指定启动命令，会发生什么呢？

#### 多个CMD指令

```dockerfile
FROM busybox

CMD echo cmd1
CMD echo cmd2
# 如果有多个CMD指令，只有最后一条命令会生效
CMD echo cmd3
```
