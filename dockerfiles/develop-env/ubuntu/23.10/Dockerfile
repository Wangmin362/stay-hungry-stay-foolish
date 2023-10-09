FROM ubuntu:23.10

# 设置容器登录密码
RUN echo root:123456 | chpasswd

USER root
WORKDIR /root
ENV container docker
STOPSIGNAL SIGRTMIN+3
VOLUME [ "/tmp", "/run", "/run/lock" ]

# 移除无用的服务
RUN rm -f /lib/systemd/system/multi-user.target.wants/* \
  /etc/systemd/system/*.wants/* \
  /lib/systemd/system/local-fs.target.wants/* \
  /lib/systemd/system/sockets.target.wants/*udev* \
  /lib/systemd/system/sockets.target.wants/*initctl* \
  /lib/systemd/system/sysinit.target.wants/systemd-tmpfiles-setup* \
  /lib/systemd/system/systemd-update-utmp*

RUN echo 'deb http://cn.archive.ubuntu.com/ubuntu mantic main restricted\n\
deb http://cn.archive.ubuntu.com/ubuntu mantic-updates main restricted\n\
deb http://cn.archive.ubuntu.com/ubuntu mantic universe\n\
deb http://cn.archive.ubuntu.com/ubuntu mantic-updates universe\n\
deb http://cn.archive.ubuntu.com/ubuntu mantic multiverse\n\
deb http://cn.archive.ubuntu.com/ubuntu mantic-updates multiverse\n\
deb http://cn.archive.ubuntu.com/ubuntu mantic-backports main restricted universe multiverse\n\
deb http://cn.archive.ubuntu.com/ubuntu mantic-security main restricted\n\
deb http://cn.archive.ubuntu.com/ubuntu mantic-security universe\n\
deb http://cn.archive.ubuntu.com/ubuntu mantic-security multiverse' > /etc/apt/sources.list

# 不知道为什么，这个必须写在这里，否则文件中的内容不会换行，不知道是不是因为把dash改为bash的原因
# 容器启动后执行：timedatectl set-timezone "Asia/Shanghai" 命令，设置时区为中国
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
    git build-essential flex libncurses-dev bison libelf-dev libssl-dev bc openssh-server bash-completion ansible \ 
    unzip libc-ares-dev libtool pkg-config libsystemd-dev file binutils texinfo


# github加速（非常重要，否则github下载代码将会相当之慢）
RUN git config --global url."https://hub.fgit.cf".insteadof "https://github.com"


# 开启SSHD服务, 错误的构建方式，镜像在构建过程中不应该启动sshd服务，应该在镜像的启动点启动sshd服务 (当然，如果使用Systemd作为一号进程，那么ssh服务自然会被启动)
RUN sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config && \
    sed -i 's/#AddressFamily any/AddressFamily any/g' /etc/ssh/sshd_config && \
    sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config && \
    sed -i 's/#PermitRootLogin yes/PermitRootLogin yes/g' /etc/ssh/sshd_config && \
    sed -i "s/UsePAM.*/UsePAM no/g" /etc/ssh/sshd_config

# 切换脚本解释器为bash, Ubuntu默认的脚本解释器为dash
RUN ln -fs /bin/bash /bin/sh

# golang环境 https://go.dev/doc/install  多版本管理 https://blog.csdn.net/weixin_41910261/article/details/112969035
# 如果dl.google.com解析不了，需要添加dns解析，那么可以使用如下命令：docker build --add-host=dl.google.com:58.63.233.97  --progress=plain .
# 使用此命令切换go的版本：update-alternatives --config go
ENV GOLANGVERSION=1.21.1
RUN curl -SL https://mirrors.aliyun.com/golang/go${GOLANGVERSION}.linux-amd64.tar.gz?spm=a2c6h.25603864.0.0.a6b07c45EjitSM -o go${GOLANGVERSION}.linux-amd64.tar.gz && \
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


# todo java环境
#COPY jdk-17_linux-x64_bin.tar.gz java
#COPY jdk-19_linux-x64_bin.tar.gz java
# todo 配置java环境
# maven依赖
#COPY apache-maven-3.8.6-bin.tar.gz maven
# todo 配置maven环境
# gradle依赖s


