FROM ubuntu:23.10

# 构建命令：
# 以docker compose的方式启动容器命令

# TODO List
# TODO 集成Socks代理


USER root
WORKDIR /root
ENV container docker
STOPSIGNAL SIGRTMIN+3
VOLUME [ "/tmp", "/run", "/run/lock" ]

# 设置root密码，方便使用ssh进行登录
RUN echo root:123456 | chpasswd

# 移除无用的服务
RUN rm -f /lib/systemd/system/multi-user.target.wants/* \
  /etc/systemd/system/*.wants/* \
  /lib/systemd/system/local-fs.target.wants/* \
  /lib/systemd/system/sockets.target.wants/*udev* \
  /lib/systemd/system/sockets.target.wants/*initctl* \
  /lib/systemd/system/sysinit.target.wants/systemd-tmpfiles-setup* \
  /lib/systemd/system/systemd-update-utmp*

# 安装证书
RUN apt update -y && apt install -y --no-install-recommends ca-certificates apt-transport-https

# 修改apt源为清华源
RUN cat <<EOF > /etc/apt/sources.list
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ lunar main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ lunar-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ lunar-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ lunar-security main restricted universe multiverse

EOF

RUN apt update -y && apt upgrade -y && \
    DEBIAN_FRONTEND=noninteractive TZ=Asia/Shanghai apt -y install tzdata && \
    apt install -y man-db manpages-posix && apt install -y manpages-dev manpages-posix-dev && yes | unminimize && \
    apt install -y net-tools telnet sysstat bridge-utils bash-completion vim jq tar openssl iputils-ping lsof lvm2 \
    dnsutils curl gcc g++ automake autoconf make tree stress htop atop sysbench ipvsadm ipset conntrack ufw git \
    build-essential flex libncurses-dev bison libelf-dev libssl-dev bc openssh-server ansible unzip binutils \
    libc-ares-dev libtool pkg-config libsystemd-dev file texinfo nmap zsh autojump byobu language-pack-en \
    language-pack-zh-hans python3-pygments chroma rsync libseccomp-dev

# 先切换为bash，shell有些语法得用bash来解释。最后在切换为zsh
RUN chsh -s /bin/bash && ln -fs /bin/bash /bin/sh

# 设置语言为中文
ENV LANG en_US.utf8
ENV LC_TIME C.utf8

# 开启git日志，git clone时可以清晰的知道使用的是哪个地址，以及下载这个地址是否正常
ENV GIT_TRACE=2
# 开启更详细的git日志
# ENV GIT_CURL_VERBOSE=1

#  同步时间
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# github加速（非常重要，否则github下载代码将会相当之慢）
# RUN git config --global url."https://hub.fgit.cf".insteadof "https://github.com"
# 下面三种方式直接配置似乎有问题，git insteadOf工作好像不正常
# RUN git config --global url."https://gitclone.com/github.com".insteadof "https://github.com"
# RUN git config --global url."https://github.moeyy.xyz/https://github.com".insteadof "https://github.com"
RUN git config --global url."https://ghproxy.com/https://github.com".insteadof "https://github.com"
# ssh方式加速 目前似乎没有ssh方式加速的方式，除非自己搭VPN

# shell命令行颜色提示; git分支名自动补全
# 设置shell提示符颜色 参考链接：https://www.jianshu.com/p/a1d698d1f7c8
# 使用zsh以及ohmyzsh来替代这里的配置，如果使用bash，可以打开这里的配置
# RUN V=$(git -v|awk '{print $3}') && git clone -b v${V} https://github.com/git/git.git && \
#     cp git/contrib/completion/git-completion.bash ~/.git-completion.bash && echo -e '\nsource ~/.git-completion.bash' >> ~/.bashrc  && \
#     cp git/contrib/completion/git-prompt.sh ~/.git-prompt.sh && rm -rf git && \
#     cat <<EOF >> ~/.bashrc

# GIT_PS1_SHOWDIRTYSTATE=true
# GIT_PS1_SHOWCOLORHINTS=true
# GIT_PS1_SHOWSTASHSTATE=true
# GIT_PS1_SHOWUNTRACKEDFILES=true
# #GIT_PS1_SHOWUPSTREAM=auto
# if [ -f ~/.git-completion.bash ]; then
#   source ~/.git-prompt.sh
#   PROMPT_COMMAND='__git_ps1 "\[\e[37m\][\[\e[32m\]\u\[\e[37m\]@\[\e[35m\]\h\[\e[0m\] \[\e[36m\]\w\[\e[0m\]]" "\\\$ "'
# fi
# EOF

# zsh配置参考链接：https://github.com/skywind3000/vim/blob/master/etc/zshrc.zsh
# 安装ohmyzsh（使用raw.gitmirror.com加速raw.githubusercontent.com）
RUN echo Y | sh -c "$(curl -fsSL https://raw.gitmirror.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" && \
    # 安装powerlevel10k主题
    # git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k && \
    # 安装历史命令提示插件 参考地址：https://github.com/zsh-users/zsh-autosuggestions
    git clone --depth=1 https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions && \
    # 安装命令行语法高亮插件 参考地址：https://github.com/zsh-users/zsh-syntax-highlighting.git
    git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    # 下载自动补全 参考链接：https://github.com/zsh-users/zsh-completions
    git clone https://github.com/zsh-users/zsh-completions ${ZSH_CUSTOM:-${ZSH:-~/.oh-my-zsh}/custom}/plugins/zsh-completions && \
    # 设置默认主题为powerlevel10k
    # sed -i 's@^ZSH_THEME.*@ZSH_THEME="powerlevel10k/powerlevel10k"@g' ~/.zshrc && \
    sed -i 's@^ZSH_THEME.*@ZSH_THEME="dst"@g' ~/.zshrc && \
    # 开启错误命令自动更正 （这个特性不咋好用）
    # sed -i 's@# ENABLE_CORRECTION.*@ENABLE_CORRECTION="true"@g' ~/.zshrc && \
    # 在命令执行过程中，使用小红点进行提示
    sed -i 's@# COMPLETION_WAITING_DOTS.*@COMPLETION_WAITING_DOTS="true"@g' ~/.zshrc && \
    # 设置ohmyzsh自动更新频率为每30天更新
    sed -i "s@# zstyle ':omz:update' frequency.*@zstyle ':omz:update' frequency 30@g" ~/.zshrc && \
    # 设置中划线、下划线不敏感
    sed -i 's@# HYPHEN_INSENSITIVE="true"@HYPHEN_INSENSITIVE="true"@g' ~/.zshrc && \
    # 历史命令日期的显示格式
    sed -i 's@# HIST_STAMPS="mm/dd/yyyy"@HIST_STAMPS="yyyy/mm/dd"@g' ~/.zshrc && \
    # 启用安装的插件
    sed -i 's@^plugins.*@plugins=(zsh-autosuggestions zsh-syntax-highlighting kubectl kube-ps1 kubectx docker history z colorize colored-man-pages)@g'  ~/.zshrc && \
    # 开启zsh自动补全
    num=$(awk '/^source/{print NR}' .zshrc|head -n 1) && sed -i "${num}i fpath+=${ZSH_CUSTOM:-${ZSH:-~/.oh-my-zsh}/custom}/plugins/zsh-completions/src" ~/.zshrc && \
    # 开启颜色高亮提示
    echo -e '\nZSH_AUTOSUGGEST_HIGHLIGHT_STYLE="fg=#ff00ff,bg=cyan,bold,underline"' >> ~/.zshrc && \
    echo 'ZSH_AUTOSUGGEST_STRATEGY=(history completion)' >> ~/.zshrc && \
    echo 'ZSH_COLORIZE_STYLE="colorful"' >> ~/.zshrc && \
    echo 'ZSH_COLORIZE_CHROMA_FORMATTER=terminal256' >> ~/.zshrc


# 开启SSHD服务, 错误的构建方式，镜像在构建过程中不应该启动sshd服务，应该在镜像的启动点启动sshd服务 (当然，如果使用Systemd作为一号进程，那么ssh服务自然会被启动)
RUN sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config && \
    sed -i 's/#AddressFamily any/AddressFamily any/g' /etc/ssh/sshd_config && \
    sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config && \
    sed -i 's/#PermitRootLogin.*/PermitRootLogin yes/g' /etc/ssh/sshd_config && \
    sed -i "s/UsePAM.*/UsePAM no/g" /etc/ssh/sshd_config && \
    systemctl enable ssh.service

