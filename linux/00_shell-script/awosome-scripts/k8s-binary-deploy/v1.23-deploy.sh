#!/bin/bash

# 修改Ip地址 todo 改为前缀匹配
sed -i 's/IPADDR=192.168.11.11/IPADDR=192.168.11.72/g' /etc/sysconfig/network-scripts/ifcfg-ens33
hostnamectl set-hostname node2
reboot


hostnamectl set-hostname k8s-master1

cat >> /etc/hosts <<EOF
192.168.11.71 k8s-master1
192.168.11.72 k8s-master2
192.168.11.73 k8s-node1
EOF
cat /etc/hosts

# 生成ssh key
git config --global user.name "wangmin"
git config --global user.email "wangmin@skyguard.com.cn"
ssh-keygen -t rsa -C "wangmin@skyguard.com.cn"

# 免密登录
ssh-copy-id -i /root/.ssh/id_rsa.pub k8s-master1
ssh-copy-id -i /root/.ssh/id_rsa.pub k8s-master2
ssh-copy-id -i /root/.ssh/id_rsa.pub k8s-node1

# 修改内核参数
# 为什么要执行 modprobe br_netfilter？
# 如不执行上面步骤则在修改/etc/sysctl.d/k8s.conf 文件后再执行 sysctl -p /etc/sysctl.d/k8s.conf 会出现如下报错：
# sysctl: cannot stat /proc/sys/net/bridge/bridge-nf-call-ip6tables: No such file or directory
# sysctl: cannot stat /proc/sys/net/bridge/bridge-nf-call-iptables: No such file or directory
# 所以解决方法就是提前加载相应模块
# 加载 br_netfilter 模块
modprobe br_netfilter
# 验证模块是否加载成功：
lsmod |grep br_netfilter
# 修改内核参数
cat > /etc/sysctl.d/k8s.conf <<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF
#使刚才修改的内核参数生效
sysctl -p /etc/sysctl.d/k8s.conf

# 配置阿里云 repo 安装 rzsz scp命令
yum install lrzsz openssh-clients yum-utils -y
#配置国内阿里云 docker 的 repo 源
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

#安装 ntpdate 或chrony服务都可以
yum install ntpdate -y
#跟网络源做同步
ntpdate cn.pool.ntp.org
#把时间同步做成计划任务
crontab -e * */1 * * * /usr/sbin/ntpdate	cn.pool.ntp.org
#重启 crond 服务
service crond restart

#安装 iptables
yum install iptables-services -y
#禁用 iptables
service iptables stop	&& systemctl disable iptables
#清空防火墙规则
iptables -F


yum install -y device-mapper-persistent-data lvm2 wget net-tools nfs-utils lrzsz gcc gcc-c++ make cmake libxml2-devel openssl-devel curl curl-devel unzip sudo ntp libaio-devel wget vim ncurses-devel autoconf automake zlib-devel python-devel epel-release openssh-server socat  ipvsadm conntrack ntpdate telnet rsync

yum install docker-ce  -y
systemctl start docker && systemctl enable docker.service && systemctl status docker
