
FROM 172.30.3.150/common/busybox:1.35.0 AS builder

RUN wget --no-check-certificate https://storage.googleapis.com/kubernetes-release/release/v1.19.16/bin/linux/amd64/kubectl

FROM 172.30.3.150/ops/ttyd:1.7.3-alpine

WORKDIR /
COPY --from=builder /kubectl /usr/bin
RUN chmod +x /usr/bin/kubectl