# golang环境 https://go.dev/doc/install  多版本管理 https://blog.csdn.net/weixin_41910261/article/details/112969035
# 如果dl.google.com解析不了，需要添加dns解析，那么可以使用如下命令：docker build --add-host=dl.google.com:58.63.233.97  --progress=plain .
# 使用此命令切换go的版本：update-alternatives --config go
ENV GOLANGVERSION=1.24.5
RUN curl -SL https://mirrors.aliyun.com/golang/go${GOLANGVERSION}.linux-amd64.tar.gz -o go${GOLANGVERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go${GOLANGVERSION}.linux-amd64.tar.gz && \
    rm -f go${GOLANGVERSION}.linux-amd64.tar.gz && \
    echo -e '\nexport GOPATH=/root/go' | tee -a ~/.bashrc ~/.zshrc && \
    echo -e 'PATH=$PATH:$GOPATH/bin' | tee -a ~/.bashrc ~/.zshrc && \
    /usr/local/go/bin/go env -w GO111MODULE=on && \
    /usr/local/go/bin/go env -w GOPROXY=https://goproxy.cn,direct && \
    # 安装不同的go版本，并使用update-alternatives工具进行切换
    versions='1.13.15 1.14.15 1.15.15 1.16.15 1.17.13 1.18.10 1.19.13 1.20.8 1.21.0 1.22.12 1.23.11' && \
    priority=1 && for V in ${versions}; do \
    /usr/local/go/bin/go install golang.org/dl/go$V@latest && ~/go/bin/go$V download && \
    update-alternatives  --install /usr/bin/go go ~/sdk/go$V/bin/go $priority && priority=$(($priority+1)); done && \
    update-alternatives  --install /usr/bin/go go /usr/local/go/bin/go $priority && \
    update-alternatives --display go && \
    # 私有仓库
    /usr/local/go/bin/go env -w GOPRIVATE=gitcdteam.skyguardmis.com && \
    /usr/local/go/bin/go env -w GOINSECURE=gitcdteam.skyguardmis.com && \
    # 设置编译出来的二进制文件的根目录，否则默认的/tmp目录可能没有可执行权限
    /usr/local/go/bin/go env -w GOTMPDIR=/root && \
    # 本地k8s测试工具
    /usr/local/go/bin/go install sigs.k8s.io/kind@latest && \
    # 漏洞检测工具
    /usr/local/go/bin/go install golang.org/x/vuln/cmd/govulncheck@latest && \
    # protoc编译器
    /usr/local/go/bin/go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    # grpc
    /usr/local/go/bin/go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    # gateway
    /usr/local/go/bin/go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
    # openapi
    /usr/local/go/bin/go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && \
    # 安装delve，用于debug go代码
    /usr/local/go/bin/go install github.com/go-delve/delve/cmd/dlv@latest && \
    # 安装wire自动注入工具
    /usr/local/go/bin/go install github.com/google/wire/cmd/wire@latest && \
    # boltdb查看工具
    /usr/local/go/bin/go install github.com/br0xen/boltbrowser@latest && \
    /usr/local/go/bin/go install github.com/cweill/gotests/gotests@latest && \
    /usr/local/go/bin/go install github.com/fatih/gomodifytags@latest && \
    /usr/local/go/bin/go install github.com/josharian/impl@latest && \
    /usr/local/go/bin/go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    /usr/local/go/bin/go install honnef.co/go/tools/cmd/staticcheck@latest && \
    /usr/local/go/bin/go install golang.org/x/tools/gopls@latest && \
    /usr/local/go/bin/go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
    /usr/local/go/bin/go install mvdan.cc/gofumpt@latest && \
    /usr/local/go/bin/go install github.com/kardianos/govendor@latest


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
RUN versions='v13.14.0 v14.21.3 v15.14.0 v16.20.2 v17.9.1 v18.8.0 v19.9.0 v20.7.0 v21.7.3 v22.17.1' && \
    priority=1 && for V in ${versions}; do \
    curl -SL https://mirrors.huaweicloud.com/nodejs/$V/node-$V-linux-x64.tar.xz -o node-$V-linux-x64.tar.xz && \
    rm -rf /usr/local/node-$V && mkdir /usr/local/node-$V && tar -C /usr/local/node-$V -xJf node-$V-linux-x64.tar.xz --strip-components 1 && \
    update-alternatives --install /usr/bin/node nodejs /usr/local/node-$V/bin/node $priority \
                        --slave  /usr/bin/npm npm /usr/local/node-$V/bin/npm && priority=$(($priority+1)); done && \
    update-alternatives --display nodejs && rm -f /root/node* && \
    # npm国内源加速配置
    npm config set registry https://registry.npm.taobao.org
    # TODO 配置nodejs的环境变量

