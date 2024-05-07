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

HOSTNAME=desktop

HELMVERSION=v3.12.2
IPADDR=192.168.11.10/24
GATEWAY=192.168.11.2
ETH=ens32

GOLANGVERSION=1.20.6

# 配置主机名
hostnamectl set-hostname ${HOSTNAME}

# 配置网络
tee /etc/netplan/00-installer-config.yaml << EOF
network:
  ethernets:
    ${ETH}:
      addresses:
      - ${IPADDR}
      nameservers:
        addresses:
        - 114.114.114.114
        - 8.8.8.8
        search: []
      routes:
      - to: default
        via: ${GATEWAY}
  version: 2
  renderer: networkd
EOF

netplan apply

# 关闭apt命令下载软件时进入交互界面询问是否需要重启服务
sed -i 's/#$nrconf{restart} = '"'"'i'"'"';/$nrconf{restart} = '"'"'a'"'"';/g' /etc/needrestart/needrestart.conf

# yum源加速 设置为清华源
tee /etc/apt/sources.list << 'EOF'
# See http://help.ubuntu.com/community/UpgradeNotes for how to upgrade to
# newer versions of the distribution.
deb http://cn.archive.ubuntu.com/ubuntu jammy main restricted
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy main restricted

## Major bug fix updates produced after the final release of the
## distribution.
deb http://cn.archive.ubuntu.com/ubuntu jammy-updates main restricted
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy-updates main restricted

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team. Also, please note that software in universe WILL NOT receive any
## review or updates from the Ubuntu security team.
deb http://cn.archive.ubuntu.com/ubuntu jammy universe
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy universe
deb http://cn.archive.ubuntu.com/ubuntu jammy-updates universe
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy-updates universe

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team, and may not be under a free licence. Please satisfy yourself as to
## your rights to use the software. Also, please note that software in
## multiverse WILL NOT receive any review or updates from the Ubuntu
## security team.
deb http://cn.archive.ubuntu.com/ubuntu jammy multiverse
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy multiverse
deb http://cn.archive.ubuntu.com/ubuntu jammy-updates multiverse
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy-updates multiverse

## N.B. software from this repository may not have been tested as
## extensively as that contained in the main release, although it includes
## newer versions of some applications which may provide useful features.
## Also, please note that software in backports WILL NOT receive any review
## or updates from the Ubuntu security team.
deb http://cn.archive.ubuntu.com/ubuntu jammy-backports main restricted universe multiverse
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy-backports main restricted universe multiverse

deb http://cn.archive.ubuntu.com/ubuntu jammy-security main restricted
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy-security main restricted
deb http://cn.archive.ubuntu.com/ubuntu jammy-security universe
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy-security universe
deb http://cn.archive.ubuntu.com/ubuntu jammy-security multiverse
# deb-src http://cn.archive.ubuntu.com/ubuntu jammy-security multiverse
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

# 由于ubuntu本身使用的内核较新，并且使用apt命令升级软件时，会自动更新操作系统内核版本，所以无需下载

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

# 禁用防火墙
ufw disable
ufw status

# 内核优化
ulimit -SHn 65535
cat >> /etc/security/limits.conf <<EOF
* soft nofile 655360
* hard nofile 131072
* soft nproc 655350
* hard nproc 655350
* seft memlock unlimited
* hard memlock unlimitedd
EOF

cat >> /etc/modules-load.d/ipvs.conf <<EOF
ip_vs
ip_vs_rr
ip_vs_wrr
ip_vs_sh
nf_conntrack
ip_tables
ip_set
xt_set
ipt_set
ipt_rpfilter
ipt_REJECT
ipip
EOF



# 优化内核参数
cat <<EOF > /etc/sysctl.d/k8s.conf
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1
fs.may_detach_mounts = 1
vm.overcommit_memory=1
vm.panic_on_oom=0
fs.inotify.max_user_watches=89100
fs.file-max=52706963
fs.nr_open=52706963
net.netfilter.nf_conntrack_max=2310720

fs.inotify.max_queued_events=1048576
fs.inotify.max_user_watches=1048576
fs.inotify.max_user_instances=1048576

net.ipv4.tcp_keepalive_time = 600
net.ipv4.tcp_keepalive_probes = 3
net.ipv4.tcp_keepalive_intvl =15
net.ipv4.tcp_max_tw_buckets = 36000
net.ipv4.tcp_tw_reuse = 1
net.ipv4.tcp_max_orphans = 327680
net.ipv4.tcp_orphan_retries = 3
net.ipv4.tcp_syncookies = 1
net.ipv4.tcp_max_syn_backlog = 16384
net.ipv4.ip_conntrack_max = 65536
net.ipv4.tcp_max_syn_backlog = 16384
net.ipv4.tcp_timestamps = 0
net.core.somaxconn = 16384

net.ipv6.conf.all.disable_ipv6 = 0
net.ipv6.conf.default.disable_ipv6 = 0
net.ipv6.conf.lo.disable_ipv6 = 0
net.ipv6.conf.all.forwarding = 1
EOF

sysctl --system

# 内存加载containerd相关内核模块，当前有效，重启无效
modprobe overlay
modprobe br_netfilter

# 持久化加载containerd相关内核模块，重新有效
cat <<EOF | tee /etc/modules-load.d/containerd.conf
overlay
br_netfilter
EOF


# 安装git
git --version && \
    git config --global user.name "wangmin" && \
    git config --global user.email "wangmin@skyguard.com.cn" && \
    git config --global url.ssh://git@gitcdteam.skyguardmis.com/.insteadOf https://gitcdteam.skyguardmis.com/

# 安装go
wget https://mirrors.aliyun.com/golang/go${GOLANGVERSION}.linux-amd64.tar.gz?spm=a2c6h.25603864.0.0.a6b07c45FOi9wZ -O /opt/go${GOLANGVERSION}.linux-amd64.tar.gz && cd /opt && rm -rf go${GOLANGVERSION} && \
    mkdir go${GOLANGVERSION} &&  \
    tar -zxf go${GOLANGVERSION}.linux-amd64.tar.gz -C go${GOLANGVERSION} && \
    rm -f go${GOLANGVERSION}.linux-amd64.tar.gz && \
    echo "export GOROOT=/opt/go${GOLANGVERSION}/go" >>  /etc/profile &&  \
    echo 'export GOPATH=/root/go' >>  /etc/profile && \
    echo 'PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> /etc/profile && \
    source /etc/profile && \
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

wget https://files.m.daocloud.io/get.helm.sh/helm-${HELMVERSION}-linux-amd64.tar.gz
tar -zxvf helm-${HELMVERSION}-linux-amd64.tar.gz && \
    mv linux-amd64/helm /usr/local/bin/helm && \
    helm help && rm -f helm-${HELMVERSION}-linux-amd64.tar.gz rm linux-amd64 -rf


reboot
