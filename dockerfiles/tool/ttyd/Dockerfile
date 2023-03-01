FROM ubuntu:20.04 as builder

RUN apt-get update
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
RUN apt-get install -y --no-install-recommends build-essential cmake git libjson-c-dev libwebsockets-dev
RUN git config --global http.sslverify false && \
    git clone https://ghproxy.com/https://github.com/tsl0922/ttyd.git && \
    cd ttyd && mkdir build && cd build && cmake .. && make && make install


FROM ubuntu:20.04


RUN apt-get update && \
    # 时区设置
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo 'Asia/Shanghai' >/etc/timezone && \
    apt-get install -y --no-install-recommends tini wget libjson-c-dev libwebsockets-dev vim iputils-ping && \
    wget --no-check-certificate https://storage.googleapis.com/kubernetes-release/release/v1.19.16/bin/linux/amd64/kubectl && mv kubectl /usr/bin && chmod +x /usr/bin/kubectl
RUN apt-get install -y bash-completion && \
    # 命令行补全
    echo "source <(kubectl completion bash)" >> ~/.bashrc && \
    echo "source /usr/share/bash-completion/bash_completion" >> ~/.bashrc && \
    echo "source <(kubectl completion bash)" >> /etc/bash.bashrc && \
    echo "source /usr/share/bash-completion/bash_completion" >> /etc/bash.bashrc && \
    # 颜色提示
    echo 'PS1="\[\e[37m\][\[\e[32m\]\u\[\e[37m\]@\[\e[35m\]\h\[\e[0m\] \[\e[36m\]\w\[\e[0m\]]\\$ "' >> ~/.bashrc && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/local/bin/ttyd /usr/bin/ttyd

WORKDIR /root

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["ttyd", "bash"]
