FROM ubuntu:22.04

# TODO 设置镜像ssh root登录密码
RUN echo root:123456 | chpasswd

WORKDIR /root
STOPSIGNAL SIGRTMIN+3

# 移除无用的服务
RUN rm -f /lib/systemd/system/multi-user.target.wants/* \
  /etc/systemd/system/*.wants/* \
  /lib/systemd/system/local-fs.target.wants/* \
  /lib/systemd/system/sockets.target.wants/*udev* \
  /lib/systemd/system/sockets.target.wants/*initctl* \
  /lib/systemd/system/sysinit.target.wants/systemd-tmpfiles-setup* \
  /lib/systemd/system/systemd-update-utmp*

RUN echo 'deb http://cn.archive.ubuntu.com/ubuntu jammy main restricted\n\
deb http://cn.archive.ubuntu.com/ubuntu jammy-updates main restricted\n\
deb http://cn.archive.ubuntu.com/ubuntu jammy universe\n\
deb http://cn.archive.ubuntu.com/ubuntu jammy-updates universe\n\
deb http://cn.archive.ubuntu.com/ubuntu jammy multiverse\n\
deb http://cn.archive.ubuntu.com/ubuntu jammy-updates multiverse\n\
deb http://cn.archive.ubuntu.com/ubuntu jammy-backports main restricted universe multiverse\n\
deb http://cn.archive.ubuntu.com/ubuntu jammy-security main restricted\n\
deb http://cn.archive.ubuntu.com/ubuntu jammy-security universe\n\
deb http://cn.archive.ubuntu.com/ubuntu jammy-security multiverse' > /etc/apt/sources.list

RUN apt update -y && apt upgrade -y
RUN DEBIAN_FRONTEND=noninteractive TZ=Asia/Shanghai apt -y install tzdata
RUN apt install -y net-tools telnet sysstat bridge-utils bash-completion vim jq tar openssl iputils-ping lsof lvm2 \
    dnsutils curl gcc g++ automake autoconf make tree stress htop atop sysbench chrony ipvsadm ipset conntrack ufw  \
    git build-essential flex libncurses-dev bison libelf-dev libssl-dev bc openssh-server

# 开启SSHD服务, 错误的构建方式，镜像在构建过程中不应该启动sshd服务，应该在镜像的启动点启动sshd服务 (当然，如果使用Systemd作为一号进程，那么ssh服务自然会被启动)
# 镜像在运行的时候需要开启特权，并且需要加上NET_ADMIN权限 docker run -d --name dev-env --privileged=true --cap-add=NET_ADMIN centos:centos7.9.2009
RUN sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config && \
    sed -i 's/#AddressFamily any/AddressFamily any/g' /etc/ssh/sshd_config && \
    sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config && \
    sed -i 's/#PermitRootLogin yes/PermitRootLogin yes/g' /etc/ssh/sshd_config && \
    sed -i "s/UsePAM.*/UsePAM no/g" /etc/ssh/sshd_config

# TODO 多版本安装软件
# gcc g++ gdb glibc kubebuilder go  docker-cli kubectl make cmake python helm
# 中国时区

# 切换脚本解释器为bash, Ubuntu默认的脚本解释器为dash
RUN ln -fs /bin/bash /bin/sh

# golang环境 https://go.dev/doc/install  多版本管理 https://blog.csdn.net/weixin_41910261/article/details/112969035
ENV GOLANGVERSION=1.21.1
RUN wget https://mirrors.aliyun.com/golang/go${GOLANGVERSION}.linux-amd64.tar.gz?spm=a2c6h.25603864.0.0.a6b07c45EjitSM -O go${GOLANGVERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go${GOLANGVERSION}.linux-amd64.tar.gz && \
    rm -f go${GOLANGVERSION}.linux-amd64.tar.gz && \
    echo 'export GOPATH=/root/go' >>  /etc/bashrc && \
    echo 'PATH=$PATH:$GOPATH/bin' >> /etc/bashrc && \
    source /etc/bashrc && \
    /usr/local/go/bin/go env -w GO111MODULE=on && \
    /usr/local/go/bin/go env -w GOPROXY=https://goproxy.cn,direct && \
    # 安装不同的go版本，并使用update-alternatives工具进行切换
    /usr/local/go/bin/go install golang.org/dl/go1.13.15@latest && /root/go/bin/go1.13.15 download && \
    /usr/local/go/bin/go install golang.org/dl/go1.14.15@latest && /root/go/bin/go1.14.15 download && \
    /usr/local/go/bin/go install golang.org/dl/go1.15.15@latest && /root/go/bin/go1.15.15 download && \
    /usr/local/go/bin/go install golang.org/dl/go1.16.15@latest && /root/go/bin/go1.16.15 download && \
    /usr/local/go/bin/go install golang.org/dl/go1.17.13@latest && /root/go/bin/go1.17.13 download && \
    /usr/local/go/bin/go install golang.org/dl/go1.18.10@latest && /root/go/bin/go1.18.10 download && \
    /usr/local/go/bin/go install golang.org/dl/go1.19.13@latest && /root/go/bin/go1.19.13 download && \
    /usr/local/go/bin/go install golang.org/dl/go1.20.8@latest && /root/go/bin/go1.20.8 download && \
    /usr/local/go/bin/go install golang.org/dl/go1.21.0@latest && /root/go/bin/go1.21.0 download && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go1.13.15/bin/go 1 && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go1.14.15/bin/go 2 && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go1.15.15/bin/go 3 && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go1.16.15/bin/go 4 && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go1.17.13/bin/go 5 && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go1.18.10/bin/go 6 && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go1.19.13/bin/go 7 && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go1.20.8/bin/go 8 && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go1.21.0/bin/go 9 && \
    update-alternatives  --install /usr/bin/go go /usr/local/go/bin/go 10 && \
    update-alternatives --display go && \
    # 私有仓库
    go env -w GOPRIVATE=gitcdteam.skyguardmis.com && \
    go env -w GOINSECURE=gitcdteam.skyguardmis.com && \
    # 本地k8s测试工具
    go install sigs.k8s.io/kind@latest && kind --version && \
    # 漏洞检测工具
    go install golang.org/x/vuln/cmd/govulncheck@latest && \
    # protoc编译器
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    # grpc
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    # gateway
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
    # openapi
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && \
    # 安装delve，用于debug go代码
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    # 安装wire自动注入工具
    go install github.com/google/wire/cmd/wire@latest
