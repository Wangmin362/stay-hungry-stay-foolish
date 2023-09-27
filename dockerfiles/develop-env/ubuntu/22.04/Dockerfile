FROM ubuntu:22.04

# TODO 设置镜像ssh root登录密码
RUN echo root:123456 | chpasswd

USER root
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

RUN echo '[Unit]\n\
Description=sync time\n\
After=chrony.service\n\
[Service]\n\
ExecStart=timedatectl set-timezone "Asia/Shanghai"\n\
[Install]\n\
WantedBy=multi-user.target' > /etc/systemd/system/sync-time.service


RUN apt update -y && apt upgrade -y
RUN DEBIAN_FRONTEND=noninteractive TZ=Asia/Shanghai apt -y install tzdata
RUN apt install -y net-tools telnet sysstat bridge-utils bash-completion vim jq tar openssl iputils-ping lsof lvm2 \
    dnsutils curl gcc g++ automake autoconf make tree stress htop atop sysbench chrony ipvsadm ipset conntrack ufw  \
    git build-essential flex libncurses-dev bison libelf-dev libssl-dev bc openssh-server bash-completion

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
# 如果dl.google.com解析不了，需要添加dns解析，那么可以使用如下命令：docker build --add-host=dl.google.com:58.63.233.97  --progress=plain .
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
    versions='1.13.15 1.14.15 1.15.15 1.16.15 1.17.13 1.18.10 1.19.13 1.20.8 1.21.0' && \
    priority=1 && for V in ${versions}; do \
    /usr/local/go/bin/go install golang.org/dl/go$V@latest && go$V download && \
    update-alternatives  --install /usr/bin/go go /root/sdk/go$V/bin/go $priority && ((priority=priority+1)); done && \
    update-alternatives  --install /usr/bin/go go /usr/local/go/bin/go $priority && \
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
    go install github.com/google/wire/cmd/wire@latest && \
    # boltdb查看工具
    go install github.com/br0xen/boltbrowser@latest


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
RUN versions='v13.14.0 v14.21.3 v15.14.0 v16.20.2 v17.9.1 v18.8.0 v19.9.0 v20.7.0' && \
    priority=1 && for V in ${versions}; do \
    wget https://mirrors.huaweicloud.com/nodejs/$V/node-$V-linux-x64.tar.xz && \
    rm -rf /usr/local/node-$V && mkdir /usr/local/node-$V && tar -C /usr/local/node-$V -xvJf node-$V-linux-x64.tar.xz --strip-components 1 && \
    update-alternatives --install /usr/bin/node nodejs /usr/local/node-$V/bin/node $priority && ((priority=priority+1)); done && \
    update-alternatives --display nodejs && rm -f /root/node*
    # TODO 考虑是否有必要配置nodejs的环境变量

# 安装kubectl客户端
RUN versions='v1.10.13 v1.11.10 v1.12.10 v1.13.12 v1.14.10 v1.15.12 v1.16.15 v1.17.17 v1.18.20 v1.19.16 v1.20.15 v1.21.14 v1.22.17 v1.23.17 v1.24.17 v1.25.14 v1.26.9 v1.27.6 v1.28.2' && \
    priority=1 && for V in ${versions}; do \
    mkdir /usr/local/k8s-$V && curl -SL https://files.m.daocloud.io/storage.googleapis.com/kubernetes-release/release/$V/bin/linux/amd64/kubectl \
    -o /usr/local/k8s-$V/kubectl && chmod +x /usr/local/k8s-$V/kubectl && \
    update-alternatives --install /usr/bin/kubectl kubectl /usr/local/k8s-$V/kubectl $priority && ((priority=priority+1)); done && \
    update-alternatives --display kubectl && echo "source <(kubectl completion bash)" >> ~/.bashrc
# todo dockefile中这个应该怎么做？？
#RUN source <(kubectl completion bash)

