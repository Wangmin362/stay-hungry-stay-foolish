#!/bin/bash

# 执行命令遇到错误就退出
set -e
# 脚本中遇到不存在的变量就退出
# set -u
# 执行指令的时候，同时把指令输出，方便观察结果
set -x
# 执行管道的时候，如果前面的命令出错，管道后面的命令会停止
set -o pipefail

# TODO 对于可能的虚拟机，需要考虑修改网卡的UUID以及MAC地址（VMWare克隆时，会自动修改Mac地址）

# TODO 若需要修改网络打开下面的注释

#HOSTNAME=desktop
#HELMVERSION=v3.12.2
#IPADDR=192.168.11.10/24
#GATEWAY=192.168.11.2
#ETH=ens32
#
#GOLANGVERSION=1.20.6
#
## 配置主机名
#hostnamectl set-hostname ${HOSTNAME}
#
## 配置网络
#tee /etc/netplan/00-installer-config.yaml << EOF
#network:
#  ethernets:
#    ${ETH}:
#      addresses:
#      - ${IPADDR}
#      nameservers:
#        addresses:
#        - 114.114.114.114
#        - 8.8.8.8
#        search: []
#      routes:
#      - to: default
#        via: ${GATEWAY}
#  version: 2
#  renderer: networkd
#EOF
#
#netplan apply


# ubuntu 2004 关闭apt命令下载软件时进入交互界面询问是否需要重启服务，注意关闭之后，需要重新启动操作系统部分软件才会生效
export DEBIAN_FRONTEND=noninteractive

# ubuntu 2204 关闭apt命令下载软件时进入交互界面询问是否需要重启服务，注意关闭之后，需要重新启动操作系统部分软件才会生效
#sed -i 's/#$nrconf{restart} = '"'"'i'"'"';/$nrconf{restart} = '"'"'a'"'"';/g' /etc/needrestart/needrestart.conf

# 更新apt源为国内镜像源，方便更新软件
tee /etc/apt/sources.list << 'EOF'
#添加清华源
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security main restricted universe multiverse
EOF

apt update -y && apt upgrade -y


# 安装软件
apt update -y && apt upgrade -y && \
    DEBIAN_FRONTEND=noninteractive TZ=Asia/Shanghai apt -y install tzdata && \
    apt install -y man-db manpages-posix && apt install -y manpages-dev manpages-posix-dev && yes | unminimize && \
    apt install -y net-tools telnet sysstat bridge-utils bash-completion vim jq tar openssl iputils-ping lsof lvm2 \
    dnsutils curl gcc g++ automake autoconf make tree stress htop atop sysbench ipvsadm ipset conntrack ufw git \
    build-essential flex libncurses-dev bison libelf-dev libssl-dev bc openssh-server ansible unzip binutils \
    libc-ares-dev libtool pkg-config libsystemd-dev file texinfo nmap zsh autojump byobu language-pack-en \
    language-pack-zh-hans python3-pygments chroma rsync libseccomp-dev

# 命令自动补全
sed -i 's@^#if ! shopt -oq posix; then@if ! shopt -oq posix; then@g' /etc/bash.bashrc
sed -i 's@#  if \[ -f /usr/share/bash-completion/bash_completion \]; then@  if \[ -f /usr/share/bash-completion/bash_completion \]; then@g' /etc/bash.bashrc
sed -i 's@#    . /usr/share/bash-completion/bash_completion@    . /usr/share/bash-completion/bash_completion@g' /etc/bash.bashrc
sed -i 's@#  elif \[ -f /etc/bash_completion \]; then@  elif \[ -f /etc/bash_completion \]; then@g' /etc/bash.bashrc
sed -i 's@#    . /etc/bash_completion@    . /etc/bash_completion@g' /etc/bash.bashrc
sed -i 's@#  fi@  fi@g' /etc/bash.bashrc
sed -i 's@#fi@fi@g' /etc/bash.bashrc
source /etc/bash.bashrc


# 安装SSH服务
apt install -y openssh-server
sed -i 's/^#Port.*/Port 22/g' /etc/ssh/sshd_config
sed -i 's/^#AddressFamily.*/AddressFamily any/g' /etc/ssh/sshd_config
sed -i 's/^#ListenAddress.*/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config
sed -i 's/^#PermitRootLogin.*/PermitRootLogin yes/g' /etc/ssh/sshd_config
systemctl restart sshd.service

# 同步时间
apt install -y chrony
timedatectl set-timezone "Asia/Shanghai"
sed -i 's/ntp.ubuntu.com/ntp1.aliyun.com/'  /etc/chrony/chrony.conf
sed -i 's/0.ubuntu.pool.ntp.org/ntp2.aliyun.com/'  /etc/chrony/chrony.conf
sed -i 's/1.ubuntu.pool.ntp.org/ntp3.aliyun.com/'  /etc/chrony/chrony.conf
sed -i 's/2.ubuntu.pool.ntp.org/ntp4.aliyun.com/'  /etc/chrony/chrony.conf
if ! cat /etc/chrony/chrony.conf | grep ntp5.aliyun.com; then
  sed -i "24 a server ntp5.aliyun.com iburst\nserver ntp6.aliyun.com iburst\nserver ntp7.aliyun.com iburst\nserver 0.cn.pool.ntp.org iburst\nserver 1.cn.pool.ntp.org iburst\nserver 2.cn.pool.ntp.org iburst\nserver 3.cn.pool.ntp.org iburst\nserver time1.cloud.tencent.com iburst\nserver time2.cloud.tencent.com iburst\nserver time3.cloud.tencent.com iburst\nserver time4.cloud.tencent.com iburst\n\n"  /etc/chrony/chrony.conf
fi
systemctl restart chronyd
chronyc sources -v

# 设置为24小时格式
tee /etc/default/locale << 'EOF'
#添加清华源
LANG=en_US.UTF-8
LC_TIME=en_DK.UTF-8
EOF