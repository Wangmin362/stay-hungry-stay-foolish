FROM centos:centos7.9.2009
MAINTAINER wangmin

WORKDIR /root

# 设置镜像root密码
RUN echo root:skyguard888 | chpasswd

# 更新yum源
RUN sed -e 's|^mirrorlist=|#mirrorlist=|g'  \
    -e 's|^#baseurl=http://mirror.centos.org|baseurl=https://mirrors.tuna.tsinghua.edu.cn|g'  \
    -i.bak /etc/yum.repos.d/CentOS-*.repo && yum clean all && yum makecache && yum update -y

# 安装需要的软件
RUN yum install -y net-tools wget ntp ntpdate telnet nmap sysstat lrszs dos2unix bind-utils bridge-utils  \
    bash-completion vim jq psmisc nfs-utils yum-utils device-mapper-persistent-data lvm2 git network-scripts  \
    tar iproute passwd openssl curl && ntpdate cn.pool.ntp.org

# TODO 升级内核
# 实际上，在容器内部是无法升级容器的内核版本的，只能通过升级宿主机内核版本的方式来升级内核

# 开启SSHD服务, 错误的构建方式，镜像在构建过程中不应该启动sshd服务，应该在镜像的启动点启动sshd服务
# 镜像在运行的时候需要开启特权，并且需要加上NET_ADMIN权限 docker run -d --name dev-env --privileged=true --cap-add=NET_ADMIN centos:centos7.9.2009
RUN yum install -y openssh-server openssh-clients && \
    sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config && \
    sed -i 's/#AddressFamily any/AddressFamily any/g' /etc/ssh/sshd_config && \
    sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config && \
    sed -i 's/#PermitRootLogin yes/PermitRootLogin yes/g' /etc/ssh/sshd_config && \
    sed -i "s/UsePAM.*/UsePAM no/g" /etc/ssh/sshd_config
#     这种方式是错误的构建方式，镜像在构建过程中不应该启动sshd服务，应该在镜像的启动点启动sshd服务
#    /bin/systemctl start sshd.service

RUN ssh-keygen -q -t rsa -b 2048 -f /etc/ssh/ssh_host_rsa_key -N '' && \
    ssh-keygen -q -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -N '' && \
    ssh-keygen -t dsa -f /etc/ssh/ssh_host_ed25519_key -N ''

# 创建启动sshd服务的脚本
RUN touch run.sh && echo '#!/bin/bash' >> run.sh && echo '/usr/sbin/sshd -D' >> run.sh && chmod 777 run.sh

# todo 挂载外部的docker.sock以及docker客户端，实现out-side of docker

WORKDIR /opt

# 安装git服务，并且生成ssh key
RUN wget https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.38.1.tar.gz --no-check-certificate
RUN yum install -y zlib-devel curl-devel bzip2-devel openssl-devel ncurses-devel gcc package automake autoconf  \
    make gcc-c++ cpio expat-devel gettext-devel zlib
RUN tar -zxf git-2.38.1.tar.gz && rm -f  git-2.38.1.tar.gz
WORKDIR git-2.38.1
RUN sh configure --prefix=/usr/local/git all && make && make install
RUN yum remove -y git &&  \
    echo "export PATH=$PATH:/usr/local/git/bin" >> /etc/bashrc && source /etc/bashrc &&  \
    git --version && \
    git config --global user.name "wangmin" && \
    git config --global user.email "wangmin@skyguard.com.cn" && \
    git config --global url.ssh://git@gitcdteam.skyguardmis.com/.insteadOf https://gitcdteam.skyguardmis.com/

WORKDIR /root

# golang环境
RUN wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz && mkdir go1.19.2 &&  \
    tar -zxvf go1.19.2.linux-amd64.tar.gz -C go1.19.2 && rm -f go1.19.2.linux-amd64.tar.gz
# fixme 配置golang环境 注意，这里不能使用/etc/profile来配置，否则docker run 之后会发现环境变量没有配置
RUN echo 'export GOROOT=/opt/go1.19.2/go' >>  /etc/bashrc &&  \
    echo 'export GOPATH=/root/go' >>  /etc/bashrc && \
    echo 'PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> /etc/bashrc && \
    source /etc/bashrc && \
    go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go env -w GOPRIVATE=gitcdteam.skyguardmis.com && \
    go env -w GOINSECURE=gitcdteam.skyguardmis.com


# python环境
#ADD Python-3.8.15.tgz python
#COPY Python-3.9.15.tgz python
#COPY Python-3.10.8.tgz python
#COPY Python-3.11.0.tgz python
# todo 配置python环境

# java环境
#COPY jdk-8u351-linux-x64.tar.gz java
#COPY jdk-11.0.17_linux-x64_bin.tar.gz java
#COPY jdk-17_linux-x64_bin.tar.gz java
#COPY jdk-19_linux-x64_bin.tar.gz java
# todo 配置java环境
# maven依赖
#COPY apache-maven-3.8.6-bin.tar.gz maven
# todo 配置maven环境
# gradle依赖

# nodejs依赖
#COPY node-v18.12.0-linux-x64.tar.xz nodejs
# todo 配置nodejs环境

# todo 安装kubectl客户端
# todo 安装helm客户端
# todo 安装ansible相关工具
# todo 安装kind，方便本地测试k8s集群新特性
# todo 考虑delve，即安装golang的远程调试环境
# todo linux性能优化工具，用于学习

WORKDIR /root
CMD ["/root/run.sh"]
#CMD ["/bin/bash"]
