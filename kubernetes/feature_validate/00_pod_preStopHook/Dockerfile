FROM ubuntu:23.10

# 修改apt源为阿里源
RUN cat <<EOF > /etc/apt/sources.list
deb http://mirrors.aliyun.com/ubuntu/ mantic main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ mantic main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ mantic-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ mantic-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ mantic-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ mantic-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ mantic-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ mantic-proposed main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ mantic-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ mantic-backports main restricted universe multiverse
EOF


RUN apt update -y && apt install -y curl

COPY main /main
RUN chmod +x /main
EXPOSE 19090
CMD ["./main"]

