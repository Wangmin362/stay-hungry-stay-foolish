 #!/bin/bash

# 执行命令遇到错误就退出
set -e
# 脚本中遇到不存在的变量就退出
set -u
# 执行指令的时候，同时把指令输出，方便观察结果
set -x
# 执行管道的时候，如果前面的命令出错，管道后面的命令会停止
set -o pipefail

# TODO 对于可能的虚拟机，需要考虑修改网卡的UUID以及MAC地址（VMWare克隆时，会自动修改Mac地址）

HOSTNAME=centos-pattern

IPADDR=192.168.11.10
NETMASK=255.255.255.0
GATEWAY=192.168.11.2
ETH=ifcfg-ens32

GITVERSION=2.41.0
GOLANGVERSION=1.20.6

# 配置主机名
hostnamectl set-hostname ${HOSTNAME}

# 配置网络
sed -i "s/ONBOOT=no/ONBOOT=yes/g" /etc/sysconfig/network-scripts/${ETH}
sed -i "s/BOOTPROTO=dhcp/BOOTPROTO=static/g" /etc/sysconfig/network-scripts/${ETH}
sed -i "s/ONBOOT=no/ONBOOT=yes/g" /etc/sysconfig/network-scripts/${ETH}
if ! cat /etc/sysconfig/network-scripts/${ETH} | grep IPADDR; then
  echo "IPADDR=${IPADDR}" >> /etc/sysconfig/network-scripts/${ETH}
else
  sed -i "s/IPADDR=.*/IPADDR=${IPADDR}/g" /etc/sysconfig/network-scripts/${ETH}
fi
if ! cat /etc/sysconfig/network-scripts/${ETH} | grep NETMASK; then
  echo "NETMASK=${NETMASK}" >> /etc/sysconfig/network-scripts/${ETH}
else
  sed -i "s/NETMASK=.*/NETMASK=${NETMASK}/g" /etc/sysconfig/network-scripts/${ETH}
fi
if ! cat /etc/sysconfig/network-scripts/${ETH} | grep GATEWAY; then
  echo "GATEWAY=${GATEWAY}" >> /etc/sysconfig/network-scripts/${ETH}
else
  sed -i "s/GATEWAY=.*/GATEWAY=${GATEWAY}/g" /etc/sysconfig/network-scripts/${ETH}
fi
if ! cat /etc/sysconfig/network-scripts/${ETH} | grep DNS1; then
  echo "DNS1=${GATEWAY}" >> /etc/sysconfig/network-scripts/${ETH}
else
  sed -i "s/DNS1=.*/DNS1=${GATEWAY}/g" /etc/sysconfig/network-scripts/${ETH}
fi
cat /etc/sysconfig/network-scripts/${ETH}
systemctl restart network

# yum源加速 设置为清华源
sed -e 's|^mirrorlist=|#mirrorlist=|g' \
    -e 's|^#baseurl=http://mirror.centos.org|baseurl=https://mirrors.tuna.tsinghua.edu.cn|g' \
    -i.bak /etc/yum.repos.d/CentOS-*.repo
yum install -y wget
# 安装EPEL源，阿里云镜像加速
wget -O /etc/yum.repos.d/epel.repo https://mirrors.aliyun.com/repo/epel-7.repo
# 更新YUM源 如果出现：[Errno 14] HTTP Error 404 - Not Found错误，可以参考：https://blog.51cto.com/waxyz/5336025
yum clean all && yum makecache && yum update -y
yum provides '*/applydeltarpm'
yum install -y deltarpm

# 安装软件
yum install -y net-tools net-tools telnet nmap sysstat lrszs dos2unix bind-utils bridge-utils \
    bash-completion vim jq psmisc nfs-utils yum-utils device-mapper-persistent-data lvm2 network-scripts tar \
    iproute passwd openssl curl zlib-devel curl-devel bzip2-devel openssl-devel ncurses-devel gcc package automake \
    autoconf make gcc-c++ cpio expat-devel gettext-devel zlib tree stress htop atop sysbench openssh-server chrony \
    ipvsadm ipset conntrack libseccomp
