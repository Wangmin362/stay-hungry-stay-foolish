FROM ubuntu:22.04

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
    git build-essential flex libncurses-dev bison libelf-dev libssl-dev bc

WORKDIR /opt

# 开启SSHD服务, 错误的构建方式，镜像在构建过程中不应该启动sshd服务，应该在镜像的启动点启动sshd服务
# 镜像在运行的时候需要开启特权，并且需要加上NET_ADMIN权限 docker run -d --name dev-env --privileged=true --cap-add=NET_ADMIN centos:centos7.9.2009
RUN sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config && \
    sed -i 's/#AddressFamily any/AddressFamily any/g' /etc/ssh/sshd_config && \
    sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config && \
    sed -i 's/#PermitRootLogin yes/PermitRootLogin yes/g' /etc/ssh/sshd_config && \
    sed -i "s/UsePAM.*/UsePAM no/g" /etc/ssh/sshd_config && \
    ssh-keygen -q -t rsa -b 2048 -f /etc/ssh/ssh_host_rsa_key -N '' && \
    ssh-keygen -q -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -N '' && \
    ssh-keygen -t dsa -f /etc/ssh/ssh_host_ed25519_key -N ''

# 安装git
#git --version && \
#    git config --global user.name "wangmin" && \
#    git config --global user.email "wangmin@skyguard.com.cn" && \
#    git config --global url.ssh://git@gitcdteam.skyguardmis.com/.insteadOf https://gitcdteam.skyguardmis.com/


# golang环境
ENV GOLANGVERSION=1.20.6
#COPY go${GOLANGVERSION}.linux-amd64.tar.gz .
RUN wget https://mirrors.aliyun.com/golang/go${GOLANGVERSION}.linux-amd64.tar.gz?spm=a2c6h.25603864.0.0.a6b07c45FOi9wZ && \
    mkdir go${GOLANGVERSION} &&  \
    tar -zxf go${GOLANGVERSION}.linux-amd64.tar.gz -C go${GOLANGVERSION} && \
    rm -f go${GOLANGVERSION}.linux-amd64.tar.gz && \
    echo "export GOROOT=/opt/go${GOLANGVERSION}/go" >>  /etc/bashrc &&  \
    echo 'export GOPATH=/root/go' >>  /etc/bashrc && \
    echo 'PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> /etc/bashrc && \
    source /etc/bashrc && \
    go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
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
RUN echo 'PS1="\[\e[37m\][\[\e[32m\]\u\[\e[37m\]@\[\e[35m\]\h\[\e[0m\] \[\e[36m\]\w\[\e[0m\]]\\$ "' >> ~/.bashrc

# 打开文件监听限制，否则当文件数量过多，linux会限制打开文件数量
RUN touch /etc/sysctl.conf && \
    echo "fs.inotify.max_user_watches = 9994288" > /etc/sysctl.conf

# vscode 不然vscode安装好之后，老是找不到命令
echo "code_latest_version=\$(ls -tral -1 --ignore=.* ~/.vscode-server/bin | sed -n '2p' | rev | cut -d' ' -f1 | rev)\n\
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