# 安装Kubectl客户端
# 使用此命令切换kubectl的版本：update-alternatives --config kubectl
RUN versions='v1.18.20 v1.19.16 v1.20.15 v1.21.14 v1.22.17 v1.23.17 v1.24.17 v1.25.14 v1.26.9 v1.27.6 v1.28.2 1.33.3' && \
    priority=1 && for V in ${versions}; do \
    mkdir /usr/local/k8s-$V && curl -SL https://files.m.daocloud.io/storage.googleapis.com/kubernetes-release/release/$V/bin/linux/amd64/kubectl \
    -o /usr/local/k8s-$V/kubectl && chmod +x /usr/local/k8s-$V/kubectl && \
    update-alternatives --install /usr/bin/kubectl kubectl /usr/local/k8s-$V/kubectl $priority && priority=$(($priority+1)); done && \
    update-alternatives --display kubectl && echo -e "\nsource <(kubectl completion zsh)" >> ~/.zshrc

# 安装Kustomize
RUN versions='3.5.1 3.8.5 4.2.0 4.4.0 4.5.7 5.1.1 5.7.0' && \
    priority=1 && for V in ${versions}; do mkdir -p /usr/local/kustomize-v${V}  && \
    curl -SLO https://files.m.daocloud.io/github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v${V}/kustomize_v${V}_linux_amd64.tar.gz && \
    tar -zxf kustomize_v${V}_linux_amd64.tar.gz && mv kustomize /usr/local/kustomize-v${V} && rm -f kustomize_v${V}_linux_amd64.tar.gz && \
    update-alternatives --install /usr/bin/kustomize kustomize /usr/local/kustomize-v${V}/kustomize $priority && priority=$(($priority+1)); done && \
    update-alternatives --display kustomize

