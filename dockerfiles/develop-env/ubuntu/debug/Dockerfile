FROM ubuntu:23.10

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

# 修改apt源为清华源
RUN cat <<EOF >> /etc/apt/sources.list
# 默认注释了源码镜像以提高 apt update 速度，如有需要可自行取消注释
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic-updates main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic-backports main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic-backports main restricted universe multiverse

# deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic-security main restricted universe multiverse
# # deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic-security main restricted universe multiverse

deb http://security.ubuntu.com/ubuntu/ mantic-security main restricted universe multiverse
# deb-src http://security.ubuntu.com/ubuntu/ mantic-security main restricted universe multiverse

# 预发布软件源，不建议启用
# deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic-proposed main restricted universe multiverse
# # deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ mantic-proposed main restricted universe multiverse
EOF

RUN apt update -y && apt upgrade -y && \
    DEBIAN_FRONTEND=noninteractive TZ=Asia/Shanghai apt -y install tzdata && \
    apt install -y man-db manpages-posix && apt install -y manpages-dev manpages-posix-dev && yes | unminimize && \
    apt install -y net-tools telnet sysstat bridge-utils bash-completion vim jq tar openssl iputils-ping lsof lvm2 \
    dnsutils curl gcc g++ automake autoconf make tree stress htop atop sysbench ipvsadm ipset conntrack ufw git \
    build-essential flex libncurses-dev bison libelf-dev libssl-dev bc openssh-server ansible unzip binutils \
    libc-ares-dev libtool pkg-config libsystemd-dev file texinfo nmap zsh autojump byobu language-pack-en \
    language-pack-zh-hans python3-pygments chroma rsync

# 设置语言为中文
ENV LANG en_US.utf8
ENV LC_TIME C.utf8

#  同步时间
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# github加速（非常重要，否则github下载代码将会相当之慢）
RUN git config --global url."https://hub.fgit.cf".insteadof "https://github.com"
# ssh方式加速 目前似乎没有ssh方式加速的方式，除非自己搭VPN

# shell命令行颜色提示; git分支名自动补全
# 设置shell提示符颜色 参考链接：https://www.jianshu.com/p/a1d698d1f7c8
# 使用zsh以及ohmyzsh来替代这里的配置，如果使用bash，可以打开这里的配置
# RUN V=$(git -v|awk '{print $3}') && git clone -b v${V} https://github.com/git/git.git && \
#     cp git/contrib/completion/git-completion.bash ~/.git-completion.bash && echo '\nsource ~/.git-completion.bash' >> ~/.bashrc  && \
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
    echo '\nZSH_AUTOSUGGEST_HIGHLIGHT_STYLE="fg=#ff00ff,bg=cyan,bold,underline"' >> ~/.zshrc && \
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


# 切换脚本解释器为bash, Ubuntu默认的脚本解释器为dash
RUN chsh -s $(which zsh) && ln -fs /bin/zsh /bin/sh






# debug命令
# docker rm -f debug && docker build -t ubuntu-debug:v1.0 --progress=plain --add-host=dl.google.com:220.181.174.225 . && docker run -d --rm --privileged --name debug --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -p 21800:22 -v /var/run/docker.sock:/var/run/docker.sock -v /home/wangmin/workspace:/root/workspace:rw ubuntu-debug:v1.0 && docker exec -it debug zsh


# 使用systemd启动容器参考连接：https://blog.csdn.net/kencaber/article/details/121980242
# 参考连接：https://blog.csdn.net/m0_37886429/article/details/80350659?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&utm_relevant_index=1
CMD [ "/lib/systemd/systemd", "log-level=info", "unit=sysinit.target" ]