# 安装NodeJS环境
# 使用此命令切换nodejs的版本：update-alternatives --config nodejs
RUN versions='v13.14.0 v14.21.3 v15.14.0 v16.20.2 v17.9.1 v18.8.0 v19.9.0 v20.7.0' && \
    priority=1 && for V in ${versions}; do \
    curl -SL https://mirrors.huaweicloud.com/nodejs/$V/node-$V-linux-x64.tar.xz -o node-$V-linux-x64.tar.xz && \
    rm -rf /usr/local/node-$V && mkdir /usr/local/node-$V && tar -C /usr/local/node-$V -xvJf node-$V-linux-x64.tar.xz --strip-components 1 && \
    update-alternatives --install /usr/bin/node nodejs /usr/local/node-$V/bin/node $priority && ((priority=priority+1)); done && \
    update-alternatives --display nodejs && rm -f /root/node*
    # TODO 考虑是否有必要配置nodejs的环境变量

# 安装Kubectl客户端
# 使用此命令切换kubectl的版本：update-alternatives --config kubectl
RUN versions='v1.10.13 v1.11.10 v1.12.10 v1.13.12 v1.14.10 v1.15.12 v1.16.15 v1.17.17 v1.18.20 v1.19.16 v1.20.15 v1.21.14 v1.22.17 v1.23.17 v1.24.17 v1.25.14 v1.26.9 v1.27.6 v1.28.2' && \
    priority=1 && for V in ${versions}; do \
    mkdir /usr/local/k8s-$V && curl -SL https://files.m.daocloud.io/storage.googleapis.com/kubernetes-release/release/$V/bin/linux/amd64/kubectl \
    -o /usr/local/k8s-$V/kubectl && chmod +x /usr/local/k8s-$V/kubectl && \
    update-alternatives --install /usr/bin/kubectl kubectl /usr/local/k8s-$V/kubectl $priority && ((priority=priority+1)); done && \
    update-alternatives --display kubectl && echo "source <(kubectl completion bash)" >> ~/.bashrc
# todo dockefile中这个应该怎么做？？
#RUN source <(kubectl completion bash)

# 安装Docker客户端
# 使用此命令切换docker的版本：update-alternatives --config docker
RUN versions='18.09.9 19.03.15 20.10.24 23.0.6 24.0.6' && \
    priority=1 && for V in ${versions}; do \
    curl -SL https://download.docker.com/linux/static/stable/x86_64/docker-$V.tgz -o docker-$V.tgz && \
    mkdir /usr/local/docker-$V && tar -xzf docker-$V.tgz --strip 1 -C /usr/local/docker-$V docker/docker && \
    update-alternatives --install /usr/bin/docker docker /usr/local/docker-$V/docker $priority && ((priority=priority+1)); done && \
    update-alternatives --display docker && \
    rm -rf /usr/libexec/docker/cli-plugins && mkdir -p /usr/libexec/docker/cli-plugins && \
    curl -SL https://files.m.daocloud.io/github.com/docker/buildx/releases/download/v0.11.2/buildx-v0.11.2.linux-amd64 -o /usr/libexec/docker/cli-plugins/docker-buildx && \
    chmod +x /usr/libexec/docker/cli-plugins/docker-buildx && \
    curl -SL https://files.m.daocloud.io/github.com/docker/compose/releases/download/v2.21.0/docker-compose-linux-x86_64 -o /usr/libexec/docker/cli-plugins/docker-compose && \
    chmod +x /usr/libexec/docker/cli-plugins/docker-compose && \
    rm -f /root/docker* && docker buildx version
    # docker 命令补全

# 安装Python环境
# 使用此命令切换python的版本：update-alternatives --config python
RUN versions='2.7 3.10 3.11' && \
    priority=1 && for V in ${versions}; do apt install -y python$V && \
    update-alternatives --install /usr/bin/python python /usr/bin/python$V $priority && ((priority=priority+1)); done && \
    update-alternatives --display python && \
    apt install -y python3-pip && pip --version && pip --version
    # pip镜像加速设置
    # pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/ && \
    # pip config set global.timeout 6000

# 安装Helm工具
ENV HELMVERSION=v3.12.3
RUN curl -SL https://files.m.daocloud.io/get.helm.sh/helm-${HELMVERSION}-linux-amd64.tar.gz -o helm-${HELMVERSION}-linux-amd64.tar.gz && \
    rm -f /usr/local/bin/helm && tar -zxf helm-${HELMVERSION}-linux-amd64.tar.gz --strip 1 -C /usr/local/bin linux-amd64/helm && \
    helm version && rm -rf helm*