# 安装Docker客户端
# 使用此命令切换docker的版本：update-alternatives --config docker
RUN versions='18.09.9 19.03.15 20.10.24 23.0.6 24.0.6 26.1.4 27.5.1 28.3.2' && \
    priority=1 && for V in ${versions}; do \
    curl -SL https://download.docker.com/linux/static/stable/x86_64/docker-$V.tgz -o docker-$V.tgz && \
    mkdir /usr/local/docker-$V && tar -xzf docker-$V.tgz --strip 1 -C /usr/local/docker-$V docker/docker && \
    update-alternatives --install /usr/bin/docker docker /usr/local/docker-$V/docker $priority && priority=$(($priority+1)); done && \
    update-alternatives --display docker && \
    rm -rf /usr/libexec/docker/cli-plugins && mkdir -p /usr/libexec/docker/cli-plugins && \
    curl -SL https://files.m.daocloud.io/github.com/docker/buildx/releases/download/v0.11.2/buildx-v0.11.2.linux-amd64 -o /usr/libexec/docker/cli-plugins/docker-buildx && \
    chmod +x /usr/libexec/docker/cli-plugins/docker-buildx && \
    curl -SL https://files.m.daocloud.io/github.com/docker/compose/releases/download/v2.21.0/docker-compose-linux-x86_64 -o /usr/libexec/docker/cli-plugins/docker-compose && \
    chmod +x /usr/libexec/docker/cli-plugins/docker-compose && \
    rm -f /root/docker*
    # docker 命令补全

