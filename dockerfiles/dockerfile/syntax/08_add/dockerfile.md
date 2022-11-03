
### ADD

- 基本语法：
  - `ADD [--chown=<user>:<group>] [--checksum=<checksum>] <src>... <dest>`
  - `ADD [--chown=<user>:<group>] ["<src>",... "<dest>"]`
- `chown`是用于改变添加文件的属主，以及组
- 可以一次性添加多个文件，源文件可以有通配符
- 源路径可以是一个`URL`，即从远端下载文件到镜像的某个目录
- 可以通过`--checksum`来校验文件的校验和


```dockerfile
FROM busybox

# 常规使用，即把abx.txt复制到容器的/opt目录下
ADD abc.txt /opt/

# 一次性添加多个文件
ADD aa.txt bb.txt cc.txt /opt/

# 源文件使用通配符匹配，即把当前目录中所有以txt结尾的文件复制到目录中
ADD *.txt /opt/


# 源文件使用通配符匹配, ?仅可以匹配一个字符
ADD abc?.txt /opt/

# 校验下载源文件的校验和
ADD --checksum=sha256:24454f830cdb571e2c4ad15481119c43b3cafd48dd869a9b2945d1036d1dc68d https://mirrors.edge.kernel.org/pub/linux/kernel/Historic/linux-0.01.tar.gz /
```