# 配置chrony，使用国内时间同步源
RUN sed -i 's/ntp.ubuntu.com/ntp1.aliyun.com/' /etc/chrony/chrony.conf && \
    sed -i 's/0.ubuntu.pool.ntp.org/ntp2.aliyun.com/' /etc/chrony/chrony.conf && \
    sed -i 's/1.ubuntu.pool.ntp.org/ntp3.aliyun.com/' /etc/chrony/chrony.conf && \
    sed -i 's/2.ubuntu.pool.ntp.org/ntp4.aliyun.com/' /etc/chrony/chrony.conf && \
    if ! cat /etc/chrony/chrony.conf | grep ntp5.aliyun.com; then \
    sed -i "24 a server ntp5.aliyun.com iburst\nserver ntp6.aliyun.com iburst\nserver ntp7.aliyun.com iburst\nserver 0.cn.pool.ntp.org iburst\nserver 1.cn.pool.ntp.org iburst\nserver 2.cn.pool.ntp.org iburst\nserver 3.cn.pool.ntp.org iburst\nserver time1.cloud.tencent.com iburst\nserver time2.cloud.tencent.com iburst\nserver time3.cloud.tencent.com iburst\nserver time4.cloud.tencent.com iburst\n\n"  /etc/chrony/chrony.conf; fi

# 安装Kubebuilder工具
# 使用此命令切换kubebuilder的版本：update-alternatives --config kubebuilder
RUN versions='v3.2.0 v3.4.0 v3.6.0 v3.9.1 v3.10.0 v3.11.1 v3.12.0' && \
    priority=1 && mkdir /usr/local/kubebuilder && for V in ${versions}; do \
    curl -SL https://files.m.daocloud.io/github.com/kubernetes-sigs/kubebuilder/releases/download/$V/kubebuilder_linux_amd64 \
    -o /usr/local/kubebuilder/kubebuilder-$V && chmod +x /usr/local/kubebuilder/kubebuilder-$V && \
    update-alternatives --install /usr/bin/kubebuilder kubebuilder /usr/local/kubebuilder/kubebuilder-$V $priority && ((priority=priority+1)); done && \
    update-alternatives --display kubebuilder

# 安装Protoc工具
# 使用此命令切换protoc的版本：update-alternatives --config protoc
RUN versions='21.9 22.0 22.5 23.0 23.4 24.0 24.3' && \
    priority=1 && for V in ${versions}; do mkdir /usr/local/protoc-v$V && \
    curl -SL https://files.m.daocloud.io/github.com/protocolbuffers/protobuf/releases/download/v$V/protoc-$V-linux-x86_64.zip -o protoc-$V-linux-x86_64.zip && \
    unzip -d /usr/local/protoc-v$V protoc-$V-linux-x86_64.zip && rm -f /root/protoc* && \
    update-alternatives --install /usr/bin/protoc protoc /usr/local/protoc-v$V/bin/protoc $priority && ((priority=priority+1)); done && \
    update-alternatives --display protoc


# 安装cmake工具
# 使用此命令切换cmake的版本：update-alternatives --config cmake
RUN versions='3.12.4 3.19.8 3.20.6 3.25.3 3.26.5 3.27.6' && \
    priority=1 && for V in ${versions}; do mkdir /usr/local/cmake-v$V && \
    curl -SL https://files.m.daocloud.io/github.com/Kitware/CMake/releases/download/v$V/cmake-$V-linux-x86_64.tar.gz -o cmake-$V-linux-x86_64.tar.gz && \
    tar -C /usr/local/cmake-v$V -xzf cmake-$V-linux-x86_64.tar.gz  --strip-components 1 && rm -f /root/cmake* && \
    update-alternatives --install /usr/bin/cmake cmake /usr/local/cmake-v$V/bin/cmake $priority && ((priority=priority+1)); done && \
    update-alternatives --display cmake


# 安装C++, Python, Ruby, Objective-C, PHP, C#语言的grpc插件
# 参考连接：https://blog.csdn.net/Aidam_Bo/article/details/117568034
RUN git clone -b v1.58.1 https://github.com/grpc/grpc.git && cd grpc && git branch && git submodule update --init && \
    mkdir -p cmake/build && cd cmake/build && cmake ../.. && make -j$(nproc) && make install