# 更新内核
uname -sr
rpm --import https://www.elrepo.org/RPM-GPG-KEY-elrepo.org
yum -y install https://www.elrepo.org/elrepo-release-7.el7.elrepo.noarch.rpm
sed -i "s@mirrorlist@#mirrorlist@g" /etc/yum.repos.d/elrepo.repo
sed -i "s@elrepo.org/linux@mirrors.tuna.tsinghua.edu.cn/elrepo@g" /etc/yum.repos.d/elrepo.repo
yum remove kernel-tools-libs.x86_64 kernel-tools.x86_64 -y
# 安装最新版本内核
yum --disablerepo=\* --enablerepo=elrepo-kernel install kernel-ml.x86_64 -y
yum --disablerepo=\* --enablerepo=elrepo-kernel install kernel-ml-tools.x86_64 -y
## 安装稳定版本内核
#yum --disablerepo=\* --enablerepo=elrepo-kernel install kernel-lt.x86_64 -y
#yum --disablerepo=\* --enablerepo=elrepo-kernel install kernel-lt-tools.x86_64 -y
grep "^menuentry" /boot/grub2/grub.cfg | cut -d "'" -f2
grub2-editenv list
sed -i 's/GRUB_TIMEOUT=.*/GRUB_TIMEOUT=0/g' /etc/default/grub
sed -i 's/GRUB_DEFAULT=.*/GRUP_DEFAULT=0/g' /etc/default/grub
grub2-mkconfig -o /boot/grub2/grub.cfg

# 禁用防火墙
sed -i 's/AllowZoneDrifting=yes/AllowZoneDrifting=no/g' /etc/firewalld/firewalld.conf
systemctl restart firewalld.service
systemctl disable firewalld.service
systemctl status firewalld.service

# 禁用SWAP分区
sed -ri 's/.*swap.*/#&/' /etc/fstab
swapoff -a && sysctl -w vm.swappiness=0
cat /etc/fstab

# 禁用SeLinux
setenforce 0
sed -i 's#SELINUX=enforcing#SELINUX=disabled#g' /etc/selinux/config
sestatus

# 安装SSH服务
yum install -y openssh-server
sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config
sed -i 's/#AddressFamily any/AddressFamily any/g' /etc/ssh/sshd_config
sed -i 's/#ListenAddress 0.0.0.0/ListenAddress 0.0.0.0/g' /etc/ssh/sshd_config
sed -i 's/#PermitRootLogin yes/PermitRootLogin yes/g' /etc/ssh/sshd_config
systemctl enable sshd.service && systemctl start sshd.service

# 同步时间
yum install -y chrony
timedatectl set-timezone "Asia/Shanghai"
sed -i 's/0.centos.pool.ntp.org/ntp1.aliyun.com/'  /etc/chrony.conf
sed -i 's/1.centos.pool.ntp.org/ntp2.aliyun.com/'  /etc/chrony.conf
sed -i 's/2.centos.pool.ntp.org/ntp3.aliyun.com/'  /etc/chrony.conf
sed -i 's/3.centos.pool.ntp.org/ntp4.aliyun.com/'  /etc/chrony.conf
if ! cat /etc/chrony.conf | grep ntp5.aliyun.com; then
  sed -i "7 a server ntp5.aliyun.com iburst\nserver ntp6.aliyun.com iburst\nserver ntp7.aliyun.com iburst\nserver 0.cn.pool.ntp.org iburst\nserver 1.cn.pool.ntp.org iburst\nserver 2.cn.pool.ntp.org iburst\nserver 3.cn.pool.ntp.org iburst\nserver time1.cloud.tencent.com iburst\nserver time2.cloud.tencent.com iburst\nserver time3.cloud.tencent.com iburst\nserver time4.cloud.tencent.com iburst\n\n"  /etc/chrony.conf
fi
systemctl enable chronyd && systemctl start chronyd && systemctl restart chronyd && systemctl status chronyd
chronyc sources -v

# K8S 网络配置
cat > /etc/NetworkManager/conf.d/calico.conf << EOF
[keyfile]
unmanaged-devices=interface-name:cali*;interface-name:tunl*
EOF
systemctl restart NetworkManager

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
lsmod | grep -e ip_vs -e nf_conntrack

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

systemctl restart systemd-modules-load.service

# 安装git
set +u
yum remove -y git
wget -P /opt https://mirrors.edge.kernel.org/pub/software/scm/git/git-${GITVERSION}.tar.xz --no-check-certificate && cd /opt && \
    tar -xvf git-${GITVERSION}.tar.xz && \
    rm -f  git-${GITVERSION}.tar.xz && \
    rm -f /usr/bin/git &&  \
    cd git-${GITVERSION} && \
    sh configure --prefix=/usr/local/git all && \
    make && \
    make install &&  \
    cd .. && rm -rf  git-${GITVERSION} && \
    echo "export PATH=$PATH:/usr/local/git/bin" >> /etc/profile && source /etc/profile &&  \
    git --version && \
    git config --global user.name "wangmin" && \
    git config --global user.email "wangmin@skyguard.com.cn" && \
    git config --global url.ssh://git@gitcdteam.skyguardmis.com/.insteadOf https://gitcdteam.skyguardmis.com/

# 安装go
set +u
GOLANGVERSION=1.20.6
wget -P /opt https://dl.google.com/go/go${GOLANGVERSION}.linux-amd64.tar.gz && cd /opt && \
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


reboot