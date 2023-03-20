FROM centos:centos7.9.2009
MAINTAINER wangmin

WORKDIR /root

# 设置镜像root密码
RUN echo root:skyguard888 | chpasswd

# 更新yum源
RUN sed \
    -e 's|^mirrorlist=|#mirrorlist=|g' \
    -e 's|^#baseurl=http://mirror.centos.org|baseurl=https://mirrors.tuna.tsinghua.edu.cn|g' \
    -i.bak /etc/yum.repos.d/CentOS-*.repo && \
    yum clean all && \
    yum makecache && \
    yum update -y && \
    yum install -y net-tools wget ntp ntpdate telnet nmap sysstat lrszs dos2unix bind-utils bridge-utils \
    bash-completion vim jq psmisc nfs-utils yum-utils device-mapper-persistent-data lvm2 git network-scripts \
    tar iproute passwd openssl curl zlib-devel curl-devel bzip2-devel openssl-devel ncurses-devel gcc package \
    automake autoconf make gcc-c++ cpio expat-devel gettext-devel zlib tree sqlite-devel readline-devel \
    tk-devel gdbm-devel db4-devel libpcap-devel xz-devel

# TODO 升级内核 实际上，在容器内部是无法升级容器的内核版本的，只能通过升级宿主机内核版本的方式来升级内核

# 开启SSHD服务, 错误的构建方式，镜像在构建过程中不应该启动sshd服务，应该在镜像的启动点启动sshd服务
# 镜像在运行的时候需要开启特权，并且需要加上NET_ADMIN权限 docker run -d --name dev-env --privileged=true --cap-add=NET_ADMIN centos:centos7.9.2009
RUN yum install -y openssh-server openssh-clients && \
    sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config && \
    sed -i 's/#AddressFamily any/AddressFamily any/g' /etc/ssh/sshd_config && \
    sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config && \
    sed -i 's/#PermitRootLogin yes/PermitRootLogin yes/g' /etc/ssh/sshd_config && \
    sed -i "s/UsePAM.*/UsePAM no/g" /etc/ssh/sshd_config && \
    ssh-keygen -q -t rsa -b 2048 -f /etc/ssh/ssh_host_rsa_key -N '' && \
    ssh-keygen -q -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -N '' && \
    ssh-keygen -t dsa -f /etc/ssh/ssh_host_ed25519_key -N ''

WORKDIR /opt

# git
ENV GITVERSION=2.40.0
RUN wget https://mirrors.edge.kernel.org/pub/software/scm/git/git-${GITVERSION}.tar.gz --no-check-certificate && \
    tar -zxf git-${GITVERSION}.tar.gz && \
    rm -f  git-${GITVERSION}.tar.gz && \
    cd git-${GITVERSION} && \
    sh configure --prefix=/usr/local/git all && \
    make && \
    make install &&  \
    cd .. && rm -rf  git-${GITVERSION} && \
    yum remove -y git &&  \
    echo "export PATH=$PATH:/usr/local/git/bin" >> /etc/bashrc && source /etc/bashrc &&  \
    git --version && \
    git config --global user.name "wangmin" && \
    git config --global user.email "wangmin@skyguard.com.cn" && \
    git config --global url.ssh://git@gitcdteam.skyguardmis.com/.insteadOf https://gitcdteam.skyguardmis.com/

WORKDIR /opt
# golang环境
ENV GOLANGVERSION=1.20.2
#COPY go${GOLANGVERSION}.linux-amd64.tar.gz .
RUN wget https://dl.google.com/go/go${GOLANGVERSION}.linux-amd64.tar.gz && \
    mkdir go${GOLANGVERSION} &&  \
    tar -zxf go${GOLANGVERSION}.linux-amd64.tar.gz -C go${GOLANGVERSION} && \
    rm -f go${GOLANGVERSION}.linux-amd64.tar.gz && \
    echo 'export GOROOT=/opt/go${GOLANGVERSION}/go' >>  /etc/bashrc &&  \
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
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest &&\
    # 安装delve，用于debug go代码
    go install github.com/go-delve/delve/cmd/dlv@latest

