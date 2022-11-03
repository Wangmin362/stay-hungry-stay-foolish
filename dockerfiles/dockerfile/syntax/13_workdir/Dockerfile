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