# 安装bear工具  Ubuntu22.04建议直接 apt install -y bear安装  centos7可以编译安装 bear-v2.4.4版本，此版本没有太多的依赖，相对而言比较好安装
# 参考连接  https://blog.airstone.me/centos-ccls    https://www.zhihu.com/question/353722203/answer/2564104885
# 编译安装
# RUN V=3.1.3 && git clone -b $V https://github.com/rizsotto/Bear.git && cd Bear && \
#     cmake -DENABLE_UNIT_TESTS=OFF -DENABLE_FUNC_TESTS=OFF . && \
#     make all && make install && cd /root && rm Bear* -rf
# 直接安装
RUN apt install -y bear


# 安装gcc工具 参考文档：https://blog.csdn.net/RBPicsdn/article/details/79565383
# 源码地址：https://mirrors.ustc.edu.cn/gnu/gcc/
# 如果执行./contrib/download_prerequisites脚本下载依赖失败，那么只能手动下载之后放在工程的根目录下面
# 依赖镜像源：https://mirrors.tuna.tsinghua.edu.cn/gnu/gmp/gmp-6.2.1.tar.bz2  或者  http://mirror.linux-ia64.org/gnu/gmp/gmp-6.2.1.tar.bz2
# 阿里镜像源：https://mirrors.aliyun.com/gnu/gmp/gmp-6.2.1.tar.bz2
# versions=(13.2.0 12.3.0 11.4.0 10.5.0 9.5.0 8.5.0 7.5.0 6.5.0 5.5.0 4.5.0)
# 使用此命令切换gcc的版本：update-alternatives --config gcc
# 编译GCC问题总结：https://www.cnblogs.com/nopnog/p/14990738.html
RUN versions='13 12 11 10 9' && update-alternatives --remove-all cpp && update-alternatives --remove-all c++ && \
    for V in ${versions}; do apt install -y gcc-$V g++-$V && \
    update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-$V $V \
                         --slave  /usr/bin/cpp cpp /usr/bin/cpp-$V \
                         --slave  /usr/bin/g++ g++ /usr/bin/g++-$V \
                         --slave  /usr/bin/c++ c++ /usr/bin/g++-$V \
                         --slave  /usr/bin/gcc-ar gcc-ar /usr/bin/gcc-ar-$V \
                         --slave  /usr/bin/gcc-nm gcc-nm /usr/bin/gcc-nm-$V \
                         --slave  /usr/bin/gcc-ranlib gcc-ranlib /usr/bin/gcc-ranlib-$V \
                         --slave  /usr/bin/gcov gcov /usr/bin/gcov-$V \
                         --slave  /usr/bin/gcov-dump gcov-dump /usr/bin/gcov-dump-$V \
                         --slave  /usr/bin/gcov-tool gcov-tool /usr/bin/gcov-tool-$V \
                         --slave  /usr/bin/lto-dump lto-dump /usr/bin/lto-dump-$V; done && \
    update-alternatives --display gcc

# 切换到gcc9，不然编译器版本太高，总是报很多警告
RUN echo 5 | update-alternatives --config gcc
RUN gcc -v

RUN V=8.5.0 && priority=8 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz && \
    cd gcc-$V && if [ -f ./contrib/download_prerequisites ]; then \
    sed -i 's@ftp://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@http://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@https://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority
RUN V=8.5.0 && update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V 8 \
                         --slave  /usr/bin/cpp cpp /usr/local/gcc-v$V/bin/cpp-$V \
                         --slave  /usr/bin/g++ g++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/c++ c++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/gcc-ar gcc-ar /usr/local/gcc-v$V/bin/gcc-ar-$V \
                         --slave  /usr/bin/gcc-nm gcc-nm /usr/local/gcc-v$V/bin/gcc-nm-$V \
                         --slave  /usr/bin/gcc-ranlib gcc-ranlib /usr/local/gcc-v$V/bin/gcc-ranlib-$V \
                         --slave  /usr/bin/gcov gcov /usr/local/gcc-v$V/bin/gcov-$V \
                         --slave  /usr/bin/gcov-dump gcov-dump /usr/local/gcc-v$V/bin/gcov-dump-$V \
                         --slave  /usr/bin/gcov-tool gcov-tool /usr/local/gcc-v$V/bin/gcov-tool-$V \
                         --slave  /usr/bin/lto-dump lto-dump /usr/local/gcc-v$V/bin/lto-dump-$V