# go install boltdbbrower

RUN ls -ll /root/sdk
RUN AA
# RUN bolt工具
# todo java环境
#COPY jdk-17_linux-x64_bin.tar.gz java
#COPY jdk-19_linux-x64_bin.tar.gz java
# todo 配置java环境
# maven依赖
#COPY apache-maven-3.8.6-bin.tar.gz maven
# todo 配置maven环境
# gradle依赖

# nodejs依赖
ENV NODEJSVERSION=17.9.1
RUN wget https://nodejs.org/dist/v${NODEJSVERSION}/node-v${NODEJSVERSION}-linux-x64.tar.xz && \
    tar -xvf node-v${NODEJSVERSION}-linux-x64.tar.xz && \
    rm -f node-v${NODEJSVERSION}-linux-x64.tar.xz && \
    mv node-v${NODEJSVERSION}-linux-x64 /usr/local/nodejs && \
    echo 'PATH=$PATH:/usr/local/nodejs/bin' >> /etc/bashrc

# 安装kubectl客户端
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && \
    chmod +x kubectl && \
    mv kubectl /usr/local/bin && \
    yum install -y bash-completion && \
    source /usr/share/bash-completion/bash_completion && \
    echo "source <(kubectl completion bash)" >> ~/.bashrc
# todo dockefile中这个应该怎么做？？
#RUN source <(kubectl completion bash)

# 安装docker客户端
ENV DOCKERVERSION=24.0.5
RUN curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKERVERSION}.tgz && \
    tar xzvf docker-${DOCKERVERSION}.tgz --strip 1 -C /usr/local/bin docker/docker && \
    rm docker-${DOCKERVERSION}.tgz

# python环境
RUN ln -s /usr/bin/python3 /usr/bin/python


# 安装helm客户端 这里通过网络去下载，相当缓慢
ENV HELMVERSION=v3.12.2
RUN wget https://files.m.daocloud.io/get.helm.sh/helm-${HELMVERSION}-linux-amd64.tar.gz
RUN tar -zxvf helm-${HELMVERSION}-linux-amd64.tar.gz && \
    mv linux-amd64/helm /usr/local/bin/helm && \
    helm help && rm -f helm-${HELMVERSION}-linux-amd64.tar.gz rm linux-amd64 -rf

# todo 安装ansible相关工具
# todo linux性能优化工具，用于学习

WORKDIR /root

# 指定终端颜色
# TODO 如果当前目录是git管理的，那么显示当前版本
RUN echo 'PS1="\[\e[37m\][\[\e[32m\]\u\[\e[37m\]@\[\e[35m\]\h\[\e[0m\] \[\e[36m\]\w\[\e[0m\]]\\$ "' >> ~/.bashrc

# 打开文件监听限制，否则当文件数量过多，linux会限制打开文件数量
RUN touch /etc/sysctl.conf && \
    echo "fs.inotify.max_user_watches = 9994288" > /etc/sysctl.conf

# vscode 不然vscode安装好之后，老是找不到命令
RUN echo "code_latest_version=\$(ls -tral -1 --ignore=.* ~/.vscode-server/bin | sed -n '2p' | rev | cut -d' ' -f1 | rev)\n\
    export PATH=\${HOME}/.vscode-server/bin/\${code_latest_version}/bin/remote-cli:\$PATH" >> ~/.bashrc



# 用于指定需要容器运行时需要启动的服务
RUN touch run.sh &&  \
    echo '#!/bin/bash' >> run.sh &&  \
    echo 'ntpdate cn.pool.ntp.org' >> run.sh &&  \
    echo 'source /etc/bashrc' >> run.sh &&  \
    echo 'sysctl -p --system' >> run.sh &&  \
    echo 'systemctl start ntpd && systemctl enable ntpd && systemctl status ntpd' >> run.sh && \
    # 这里会后台运行，容器不会退出，用户也不需要指定Command
    echo '/usr/sbin/sshd -D' >> run.sh &&  \
    chmod +x run.sh

# 这里如果是CMD命令，那么会被docker run <command>的command覆盖掉
ENTRYPOINT ["/root/run.sh"]
CMD [ "/lib/systemd/systemd", "log-level=info", "unit=sysinit.target" ]