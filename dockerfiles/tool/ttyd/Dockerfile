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
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo 'Asia/Shanghai' >/etc/timezone && \
    apt-get install -y --no-install-recommends tini wget libjson-c-dev libwebsockets-dev && \
    wget --no-check-certificate https://storage.googleapis.com/kubernetes-release/release/v1.19.16/bin/linux/amd64/kubectl && mv kubectl /usr/bin && chmod +x /usr/bin/kubectl && \
    apt-get install -y bash-completion && \
    echo "source <(kubectl completion bash)" >> ~/.bashrc && \
    echo "source /usr/share/bash-completion/bash_completion" >> ~/.bashrc && \
    echo "source <(kubectl completion bash)" >> /etc/bash.bashrc && \
    echo "source /usr/share/bash-completion/bash_completion" >> /etc/bash.bashrc && \
    cat "/etc/bash.bashrc" && \
    rm -rf /var/lib/apt/lists/* && \
    rm /bin/sh && ln -s /bin/bash /bin/sh
COPY --from=builder /usr/local/bin/ttyd /usr/bin/ttyd

WORKDIR /root

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["ttyd", "bash"]