# 安装docker客户端
RUN versions='18.09.9 19.03.15 20.10.24 23.0.6 24.0.6' && \
    priority=1 && for V in ${versions}; do \
    curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-$V.tgz && \
    mkdir /usr/local/docker-$V && tar xzvf docker-$V.tgz --strip 1 -C /usr/local/docker-$V docker/docker && \
    update-alternatives --install /usr/bin/docker docker /usr/local/docker-$V/docker $priority && ((priority=priority+1)); done && \
    update-alternatives --display docker && \
    rm -rf /usr/libexec/docker/cli-plugins && mkdir -p /usr/libexec/docker/cli-plugins && \
    curl -SL https://ghproxy.com/https://github.com/docker/buildx/releases/download/v0.11.2/buildx-v0.11.2.linux-amd64 -o /usr/libexec/docker/cli-plugins/docker-buildx && \
    chmod +x /usr/libexec/docker/cli-plugins/docker-buildx && \
    curl -SL https://ghproxy.com/https://github.com/docker/compose/releases/download/v2.21.0/docker-compose-linux-x86_64 -o /usr/libexec/docker/cli-plugins/docker-compose && \
    chmod +x /usr/libexec/docker/cli-plugins/docker-compose && \
    rm -f /root/docker* && docker buildx version
    # docker 命令补全


# Python环境
RUN versions='2.7 3.10 3.11' && \
    priority=1 && for V in ${versions}; do apt install -y python$V && \
    update-alternatives --install /usr/bin/python python /usr/bin/python$V $priority && ((priority=priority+1)); done && \
    update-alternatives --display python && \
    apt install -y python3-pip && pip --version && pip --version

# 安装Helm工具
ENV HELMVERSION=v3.12.3
RUN wget https://files.m.daocloud.io/get.helm.sh/helm-${HELMVERSION}-linux-amd64.tar.gz
RUN rm -f /usr/local/bin/helm && tar -zxvf helm-${HELMVERSION}-linux-amd64.tar.gz --strip 1 -C /usr/local/bin linux-amd64/helm && \
    helm version && rm -rf helm*

# todo 安装ansible相关工具
RUN apt install -y ansible && ansible --version

# 配置chrony，使用国内时间同步源
RUN sed -i 's/ntp.ubuntu.com/ntp1.aliyun.com/' /etc/chrony/chrony.conf && \
    sed -i 's/0.ubuntu.pool.ntp.org/ntp2.aliyun.com/' /etc/chrony/chrony.conf && \
    sed -i 's/1.ubuntu.pool.ntp.org/ntp3.aliyun.com/' /etc/chrony/chrony.conf && \
    sed -i 's/2.ubuntu.pool.ntp.org/ntp4.aliyun.com/' /etc/chrony/chrony.conf && \
    if ! cat /etc/chrony/chrony.conf | grep ntp5.aliyun.com; then \
    sed -i "24 a server ntp5.aliyun.com iburst\nserver ntp6.aliyun.com iburst\nserver ntp7.aliyun.com iburst\nserver 0.cn.pool.ntp.org iburst\nserver 1.cn.pool.ntp.org iburst\nserver 2.cn.pool.ntp.org iburst\nserver 3.cn.pool.ntp.org iburst\nserver time1.cloud.tencent.com iburst\nserver time2.cloud.tencent.com iburst\nserver time3.cloud.tencent.com iburst\nserver time4.cloud.tencent.com iburst\n\n"  /etc/chrony/chrony.conf; fi

# 指定终端颜色
# TODO 如果当前目录是git管理的，那么显示当前版本
RUN echo 'PS1="\[\e[37m\][\[\e[32m\]\u\[\e[37m\]@\[\e[35m\]\h\[\e[0m\] \[\e[36m\]\w\[\e[0m\]]\\$ "' >> ~/.bashrc

# 打开文件监听限制，否则当文件数量过多，linux会限制打开文件数量
RUN touch /etc/sysctl.conf && \
    echo "fs.inotify.max_user_watches = 9994288" > /etc/sysctl.conf

# vscode 不然vscode安装好之后，老是找不到命令
RUN echo "code_latest_version=\$(ls -tral -1 --ignore=.* ~/.vscode-server/bin | sed -n '2p' | rev | cut -d' ' -f1 | rev)\n\
    export PATH=\${HOME}/.vscode-server/bin/\${code_latest_version}/bin/remote-cli:\$PATH" >> ~/.bashrc

RUN cat /etc/apt/sources.list
RUN cat /etc/systemd/system/sync-time.service


# docker启动容器：docker build -t systemd_test .
# 测试：    docker run -it --rm --privileged --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private  ubuntu-dev:22.04
# 实际使用：

CMD [ "/lib/systemd/systemd", "log-level=info", "unit=sysinit.target" ]