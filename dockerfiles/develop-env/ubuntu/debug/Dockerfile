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

RUN apt update -y && apt upgrade -y
RUN DEBIAN_FRONTEND=noninteractive TZ=Asia/Shanghai apt -y install tzdata
RUN apt install -y net-tools telnet sysstat bridge-utils bash-completion vim jq tar openssl iputils-ping lsof git  openssh-server curl \
    zsh autojump

# 切换脚本解释器为bash, Ubuntu默认的脚本解释器为dash
RUN chsh -s $(which zsh) && ln -fs /bin/zsh /bin/sh

# github加速（非常重要，否则github下载代码将会相当之慢）
RUN git config --global url."https://hub.fgit.cf".insteadof "https://github.com"

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime


# 开启SSHD服务, 错误的构建方式，镜像在构建过程中不应该启动sshd服务，应该在镜像的启动点启动sshd服务 (当然，如果使用Systemd作为一号进程，那么ssh服务自然会被启动)
RUN sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config && \
    sed -i 's/#AddressFamily any/AddressFamily any/g' /etc/ssh/sshd_config && \
    sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config && \
    sed -i 's/#PermitRootLogin.*/PermitRootLogin yes/g' /etc/ssh/sshd_config && \
    sed -i "s/UsePAM.*/UsePAM no/g" /etc/ssh/sshd_config && \
    systemctl enable ssh.service


# vscode 不然vscode安装好之后，老是找不到命令
RUN cat <<EOF >> ~/.bashrc

# 添加vscode可执行文件环境变量配置，否则有可能找不到vscode可执行文件
if [ -d ~/.vscode-server ];then
    code_latest_version=\$(ls -tral -1 --ignore=.* ~/.vscode-server/bin | sed -n '2p' | rev | cut -d' ' -f1 | rev)
    export PATH=\${HOME}/.vscode-server/bin/\${code_latest_version}/bin/remote-cli:\$PATH
fi
EOF

# zsh配置参考链接：https://github.com/skywind3000/vim/blob/master/etc/zshrc.zsh
# 安装ohmyzsh（使用raw.gitmirror.com加速raw.githubusercontent.com）
RUN echo Y | sh -c "$(curl -fsSL https://raw.gitmirror.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" && \
    # 安装powerlevel10k主题
    # git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k && \
    # 安装历史命令提示插件
    git clone --depth=1 https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions && \
    # 安装命令行语法高亮插件
    git clone --depth=1 https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    # 下载字体
    # curl -SLO https://raw.gitmirror.com/powerline/powerline/develop/font/10-powerline-symbols.conf && \
    # curl -SLO https://raw.gitmirror.com/powerline/powerline/develop/font/PowerlineSymbols.otf && \
    # sudo mkdir /usr/share/fonts/OTF && \
    # sudo cp 10-powerline-symbols.conf /usr/share/fonts/OTF/  && \
    # sudo mv 10-powerline-symbols.conf /etc/fonts/conf.d/ && \
    # sudo mv PowerlineSymbols.otf /usr/share/fonts/OTF/ && \
    # fc-cache -vf  /usr/share/fonts/OTF/   && \
    # 设置默认主题为powerlevel10k
    # sed -i 's@^ZSH_THEME.*@ZSH_THEME="powerlevel10k/powerlevel10k"@g' ~/.zshrc && \
    sed -i 's@^ZSH_THEME.*@ZSH_THEME="dst"@g' ~/.zshrc && \
    # 开启错误命令自动更正
    sed -i 's@#ENABLE_CORRECTION.*@ENABLE_CORRECTION="true"@g' ~/.zshrc && \
    # 在命令执行过程中，使用小红点进行提示
    sed -i 's@#COMPLETION_WAITING_DOTS.*@COMPLETION_WAITING_DOTS="true"@g' ~/.zshrc && \
    # 启用安装的插件
    sed -i 's@^plugins.*@plugins=(extract autojump zsh-autosuggestions zsh-syntax-highlighting)@g'  ~/.zshrc




# debug命令
# docker rm -f debug && docker build -t ubuntu-debug:v1.0 --progress=plain --add-host=dl.google.com:220.181.174.225 . && docker run -d --rm --privileged --name debug --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -p 21800:22 -v /var/run/docker.sock:/var/run/docker.sock -v /home/wangmin/workspace:/root/workspace:rw ubuntu-debug:v1.0 && docker exec -it debug zsh


# 实际使用：docker run d --name wangmin --privileged --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -v /var/run/docker.sock:/var/run/docker.sock -v home/wangmin/workspace:/root/workspace:rw -p 2800:22 -p 2810-2820:8010:8020 ubuntu-dev:22.04


# 使用systemd启动容器参考连接：https://blog.csdn.net/kencaber/article/details/121980242
# 参考连接：https://blog.csdn.net/m0_37886429/article/details/80350659?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&utm_relevant_index=1
CMD [ "/lib/systemd/systemd", "log-level=info", "unit=sysinit.target" ]