# todo java环境
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
ENV NODEJSVERSION=17.9.1
RUN wget https://nodejs.org/dist/v${NODEJSVERSION}/node-v${NODEJSVERSION}-linux-x64.tar.xz && \
    tar -xvf node-v${NODEJSVERSION}-linux-x64.tar.xz && \
    rm -f node-v${NODEJSVERSION}-linux-x64.tar.xz && \
    mv node-v${NODEJSVERSION}-linux-x64 /usr/local/nodejs && \
    echo 'PATH=$PATH:/usr/local/nodejs/bin' >> /etc/bashrc

# todo 配置nodejs环境

# todo 生成ssh key

# todo 安装kubectl客户端
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && \
    chmod +x kubectl && \
    mv kubectl /usr/local/bin && \
    yum install -y bash-completion && \
    source /usr/share/bash-completion/bash_completion && \
    echo "source <(kubectl completion bash)" >> ~/.bashrc
# todo dockefile中这个应该怎么做？？
#RUN source <(kubectl completion bash)

# 安装docker客户端
ENV DOCKERVERSION=20.10.23
RUN curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKERVERSION}.tgz && \
    tar xzvf docker-${DOCKERVERSION}.tgz --strip 1 -C /usr/local/bin docker/docker && \
    rm docker-${DOCKERVERSION}.tgz

# python环境
#RUN yum install -y python3
# todo 安装最新的python版本, 移除python2版本,并且把python命令链接到python3
WORKDIR /opt
ENV PYTHONVERSION=3.11.2
COPY Python-3.11.2.tar.xz ./
RUN #wget https://www.python.org/ftp/python/${PYTHONVERSION}/Python-${PYTHONVERSION}.tar.xz
RUN yum -y groupinstall "Development tools"
RUN tar -vxf Python-${PYTHONVERSION}.tar.xz && \
    cd Python-${PYTHONVERSION} && \
    sh configure --prefix=/usr/local/python3 --enable-optimizations && \
    make -j 8 && \
    make altinstall && \
    rm -rf  Python-${PYTHONVERSION}.tar.xz && \
    rm -rf Python-${PYTHONVERSION}

# 安装docker-compose工具
ENV DOCKERCOMPOSEVERSION=2.16.0
RUN curl -L "https://ghproxy.com/https://github.com/docker/compose/releases/download/v${DOCKERCOMPOSEVERSION}/docker-compose-linux-x86_64" -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose && \
    docker-compose --version

# 安装helm客户端 这里通过网络去下载，相当缓慢
ENV HELMVERSION=v3.11.2
#RUN wget https://get.helm.sh/helm-${HELMVERSION}-linux-amd64.tar.gz
COPY helm-v3.11.2-linux-amd64.tar.gz .
RUN tar -zxvf helm-${HELMVERSION}-linux-amd64.tar.gz && \
    mv linux-amd64/helm /usr/local/bin/helm && \
    helm help && rm -f helm-${HELMVERSION}-linux-amd64.tar.gz rm linux-amd64 -rf

# todo 安装ansible相关工具
# todo linux性能优化工具，用于学习

WORKDIR /root

# 指定终端颜色
RUN echo 'PS1="\[\e[37m\][\[\e[32m\]\u\[\e[37m\]@\[\e[35m\]\h\[\e[0m\] \[\e[36m\]\w\[\e[0m\]]\\$ "' >> ~/.bashrc

# 用于指定需要容器运行时需要启动的服务
RUN touch run.sh &&  \
    echo '#!/bin/bash' >> run.sh &&  \
    echo 'ntpdate cn.pool.ntp.org' >> run.sh &&  \
    # 这里会后台运行，容器不会退出，用户也不需要指定Command
    echo '/usr/sbin/sshd -D' >> run.sh &&  \
    chmod +x run.sh

# 这里如果是CMD命令，那么会被docker run <command>的command覆盖掉
ENTRYPOINT ["/root/run.sh"]
#CMD ["/bin/bash"]