# 安装Python环境
# 使用此命令切换python的版本：update-alternatives --config python
RUN versions='2.7 3.10 3.11' && apt install -y python3-requests python3-venv && \
    priority=1 && for V in ${versions}; do apt install -y python$V && \
    update-alternatives --install /usr/bin/python python /usr/bin/python$V $priority && priority=$(($priority+1)); done && \
    update-alternatives --display python && \
    apt install -y python3-pip && pip --version && pip --version && \
    # pip镜像加速设置
    pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/ && \
    pip config set global.timeout 6000

# 安装Helm工具
ENV HELMVERSION=v3.18.4
RUN curl -SL https://files.m.daocloud.io/get.helm.sh/helm-${HELMVERSION}-linux-amd64.tar.gz -o helm-${HELMVERSION}-linux-amd64.tar.gz && \
    rm -f /usr/local/bin/helm && tar -zxf helm-${HELMVERSION}-linux-amd64.tar.gz --strip 1 -C /usr/local/bin linux-amd64/helm && \
    helm version && rm -rf helm*

# 安装Kubebuilder工具
# 使用此命令切换kubebuilder的版本：update-alternatives --config kubebuilder
RUN versions='v3.2.0 v3.4.0 v3.6.0 v3.9.1 v3.10.0 v3.11.1 v3.12.0 v4.7.0' && \
    priority=1 && mkdir /usr/local/kubebuilder && for V in ${versions}; do \
    curl -SL https://files.m.daocloud.io/github.com/kubernetes-sigs/kubebuilder/releases/download/$V/kubebuilder_linux_amd64 \
    -o /usr/local/kubebuilder/kubebuilder-$V && chmod +x /usr/local/kubebuilder/kubebuilder-$V && \
    update-alternatives --install /usr/bin/kubebuilder kubebuilder /usr/local/kubebuilder/kubebuilder-$V $priority && priority=$(($priority+1)); done && \
    update-alternatives --display kubebuilder

# 安装Protoc工具
# 使用此命令切换protoc的版本：update-alternatives --config protoc
RUN versions='21.9 22.0 22.5 23.0 23.4 24.0 24.3 31.1' && \
    priority=1 && for V in ${versions}; do mkdir /usr/local/protoc-v$V && \
    curl -SL https://files.m.daocloud.io/github.com/protocolbuffers/protobuf/releases/download/v$V/protoc-$V-linux-x86_64.zip -o protoc-$V-linux-x86_64.zip && \
    unzip -d /usr/local/protoc-v$V protoc-$V-linux-x86_64.zip && rm -f /root/protoc* && \
    update-alternatives --install /usr/bin/protoc protoc /usr/local/protoc-v$V/bin/protoc $priority && priority=$(($priority+1)); done && \
    update-alternatives --display protoc


# 安装cmake工具
# 使用此命令切换cmake的版本：update-alternatives --config cmake
RUN versions='3.12.4 3.19.8 3.20.6 3.25.3 3.26.5 3.30.9 4.0.3' && \
    priority=1 && for V in ${versions}; do mkdir /usr/local/cmake-v$V && \
    curl -SL https://files.m.daocloud.io/github.com/Kitware/CMake/releases/download/v$V/cmake-$V-linux-x86_64.tar.gz -o cmake-$V-linux-x86_64.tar.gz && \
    tar -C /usr/local/cmake-v$V -xzf cmake-$V-linux-x86_64.tar.gz  --strip-components 1 && rm -f /root/cmake* && \
    update-alternatives --install /usr/bin/cmake cmake /usr/local/cmake-v$V/bin/cmake $priority && priority=$(($priority+1)); done && \
    update-alternatives --display cmake