# RUN echo 6 | update-alternatives --config gcc
RUN gcc -v

# 参考连接：https://blog.csdn.net/yangchunyuycy/article/details/120848335
RUN V=7.5.0 && priority=7 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz && \
    cd gcc-$V && if [ -f ./contrib/download_prerequisites ]; then \
    sed -i 's@ftp://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@http://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@https://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && sed -i '/target-libsanitizer \\/d' ../configure && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority
RUN V=7.5.0 && update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V 7 \
                         --slave  /usr/bin/cpp cpp /usr/local/gcc-v$V/bin/cpp-$V \
                         --slave  /usr/bin/g++ g++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/c++ c++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/gcc-ar gcc-ar /usr/local/gcc-v$V/bin/gcc-ar-$V \
                         --slave  /usr/bin/gcc-nm gcc-nm /usr/local/gcc-v$V/bin/gcc-nm-$V \
                         --slave  /usr/bin/gcc-ranlib gcc-ranlib /usr/local/gcc-v$V/bin/gcc-ranlib-$V \
                         --slave  /usr/bin/gcov gcov /usr/local/gcc-v$V/bin/gcov-$V \
                         --slave  /usr/bin/gcov-dump gcov-dump /usr/local/gcc-v$V/bin/gcov-dump-$V \
                         --slave  /usr/bin/gcov-tool gcov-tool /usr/local/gcc-v$V/bin/gcov-tool-$V \
                         --slave  /usr/bin/lto-dump lto-dump /usr/local/gcc-v$V/bin/lto-dump-$V

# 参考连接：https://blog.csdn.net/yangchunyuycy/article/details/120848335
RUN V=6.5.0 && priority=6 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz && \
    cd gcc-$V && if [ -f ./contrib/download_prerequisites ]; then \
    sed -i 's@ftp://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@http://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@https://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && sed -i '/target-libsanitizer \\/d' ../configure && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority
RUN V=6.5.0 && update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V 6 \
                         --slave  /usr/bin/cpp cpp /usr/local/gcc-v$V/bin/cpp-$V \
                         --slave  /usr/bin/g++ g++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/c++ c++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/gcc-ar gcc-ar /usr/local/gcc-v$V/bin/gcc-ar-$V \
                         --slave  /usr/bin/gcc-nm gcc-nm /usr/local/gcc-v$V/bin/gcc-nm-$V \
                         --slave  /usr/bin/gcc-ranlib gcc-ranlib /usr/local/gcc-v$V/bin/gcc-ranlib-$V \
                         --slave  /usr/bin/gcov gcov /usr/local/gcc-v$V/bin/gcov-$V \
                         --slave  /usr/bin/gcov-dump gcov-dump /usr/local/gcc-v$V/bin/gcov-dump-$V \
                         --slave  /usr/bin/gcov-tool gcov-tool /usr/local/gcc-v$V/bin/gcov-tool-$V \
                         --slave  /usr/bin/lto-dump lto-dump /usr/local/gcc-v$V/bin/lto-dump-$V
RUN echo 6 | update-alternatives --config gcc
RUN gcc -v

# 参考连接：https://juejin.cn/post/7087589940996882439
RUN V=5.5.0 && priority=5 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz && \
    cd gcc-$V && if [ -f ./contrib/download_prerequisites ]; then \
    sed -i 's@ftp://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@http://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@https://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && sed -i '/target-libsanitizer \\/d' ../configure && \
    # --build=x86_64-linux-gnu  加上这个参数有啥效果？
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority
RUN V=5.5.0 && update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V 5 \
                         --slave  /usr/bin/cpp cpp /usr/local/gcc-v$V/bin/cpp-$V \
                         --slave  /usr/bin/g++ g++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/c++ c++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/gcc-ar gcc-ar /usr/local/gcc-v$V/bin/gcc-ar-$V \
                         --slave  /usr/bin/gcc-nm gcc-nm /usr/local/gcc-v$V/bin/gcc-nm-$V \
                         --slave  /usr/bin/gcc-ranlib gcc-ranlib /usr/local/gcc-v$V/bin/gcc-ranlib-$V \
                         --slave  /usr/bin/gcov gcov /usr/local/gcc-v$V/bin/gcov-$V \
                         --slave  /usr/bin/gcov-dump gcov-dump /usr/local/gcc-v$V/bin/gcov-dump-$V \
                         --slave  /usr/bin/gcov-tool gcov-tool /usr/local/gcc-v$V/bin/gcov-tool-$V \
                         --slave  /usr/bin/lto-dump lto-dump /usr/local/gcc-v$V/bin/lto-dump-$V
