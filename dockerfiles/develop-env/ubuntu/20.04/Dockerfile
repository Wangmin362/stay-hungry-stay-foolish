FROM ubuntu:20.04

RUN echo 'deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted universe multiverse\n\
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted universe multiverse\n\
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse\n\
deb http://security.ubuntu.com/ubuntu/ focal-security main restricted universe multiverse' > /etc/apt/sources.list

RUN cat /etc/apt/sources.list && apt install -y apt-transport-https ca-certificates && apt update -y && apt install -y net-tools telnet sysstat bridge-utils bash-completion vim jq tar \
    openssl iputils-ping lsof lvm2 dnsutils curl gcc automake autoconf make tree stress htop atop sysbench chrony \
    ipvsadm ipset conntrack ufw git