# 安装C++, Python, Ruby, Objective-C, PHP, C#语言的grpc插件
# 参考连接：https://blog.csdn.net/Aidam_Bo/article/details/117568034
RUN git clone -b v1.58.1 https://github.com/grpc/grpc.git && cd grpc && git submodule update --init && \
    mkdir -p cmake/build && cd cmake/build && cmake ../.. && make -j$(nproc) && make install && cd /root \
    # 删除编译grpc自动安装的低版本protoc
    && rm -rf grpc && rm -f /usr/local/bin/protoc*

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
RUN versions='15 14 13 12 11 10 9' && update-alternatives --remove-all cpp && update-alternatives --remove-all c++ && \
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

RUN V=8.5.0 && priority=8 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz && \
    cd gcc-$V && if [ -f ./contrib/download_prerequisites ]; then \
    sed -i 's@ftp://gcc.gnu.org@https://gcc.gnu.org@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority \
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
RUN V=7.5.0 && priority=7 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz && \
    cd gcc-$V && if [ -f ./contrib/download_prerequisites ]; then \
    sed -i 's@ftp://gcc.gnu.org@https://gcc.gnu.org@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu --disable-libsanitizer && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority \
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
    sed -i 's@ftp://gcc.gnu.org@https://gcc.gnu.org@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu --disable-libsanitizer && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority \
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

# 切换到gcc9，不然编译器版本太高，总是报很多警告
RUN echo 5 | update-alternatives --config gcc

# 参考连接：https://juejin.cn/post/7087589940996882439
RUN V=5.5.0 && priority=5 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz && \
    cd gcc-$V && if [ -f ./contrib/download_prerequisites ]; then \
    sed -i 's@ftp://gcc.gnu.org@https://gcc.gnu.org@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites; fi && \
    mkdir build-gcc-$V && cd build-gcc-$V && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu --disable-libsanitizer && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority \
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

# 问题解决方案：https://unix.stackexchange.com/questions/335717/how-to-handle-error-compiling-gcc-4-7-0-using-gcc-6-2-1
# https://stackoverflow.com/questions/41204632/unable-to-build-gcc-due-to-c11-errors
# https://blog.csdn.net/qq_31347869/article/details/120999082
# configure增加 --disable-libsanitizer参数
RUN V=4.8.2 && priority=4 && mkdir /usr/local/gcc-v$V && \
    curl -SL https://mirrors.ustc.edu.cn/gnu/gcc/gcc-$V/gcc-$V.tar.gz -o gcc-$V.tar.gz && tar -zxf gcc-$V.tar.gz && \
    cd gcc-$V && sed -i 's@ftp://gcc.gnu.org@https://gcc.gnu.org@g' ./contrib/download_prerequisites && \
    sed -i 's@\${fetch} --no-verbose@curl -SLOk @g' ./contrib/download_prerequisites && ./contrib/download_prerequisites && \
    sed -i 's@ANSI-C@C++@g' gcc/cp/Make-lang.in && sed -i '19,26d' gcc/cp/cfns.gperf && sed -i '1i %language=C++\n%define class-name libc_name' gcc/cp/cfns.gperf && \
    sed -i 's@ANSI-C code@C++ code@g' gcc/cp/cfns.h && sed -i 's@3.0.3@3.0.4@g' gcc/cp/cfns.h && sed -i 's@ANSI-C cfns.gperf@C++ --output-file cfns.h cfns.gperf@g' gcc/cp/cfns.h && \
    sed -i 's@line 1 "cfns.gperf"@line 3 "cfns.gperf"@g' gcc/cp/cfns.h && sed -i '50,57d;60,68d;125,130d' gcc/cp/cfns.h && \
    sed -i '52i class libc_name\n{\nprivate:\n  static inline unsigned int hash (const char *str, unsigned int len);\npublic:\n  static const char *libc_name_p (const char *str, unsigned int len);\n};\n\ninline unsigned int\nlibc_name::hash (register const char *str, register unsigned int len)' gcc/cp/cfns.h && \
    sed -i 's@^libc_name_p@libc_name::libc_name_p@g' gcc/cp/cfns.h && sed -i 's@libc_name_p@libc_name::libc_name_p@g' gcc/cp/except.c && \
    sed -i 's@ucontext \*uc@ucontext_t \*uc@g' libgcc/config/i386/linux-unwind.h && \
    mkdir build-gcc-$V && cd build-gcc-$V && \
    ../configure --prefix=/usr/local/gcc-v$V --enable-checking=release --enable-languages=c,c++ --disable-multilib --program-suffix=-$V --build=x86_64-linux-gnu --disable-libsanitizer && \
    make -j$(nproc) && make install && cd /root && rm -rf /root/gcc* && \
    update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-v$V/bin/gcc-$V $priority \
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

