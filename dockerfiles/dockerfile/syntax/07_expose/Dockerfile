FROM busybox

# 基本使用方式（没有声明协议，默认就是tcp协议）
EXPOSE 8001

# 设置该端口暴露那种协议
EXPOSE 8002/tcp

# 如果一个端口，须同时暴露两种协议，那么需要声明两次
EXPOSE 8003/tcp
EXPOSE 8003/udp

CMD ["/bin/sh", "-c", "ping baidu.com"]