RUN gcc -v
RUN gcc -v

RUN V=4.8.2 && priority=4 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz && \
    cd gcc-$V && if [ -f ./contrib/download_prerequisites ]; then \
    sed -i 's@ftp://gcc.gnu.org@https://gcc.gnu.org@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority
RUN V=4.8.2 && update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V 4 \
                         --slave  /usr/bin/cpp cpp /usr/local/gcc-v$V/bin/cpp-$V \
                         --slave  /usr/bin/g++ g++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/c++ c++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/gcc-ar gcc-ar /usr/local/gcc-v$V/bin/gcc-ar-$V \
                         --slave  /usr/bin/gcc-nm gcc-nm /usr/local/gcc-v$V/bin/gcc-nm-$V \
                         --slave  /usr/bin/gcc-ranlib gcc-ranlib /usr/local/gcc-v$V/bin/gcc-ranlib-$V \
                         --slave  /usr/bin/gcov gcov /usr/local/gcc-v$V/bin/gcov-$V \
                         --slave  /usr/bin/gcov-dump gcov-dump /usr/local/gcc-v$V/bin/gcov-dump-$V \
                         --slave  /usr/bin/gcov-tool gcov-tool /usr/local/gcc-v$V/bin/gcov-tool-$V \
                         --slave  /usr/bin/lto-dump lto-dump /usr/local/gcc-v$V/bin/lto-dump-$V
RUN AA

RUN V=4.5.0 && priority=4 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz --build=x86_64-linux-gnu && \
    cd gcc-$V && cat ./contrib/download_prerequisites && if [ -f ./contrib/download_prerequisites ]; then \
    sed -i 's@ftp://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@http://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@https://gcc.gnu.org/pub@http://mirror.linux-ia64.org/gnu@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority
RUN V=4.5.0 && update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V 4 \
                         --slave  /usr/bin/cpp cpp /usr/local/gcc-v$V/bin/cpp-$V \
                         --slave  /usr/bin/g++ g++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/c++ c++ /usr/local/gcc-v$V/bin/g++-$V \
                         --slave  /usr/bin/gcc-ar gcc-ar /usr/local/gcc-v$V/bin/gcc-ar-$V \
                         --slave  /usr/bin/gcc-nm gcc-nm /usr/local/gcc-v$V/bin/gcc-nm-$V \
                         --slave  /usr/bin/gcc-ranlib gcc-ranlib /usr/local/gcc-v$V/bin/gcc-ranlib-$V \
                         --slave  /usr/bin/gcov gcov /usr/local/gcc-v$V/bin/gcov-$V \
                         --slave  /usr/bin/gcov-dump gcov-dump /usr/local/gcc-v$V/bin/gcov-dump-$V \
                         --slave  /usr/bin/gcov-tool gcov-tool /usr/local/gcc-v$V/bin/gcov-tool-$V \
                         --slave  /usr/bin/lto-dump lto-dump /usr/local/gcc-v$V/bin/lto-dump-$V
RUN AA

# 安装gcc g++工具

# TODO 多版本安装软件
# gcc g++ gdb

# TODO 清理docker build过程中产生的缓存文件


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

# docker构建时提示错误：output clipped, log limit 1MiB reached  解决方案如下，分别执行两条命令即可
# docker buildx create --use --name larger_log --driver-opt env.BUILDKIT_STEP_LOG_MAX_SIZE=50000000
# docker buildx build --progress plain .

# 测试：    docker run -it --rm --privileged --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private  ubuntu-dev:22.04

# 实际使用：docker run d --name wangmin --privileged --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -v /var/run/docker.sock:/var/run/docker.sock -p 2800:22 -p 2800-2820:8000:8020 ubuntu-dev:22.04


# 使用systemd启动容器参考连接：https://blog.csdn.net/kencaber/article/details/121980242
# 参考连接：https://blog.csdn.net/m0_37886429/article/details/80350659?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&utm_relevant_index=1
CMD [ "/lib/systemd/systemd", "log-level=info", "unit=sysinit.target" ]