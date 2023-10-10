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
    git build-essential flex libncurses-dev bison libelf-dev libssl-dev bc openssh-server bash-completion ansible


# github加速（非常重要，否则github下载代码将会相当之慢）
RUN git config --global url."https://hub.fgit.cf".insteadof "https://github.com"


# 开启SSHD服务, 错误的构建方式，镜像在构建过程中不应该启动sshd服务，应该在镜像的启动点启动sshd服务 (当然，如果使用Systemd作为一号进程，那么ssh服务自然会被启动)
RUN sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config && \
    sed -i 's/#AddressFamily any/AddressFamily any/g' /etc/ssh/sshd_config && \
    sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config && \
    sed -i 's/#PermitRootLogin yes/PermitRootLogin yes/g' /etc/ssh/sshd_config && \
    sed -i "s/UsePAM.*/UsePAM no/g" /etc/ssh/sshd_config



# shell命令行颜色提示; git分支名自动补全
# 设置shell提示符颜色 参考链接：https://www.jianshu.com/p/a1d698d1f7c8
RUN V=$(git -v|awk '{print $3}') && git clone -b v${V} https://github.com/git/git.git
RUN git clone https://github.com/Dreamacro/clash-ctl.git

# 切换脚本解释器为bash, Ubuntu默认的脚本解释器为dash
RUN chsh -s /bin/bash

# RUN echo '\nPS1="\[\e[37m\][\[\e[32m\]\u\[\e[37m\]@\[\e[35m\]\h\[\e[0m\] \[\e[36m\]\w\[\e[0m\]]\\$ "' >> ~/.bashrc

RUN cp git/contrib/completion/git-completion.bash ~/.git-completion.bash && echo '\nsource ~/.git-completion.bash' >> ~/.bashrc  && \
    cp git/contrib/completion/git-prompt.sh ~/.git-prompt.sh && \
    cat <<EOF >> ~/.bashrc

GIT_PS1_SHOWDIRTYSTATE=true
GIT_PS1_SHOWCOLORHINTS=true
GIT_PS1_SHOWSTASHSTATE=true
GIT_PS1_SHOWUNTRACKEDFILES=true
#GIT_PS1_SHOWUPSTREAM=auto
if [ -f ~/.git-completion.bash ]; then
  source ~/.git-prompt.sh
  PROMPT_COMMAND='__git_ps1 "\[\e[37m\][\[\e[32m\]\u\[\e[37m\]@\[\e[35m\]\h\[\e[0m\] \[\e[36m\]\w\[\e[0m\]]" "\\\$ "'
fi
EOF

# vscode 不然vscode安装好之后，老是找不到命令
RUN cat <<EOF >> ~/.bashrc

# 添加vscode可执行文件环境变量配置，否则有可能找不到vscode可执行文件
if [ -d ~/.vscode-server ];then
    code_latest_version=\$(ls -tral -1 --ignore=.* ~/.vscode-server/bin | sed -n '2p' | rev | cut -d' ' -f1 | rev)
    export PATH=\${HOME}/.vscode-server/bin/\${code_latest_version}/bin/remote-cli:\$PATH
fi
EOF

# 打开文件监听限制，否则当文件数量过多，linux会限制打开文件数量
RUN touch /etc/sysctl.conf && \
    echo "fs.inotify.max_user_watches = 9994288" > /etc/sysctl.conf

# TODO 解决启动时需要输入命令的问题，systemd应该直接启动
# TODO 解决进入容器后，exit无法退出问题


# docker构建时提示错误：output clipped, log limit 1MiB reached  解决方案如下，分别执行两条命令即可
# docker buildx create --use --name larger_log --driver-opt env.BUILDKIT_STEP_LOG_MAX_SIZE=50000000
# docker buildx build --progress plain .

# 测试：    docker run -d --rm --privileged --name debug --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -p 2800:22 -v /var/run/docker.sock:/var/run/docker.sock ubuntu-debug:v1.0

# 实际使用：docker run d --name wangmin --privileged --cap-add SYS_ADMIN --security-opt seccomp=unconfined --cgroup-parent=docker.slice --cgroupns private --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -v /var/run/docker.sock:/var/run/docker.sock -v home/wangmin/workspace:/root/workspace:rw -p 2800:22 -p 2810-2820:8010:8020 ubuntu-dev:22.04


# 使用systemd启动容器参考连接：https://blog.csdn.net/kencaber/article/details/121980242
# 参考连接：https://blog.csdn.net/m0_37886429/article/details/80350659?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-80350659-blog-121980242.235%5Ev38%5Epc_relevant_anti_vip&utm_relevant_index=1
CMD [ "/lib/systemd/systemd", "log-level=info", "unit=multi-user.target" ]