# 切换到最新的gcc
RUN echo 4 | update-alternatives --config gcc

# 打开文件监听限制，否则当文件数量过多，linux会限制打开文件数量
RUN touch /etc/sysctl.conf && \
    echo "fs.inotify.max_user_watches = 9994288" > /etc/sysctl.conf

# vscode 不然vscode安装好之后，老是找不到命令
RUN cat <<EOF >> ~/.zshrc

# 禁用nomatch报错
unsetopt nomatch
# 添加vscode可执行文件环境变量配置，否则有可能找不到vscode可执行文件
if [ -d ~/.vscode-server ];then
    code_latest_version=\$(ls -tral -1 --ignore=.* ~/.vscode-server/bin | sed -n '2p' | rev | cut -d' ' -f1 | rev)
    export PATH=\${HOME}/.vscode-server/bin/\${code_latest_version}/bin/remote-cli:\$PATH
fi
EOF

# 切换脚本解释器为zsh, Ubuntu默认的脚本解释器为dash
RUN chsh -s $(which zsh) && ln -fs /bin/zsh /bin/sh

# 关闭git日志，否则每次进入git clone的文件夹会打印日志信息
RUN unset GIT_TRACE
# 开启更详细的git日志
# RUN unset GIT_CURL_VERBOSE


# TODO 1、分屏问题（mobaxterm无法水平分屏，vscode可以）    
#  2、自动补全感觉有时候不大好用   3、vi编辑文件有时候无法退出（ctrl + ; 被占用）  4、ctrl + ;会自动插入
#  5、tab补全问题








# 构建镜像命令
# docker build -t ubuntu-dev:v1.0 --progress=plain .


# 调试dockerfile命令
# docker rm -f dev-debug && docker build -t ubuntu-dev:v1.0 --progress=plain --add-host=dl.google.com:220.181.174.225 . && docker run -d --rm --privileged --name dev-debug --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -p 21800:22 -v /var/run/docker.sock:/var/run/docker.sock -v /home/wangmin/workspace:/root/workspace:rw ubuntu-dev:v1.0 && docker exec -it dev-debug zsh

# docker rm -f dev-debug && docker run -d --rm --privileged --name dev-debug --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -p 21800:22 -v /var/run/docker.sock:/var/run/docker.sock -v /home/wangmin/workspace:/root/workspace:rw ubuntu-dev:v1.0 && docker exec -it dev-debug zsh

# 实际使用：docker run -d --name wangmin-dev --privileged --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -v /var/run/docker.sock:/var/run/docker.sock -v /home/wangmin/workspace:/root/workspace:rw -p 2800:22 -p 2810-2820:8010-8020 ubuntu-dev:v1.0


# 使用systemd启动容器参考连接：https://blog.csdn.net/kencaber/article/details/121980242
# 参考连接：https://blog.csdn.net/m0_37886429/article/details/80350659?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&utm_relevant_index=1
CMD [ "/lib/systemd/systemd", "log-level=info", "unit=sysinit.target" ]