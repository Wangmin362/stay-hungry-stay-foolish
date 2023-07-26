FROM ubuntu:20.04

RUN echo 'deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy main restricted universe multiverse\n\
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-updates main restricted universe multiverse\n\
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-backports main restricted universe multiverse\n\
deb http://security.ubuntu.com/ubuntu/ jammy-security main restricted universe multiverse' > /etc/apt/sources.list

RUN cat /etc/apt/sources.list && apt update -y && apt upgrade -y && apt install -y net-tools telnet sysstat bridge-utils bash-completion vim jq tar \
    openssl iputils-ping lsof lvm2 dnsutils curl gcc automake autoconf make tree stress htop atop sysbench chrony \
    ipvsadm ipset conntrack ufw git