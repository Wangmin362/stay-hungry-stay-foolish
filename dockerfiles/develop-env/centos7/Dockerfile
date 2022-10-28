FROM centos:centos7.9.2009
MAINTAINER wangmin

WORKDIR /opt

# golang环境
#COPY go1.16.15.linux-amd64.tar.gz golang
#COPY go1.17.13.linux-amd64.tar.gz golang
#COPY go1.18.7.linux-amd64.tar.gz golang
ADD go1.19.2.linux-amd64.tar.gz golang

# python环境
ADD Python-3.8.15.tgz python
#COPY Python-3.9.15.tgz python
#COPY Python-3.10.8.tgz python
#COPY Python-3.11.0.tgz python

# java环境
#COPY jdk-8u351-linux-x64.tar.gz java
#COPY jdk-11.0.17_linux-x64_bin.tar.gz java
#COPY jdk-17_linux-x64_bin.tar.gz java
#COPY jdk-19_linux-x64_bin.tar.gz java
# maven依赖
#COPY apache-maven-3.8.6-bin.tar.gz maven
# gradle依赖

# nodejs依赖
#COPY node-v18.12.0-linux-x64.tar.xz nodejs

WORKDIR /root
CMD ["/bin/bash"]
