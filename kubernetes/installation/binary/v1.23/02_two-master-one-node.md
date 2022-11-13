# 1. 网段规划

[参考文档](https://www.cnblogs.com/fengdejiyixx/p/16576021.html)
> 安装单节点集群的目的: 方便调试K8S源码，理解其设计思路，因此多节点集群没有意义，反而增加了调试的复杂度

- K8S环境规划：
  - Pod网段：`10.0.0.0/16`
  - Service网段：`10.255.0.0/16`
- 实验环境规划：
  - 操作系统：`Centso7.9`
  - 配置：1GB, 2vCpu, 100G硬盘

| 集群角色 |      IP       |    主机名    |                         安装组件                         |
| ------- | ------------- | ----------- | ------------------------------------------------------- |
| 控制节点 | 192.168.11.71 | k8s-master1 | api-server, controller-manager, scheduler, etcd, docker |
| 控制节点 | 192.168.11.72 | k8s-master2 | api-server, controller-manager, scheduler, etcd, docker |
| 工作节点 | 192.168.11.73 | k8s-node1   | kubelet, kube-proxy, docker, calico, coredns            |

# 2. 环境准备
## 2.1. 配置静态IP地址以及设置主机名
```shell
sed -i 's/IPADDR=192.168.11.11/IPADDR=192.168.11.73/g' /etc/sysconfig/network-scripts/ifcfg-ens33
hostnamectl set-hostname k8s-node1
reboot

```

## 2.2. 配置HOST文件

```shell
tee /etc/hosts << 'EOF'
192.168.11.71 k8s-master1
192.168.11.72 k8s-master2
192.168.11.73 k8s-node1
EOF

```

## 2.3. 配置免密登录

生成`ssh key`
```shell
git config --global user.name "wangmin"
git config --global user.email "wangmin@skyguard.com.cn"
ssh-keygen -t rsa -C "wangmin@skyguard.com.cn"

```

配置免密登录
> <font size=4 color=Red>**注意：免密登录只需要在master1上配置，因为需要把master1上的配置文件拷贝到其他节点**</font>
```shell
ssh-copy-id -i .ssh/id_rsa.pub k8s-master1
ssh-copy-id -i .ssh/id_rsa.pub k8s-master2
ssh-copy-id -i .ssh/id_rsa.pub k8s-node1

```

## 2.4. 查看firewalld防火墙， selinux, swap交换分区

```shell
systemctl status firewalld
getenforce
free -h

[root@k8s-master1 ~]#
[root@k8s-master1 ~]# systemctl status firewalld
● firewalld.service - firewalld - dynamic firewall daemon
   Loaded: loaded (/usr/lib/systemd/system/firewalld.service; disabled; vendor preset: enabled)
   Active: inactive (dead)
     Docs: man:firewalld(1)
[root@k8s-master1 ~]# getenforce
Disabled
[root@k8s-master1 ~]# free -h
              total        used        free      shared  buff/cache   available
Mem:           946M        186M        655M        6.6M        104M        634M
Swap:            0B          0B          0B
[root@k8s-master1 ~]#
[root@k8s-master1 ~]#
```

## 2.4. 时间同步

```shell
yum install -y chrony
systemctl restart chronyd
systemctl enable chronyd
chronyc sources

```

## 2.5. 修改内核参数

```shell
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

```

## 2.6. 配置阿里云REPO

```shell
#安装 rzsz scp命令
yum install lrzsz openssh-clients yum-utils -y
#配置国内阿里云 docker 的 repo 源
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

```

## 2.7. 安装iptables

```shell
#安装 iptables
yum install iptables-services -y
#禁用 iptables
service iptables stop	&& systemctl disable iptables
#清空防火墙规则 
iptables -F

```


## 2.8. 在worker上安装docker环境

> 如果只需要`worker`运行`pod`，不需要`master`节点运行`pod`，那么无需在`master`节点上安装`docker`以及`kubelet`
```shell
yum install docker-ce  -y
systemctl start docker && systemctl enable docker.service && systemctl status docker

```

配置镜像加速
```shell
tee /etc/docker/daemon.json << 'EOF'
{
"registry-mirrors":["https://rsbud4vc.mirror.aliyuncs.com","https://registry.docker-cn.com","https://docker.mirrors.ustc.edu.cn","https://dockerhub.azk8s.cn","http://hub-mirror.c.163.com","http://qtid6917.mirror.aliyuncs.com", "https://rncxm540.mirror.aliyuncs.com"],
"exec-opts": ["native.cgroupdriver=systemd"]
}
EOF
systemctl daemon-reload && systemctl restart docker && systemctl status docker

```


## 2.9. 安装基础软件包

```shell
yum install -y device-mapper-persistent-data lvm2 wget net-tools nfs-utils lrzsz gcc gcc-c++ make cmake libxml2-devel openssl-devel curl curl-devel unzip sudo ntp libaio-devel wget vim ncurses-devel autoconf automake zlib-devel python-devel epel-release openssh-server socat  ipvsadm conntrack ntpdate telnet rsync

```


## 2.10. 配置工作目录
<font size=4 color=Gold>**以后所有的操作都在 /data/work上执行**</font>
```shell
mkdir -p /data/work
cd /data/work

```

# 3. Master节点部署etcd集群
## 3.1. 准备配置文件
<font size=4 color=Gold>**以下命令在k8s-master1上执行即可，部分命令会通过ssh的方式链接到k8s-master2上执行**</font>
```shell
mkdir -p /etc/etcd          # 配置文件存放目录
mkdir -p /etc/etcd/ssl      # 证书文件存放目录

# 在k8s master2上创建目录
ssh k8s-master2
mkdir -p /etc/etcd          # 配置文件存放目录
mkdir -p /etc/etcd/ssl      # 证书文件存放目录
exit

export CFSSL_VERSION=1.6.3
wget -O cfssl_linux-amd64 https://ghproxy.com/https://github.com/cloudflare/cfssl/releases/download/v${CFSSL_VERSION}/cfssl_${CFSSL_VERSION}_linux_amd64
wget -O cfssljson_linux-amd64 https://ghproxy.com/https://github.com/cloudflare/cfssl/releases/download/v${CFSSL_VERSION}/cfssljson_${CFSSL_VERSION}_linux_amd64
wget -O cfssl-certinfo_linux-amd64 https://ghproxy.com/https://github.com/cloudflare/cfssl/releases/download/v${CFSSL_VERSION}/cfssl-certinfo_${CFSSL_VERSION}_linux_amd64

chmod +x cfssl*
mv cfssl_linux-amd64 /usr/local/bin/cfssl
mv cfssljson_linux-amd64 /usr/local/bin/cfssljson
mv cfssl-certinfo_linux-amd64 /usr/local/bin/cfssl-certinfo

tee ca-csr.json << 'EOF'
{
  "CN": "kubernetes",
  "key": {
      "algo": "rsa",
      "size": 2048
  },
  "names": [
    {
      "C": "CN",
      "ST": "Hubei",
      "L": "Wuhan",
      "O": "k8s",
      "OU": "system"
    }
  ],
  "ca": {
          "expiry": "87600h"
  }
}
EOF

# 该命令执行之后会在当前目录中生成 ca.csr  ca-key.pem  ca.pem三个文件
cfssl gencert -initca ca-csr.json  | cfssljson -bare ca

tee ca-config.json << 'EOF'
{
  "signing": {
      "default": {
          "expiry": "87600h"
        },
      "profiles": {
          "kubernetes": {
              "usages": [
                  "signing",
                  "key encipherment",
                  "server auth",
                  "client auth"
              ],
              "expiry": "87600h"
          }
      }
  }
}
EOF

tee etcd-csr.json << 'EOF'
{
  "CN": "etcd",
  "hosts": [
    "127.0.0.1",
    "192.168.11.71",
    "192.168.11.72",
    "192.168.11.73"
  ],
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [{
    "C": "CN",
    "ST": "Hubei",
    "L": "Wuhan",
    "O": "k8s",
    "OU": "system"
  }]
}
EOF

# 该命令执行之后会在当前命令生成 etcd-key.pem  etcd.pem 这两个文件
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes etcd-csr.json | cfssljson  -bare etcd

export ETCD_VERSION=v3.4.13
wget https://ghproxy.com/https://github.com/etcd-io/etcd/releases/download/${ETCD_VERSION}/etcd-${ETCD_VERSION}-linux-amd64.tar.gz
tar -zxvf etcd-${ETCD_VERSION}-linux-amd64.tar.gz
cp -ar etcd-${ETCD_VERSION}-linux-amd64/etcd* /usr/local/bin
chmod +x /usr/local/bin/etcd*

scp -r etcd-${ETCD_VERSION}-linux-amd64/etcd* k8s-master2:/usr/local/bin
ssh root@k8s-master2
chmod +x /usr/local/bin/etcd*
exit

tee etcd.conf << 'EOF'
#[Member]
ETCD_NAME="etcd1"
ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
ETCD_LISTEN_PEER_URLS="https://192.168.11.71:2380"
ETCD_LISTEN_CLIENT_URLS="https://192.168.11.71:2379,http://127.0.0.1:2379"
#[Clustering]
ETCD_INITIAL_ADVERTISE_PEER_URLS="https://192.168.11.71:2380"
ETCD_ADVERTISE_CLIENT_URLS="https://192.168.11.71:2379"
ETCD_INITIAL_CLUSTER="etcd1=https://192.168.11.71:2380,etcd2=https://192.168.11.72:2380"
ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
ETCD_INITIAL_CLUSTER_STATE="new"
EOF

tee etcd.service << 'EOF'
[Unit]
Description=Etcd Server
After=network.target
After=network-online.target
Wants=network-online.target

[Service]
Type=notify
EnvironmentFile=-/etc/etcd/etcd.conf
WorkingDirectory=/var/lib/etcd/
ExecStart=/usr/local/bin/etcd \
  --cert-file=/etc/etcd/ssl/etcd.pem \
  --key-file=/etc/etcd/ssl/etcd-key.pem \
  --trusted-ca-file=/etc/etcd/ssl/ca.pem \
  --peer-cert-file=/etc/etcd/ssl/etcd.pem \
  --peer-key-file=/etc/etcd/ssl/etcd-key.pem \
  --peer-trusted-ca-file=/etc/etcd/ssl/ca.pem \
  --peer-client-cert-auth \
  --client-cert-auth
Restart=on-failure
RestartSec=5
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF

# 把证书以及配置文件拷贝到master02节点
mkdir -p /etc/etcd/ssl/
cp ca*.pem /etc/etcd/ssl/
cp etcd*.pem /etc/etcd/ssl/
cp etcd.conf /etc/etcd/
cp etcd.service /usr/lib/systemd/system/
for i in k8s-master2 ;do rsync -vaz etcd.conf $i:/etc/etcd/;done
for i in k8s-master2 ;do rsync -vaz etcd*.pem ca*.pem $i:/etc/etcd/ssl/;done
for i in k8s-master2 ;do rsync -vaz etcd.service $i:/usr/lib/systemd/system/;done

mkdir -p /var/lib/etcd/default.etcd
chmod 777 /var/lib/etcd/default.etcd

# 登录到k8s-master2上
ssh k8s-master2
mkdir -p /var/lib/etcd/default.etcd
chmod 777 /var/lib/etcd/default.etcd

# 修改master2上etcd的配置
tee /etc/etcd/etcd.conf << 'EOF'
#[Member]
ETCD_NAME="etcd2"
ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
ETCD_LISTEN_PEER_URLS="https://192.168.11.72:2380"
ETCD_LISTEN_CLIENT_URLS="https://192.168.11.72:2379,http://127.0.0.1:2379"
#[Clustering]
ETCD_INITIAL_ADVERTISE_PEER_URLS="https://192.168.11.72:2380"
ETCD_ADVERTISE_CLIENT_URLS="https://192.168.11.72:2379"
ETCD_INITIAL_CLUSTER="etcd1=https://192.168.11.71:2380,etcd2=https://192.168.11.72:2380"
ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
ETCD_INITIAL_CLUSTER_STATE="new"
EOF
# 退出k8s-master2,回到k8s-master1上
exit

```

## 3.2. 在master节点上启动etcd
<font size=4 color=Gold>**现在master1上启动etcd, 然后在master2上启动etcd服务**</font>
```shell
# 4. master1 以及 master2上启动服务
systemctl daemon-reload && systemctl enable etcd.service && systemctl start etcd.service

```


## 3.3. 查看etcd的状态

```shell
# 5. 查看etcd集群
ETCDCTL_API=3 && /usr/local/bin/etcdctl --write-out=table --cacert=/etc/etcd/ssl/ca.pem --cert=/etc/etcd/ssl/etcd.pem --key=/etc/etcd/ssl/etcd-key.pem --endpoints=https://192.168.11.71:2379,https://192.168.11.72:2379 endpoint health

```

# 4. 下载kubernetes组件

```shell
wget https://storage.googleapis.com/kubernetes-release/release/v1.23.13/kubernetes-server-linux-amd64.tar.gz
tar zxvf kubernetes-server-linux-amd64.tar.gz
cd  kubernetes/server/bin/
cp kube-apiserver kube-controller-manager kube-scheduler kubectl /usr/local/bin/
rsync -vaz kube-apiserver kube-controller-manager kube-scheduler kubectl k8s-master2:/usr/local/bin/
scp kubelet kube-proxy k8s-node1:/usr/local/bin/
cd /data/work/
mkdir -p /etc/kubernetes/ssl
mkdir /var/log/kubernetes

ssh k8s-master2
mkdir -p /etc/kubernetes/ssl
mkdir /var/log/kubernetes
exit

ssh k8s-node1
mkdir -p /etc/kubernetes/ssl
mkdir /var/log/kubernetes
exit

```

# 5. Master节点部署apiServer组件
## 5.1. 准备配置文件
```shell
cd /data/work

# 格式：token，用户名，UID，用户组
tee token.csv << 'EOF'
$(head -c 16 /dev/urandom | od -An -t x | tr -d ' '),kubelet-bootstrap,10001,"system:kubelet-bootstrap"
EOF

tee kube-apiserver-csr.json << 'EOF'
{
  "CN": "kubernetes",
  "hosts": [
    "127.0.0.1",
    "192.168.11.71",
    "192.168.11.72",
    "192.168.11.73",
    "10.255.0.1",
    "kubernetes",
    "kubernetes.default",
    "kubernetes.default.svc",
    "kubernetes.default.svc.cluster",
    "kubernetes.default.svc.cluster.local"
  ],
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CN",
      "ST": "Hubei",
      "L": "Wuhan",
      "O": "k8s",
      "OU": "system"
    }
  ]
}
EOF

# 生成证书
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-apiserver-csr.json | cfssljson -bare kube-apiserver

# api-server配置文件
tee kube-apiserver.conf << 'EOF'
KUBE_APISERVER_OPTS="--enable-admission-plugins=NamespaceLifecycle,NodeRestriction,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota \
  --anonymous-auth=false \
  --bind-address=192.168.11.71 \
  --secure-port=6443 \
  --advertise-address=192.168.11.71 \
  --insecure-port=0 \
  --authorization-mode=Node,RBAC \
  --runtime-config=api/all=true \
  --enable-bootstrap-token-auth \
  --service-cluster-ip-range=10.255.0.0/16 \
  --token-auth-file=/etc/kubernetes/token.csv \
  --service-node-port-range=30000-50000 \
  --tls-cert-file=/etc/kubernetes/ssl/kube-apiserver.pem  \
  --tls-private-key-file=/etc/kubernetes/ssl/kube-apiserver-key.pem \
  --client-ca-file=/etc/kubernetes/ssl/ca.pem \
  --kubelet-client-certificate=/etc/kubernetes/ssl/kube-apiserver.pem \
  --kubelet-client-key=/etc/kubernetes/ssl/kube-apiserver-key.pem \
  --service-account-key-file=/etc/kubernetes/ssl/ca-key.pem \
  --service-account-signing-key-file=/etc/kubernetes/ssl/ca-key.pem  \
  --service-account-issuer=https://kubernetes.default.svc.cluster.local \
  --etcd-cafile=/etc/etcd/ssl/ca.pem \
  --etcd-certfile=/etc/etcd/ssl/etcd.pem \
  --etcd-keyfile=/etc/etcd/ssl/etcd-key.pem \
  --etcd-servers=https://192.168.11.71:2379,https://192.168.11.72:2379 \
  --enable-swagger-ui=true \
  --allow-privileged=true \
  --apiserver-count=3 \
  --audit-log-maxage=30 \
  --audit-log-maxbackup=3 \
  --audit-log-maxsize=100 \
  --audit-log-path=/var/log/kube-apiserver-audit.log \
  --event-ttl=1h \
  --alsologtostderr=true \
  --logtostderr=false \
  --log-dir=/var/log/kubernetes \
  --v=4"
EOF

# 启动api-server
tee kube-apiserver.service << 'EOF'
[Unit]
Description=Kubernetes API Server
Documentation=https://github.com/kubernetes/kubernetes
After=etcd.service
Wants=etcd.service

[Service]
EnvironmentFile=-/etc/kubernetes/kube-apiserver.conf
ExecStart=/usr/local/bin/kube-apiserver $KUBE_APISERVER_OPTS
Restart=on-failure
RestartSec=5
Type=notify
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF

# 拷贝证书文件到相应的目录，同时也拷贝到master2节点
cp ca*.pem /etc/kubernetes/ssl
cp kube-apiserver*.pem /etc/kubernetes/ssl/
cp token.csv /etc/kubernetes/
cp kube-apiserver.conf /etc/kubernetes/
cp kube-apiserver.service /usr/lib/systemd/system/

# 同步 master2
rsync -vaz token.csv k8s-master2:/etc/kubernetes/
rsync -vaz kube-apiserver*.pem k8s-master2:/etc/kubernetes/ssl/
rsync -vaz ca*.pem k8s-master2:/etc/kubernetes/ssl/
rsync -vaz kube-apiserver.conf k8s-master2:/etc/kubernetes/
rsync -vaz kube-apiserver.service k8s-master2:/usr/lib/systemd/system/
### 注意！！！！k8s-master2 配置文件 kube-apiserver.conf 的 IP 地址修改为实际的本机 IP

ssh k8s-master2
tee /etc/kubernetes/kube-apiserver.conf << 'EOF'
KUBE_APISERVER_OPTS="--enable-admission-plugins=NamespaceLifecycle,NodeRestriction,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota \
  --anonymous-auth=false \
  --bind-address=192.168.11.72 \
  --secure-port=6443 \
  --advertise-address=192.168.11.72 \
  --insecure-port=0 \
  --authorization-mode=Node,RBAC \
  --runtime-config=api/all=true \
  --enable-bootstrap-token-auth \
  --service-cluster-ip-range=10.255.0.0/16 \
  --token-auth-file=/etc/kubernetes/token.csv \
  --service-node-port-range=30000-50000 \
  --tls-cert-file=/etc/kubernetes/ssl/kube-apiserver.pem  \
  --tls-private-key-file=/etc/kubernetes/ssl/kube-apiserver-key.pem \
  --client-ca-file=/etc/kubernetes/ssl/ca.pem \
  --kubelet-client-certificate=/etc/kubernetes/ssl/kube-apiserver.pem \
  --kubelet-client-key=/etc/kubernetes/ssl/kube-apiserver-key.pem \
  --service-account-key-file=/etc/kubernetes/ssl/ca-key.pem \
  --service-account-signing-key-file=/etc/kubernetes/ssl/ca-key.pem  \
  --service-account-issuer=https://kubernetes.default.svc.cluster.local \
  --etcd-cafile=/etc/etcd/ssl/ca.pem \
  --etcd-certfile=/etc/etcd/ssl/etcd.pem \
  --etcd-keyfile=/etc/etcd/ssl/etcd-key.pem \
  --etcd-servers=https://192.168.11.71:2379,https://192.168.11.72:2379 \
  --enable-swagger-ui=true \
  --allow-privileged=true \
  --apiserver-count=3 \
  --audit-log-maxage=30 \
  --audit-log-maxbackup=3 \
  --audit-log-maxsize=100 \
  --audit-log-path=/var/log/kube-apiserver-audit.log \
  --event-ttl=1h \
  --alsologtostderr=true \
  --logtostderr=false \
  --log-dir=/var/log/kubernetes \
  --v=4"
EOF
cat /etc/kubernetes/kube-apiserver.conf
exit

```


## 5.2. master节点启动apiserver服务
```shell
# 6. 启动kube-apiserver master1 以及 master2同时执行
systemctl daemon-reload && systemctl enable kube-apiserver && systemctl start kube-apiserver

```

## 5.3. 查看apiserver状态
```
# 6. 查看kube-apiserver启动状态
systemctl status kube-apiserver
# 7. 查看api-server授权状态
curl --insecure https://192.168.11.71:6443/
#{
# 8. "kind": "Status",
# 9. "apiVersion": "v1",
# 10. "metadata": {},
# 11. "status": "Failure",
# 12. "message": "Unauthorized",
# 13. "reason": "Unauthorized",
# 14. "code": 401
#}
# 15. 上面看到 401，这个是正常的的状态，还没认证

```

# 6. 部署kubectl组件

```shell
cd /data/work

# master1上导出这个环境变量,为kubectl提供配置文件路径
export KUBECONFIG=/etc/kubernetes/admin.conf

# 创建csr请求文件,准备创建admin.conf文件
tee admin-csr.json << 'EOF'
{
  "CN": "admin",
  "hosts": [],
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CN",
      "ST": "Hubei",
      "L": "Wuhan",
      "O": "system:masters",
      "OU": "system"
    }
  ]
}
EOF

# 生成客户端证书
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes admin-csr.json | cfssljson -bare admin
# 拷贝证书
cp admin*.pem /etc/kubernetes/ssl/

# 创建 kubeconfig安全上下文
# 注意: kube.config文件会自动生成,无需手动创建
kubectl config set-cluster kubernetes --certificate-authority=ca.pem --embed-certs=true --server=https://192.168.11.71:6443 --kubeconfig=kube.config

# 设置客户端认证参数
kubectl config set-credentials admin --client-certificate=admin.pem --client-key=admin-key.pem --embed-certs=true --kubeconfig=kube.config

# 设置上下文参数
kubectl config set-context kubernetes --cluster=kubernetes --user=admin --kubeconfig=kube.config

# 设置当前上下文
kubectl config use-context kubernetes --kubeconfig=kube.config
mkdir ~/.kube -p
cp kube.config ~/.kube/config
cp kube.config /etc/kubernetes/admin.conf

# 授权kubernetes证书访问kubelet api权限
kubectl create clusterrolebinding kube-apiserver:kubelet-apis --clusterrole=system:kubelet-api-admin --user kubernetes

# 查看集群信息
kubectl cluster-info
#Kubernetes control plane is running at https://192.168.7.10:6443
#To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

# 查看集群组件状态
kubectl get componentstatuses
#Warning: v1 ComponentStatus is deprecated in v1.19+
#NAME                 STATUS      MESSAGE              ERROR
#controller-manager   Unhealthy   Get "https://127.0.0.1:10257/healthz": dial tcp 127.0.0.1:10257: connect: connection refused
#scheduler            Unhealthy   Get "https://127.0.0.1:10259/healthz": dial tcp 127.0.0.1:10259: connect: connection refused
#etcd-0               Healthy     {"health":"true"}
#etcd-1               Healthy     {"health":"true"}


kubectl get all --all-namespaces
#NAMESPACE   NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
#default     service/kubernetes   ClusterIP   10.255.0.1   <none>        443/TCP   36m

# 同步kubeconfig文件到master2
rsync -vaz /root/.kube/config k8s-master2:/root/.kube/

# kubectl 命令补全 master1和master2都需要执行
yum install -y bash-completion
source /usr/share/bash-completion/bash_completion
source <(kubectl completion bash)
kubectl completion bash > ~/.kube/completion.bash.inc
source '/root/.kube/completion.bash.inc'
source $HOME/.bash_profile

```

# 7. 部署kube-controller-mansger组件
## 7.1. 准备配置文件

```shell
cd /data/work

# 创建 controller-manager csr请求文档
tee kube-controller-manager-csr.json << 'EOF'
{
    "CN": "system:kube-controller-manager",
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "hosts": [
      "127.0.0.1",
      "192.168.11.71",
      "192.168.11.72",
      "192.168.11.73"
    ],
    "names": [
      {
        "C": "CN",
        "ST": "Hubei",
        "L": "Wuhan",
        "O": "system:kube-controller-manager",
        "OU": "system"
      }
    ]
}
EOF


# 创建controller-manager证书
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-controller-manager-csr.json | cfssljson -bare kube-controller-manager


# 创建Kube-controller-manager的kubeconfig
kubectl config set-cluster kubernetes --certificate-authority=ca.pem --embed-certs=true --server=https://192.168.11.71:6443 --kubeconfig=kube-controller-manager.kubeconfig

# 查看创建的kubeconfig文件
cat kube-controller-manager.kubeconfig

# 设置客户端认证参数
kubectl config set-credentials system:kube-controller-manager --client-certificate=kube-controller-manager.pem --client-key=kube-controller-manager-key.pem --embed-certs=true --kubeconfig=kube-controller-manager.kubeconfig


# 设置上下文参数
kubectl config set-context system:kube-controller-manager --cluster=kubernetes --user=system:kube-controller-manager --kubeconfig=kube-controller-manager.kubeconfig


# 设置当前上下文
kubectl config use-context system:kube-controller-manager --kubeconfig=kube-controller-manager.kubeconfig

# 创建controller-manager的配置文件
tee kube-controller-manager.conf << 'EOF'
KUBE_CONTROLLER_MANAGER_OPTS="--port=0 \
  --secure-port=10257 \
  --bind-address=127.0.0.1 \
  --kubeconfig=/etc/kubernetes/kube-controller-manager.kubeconfig \
  --service-cluster-ip-range=10.255.0.0/16 \
  --cluster-name=kubernetes \
  --cluster-signing-cert-file=/etc/kubernetes/ssl/ca.pem \
  --cluster-signing-key-file=/etc/kubernetes/ssl/ca-key.pem \
  --allocate-node-cidrs=true \
  --cluster-cidr=10.0.0.0/16 \
  --experimental-cluster-signing-duration=87600h \
  --root-ca-file=/etc/kubernetes/ssl/ca.pem \
  --service-account-private-key-file=/etc/kubernetes/ssl/ca-key.pem \
  --leader-elect=true \
  --feature-gates=RotateKubeletServerCertificate=true \
  --controllers=*,bootstrapsigner,tokencleaner \
  --horizontal-pod-autoscaler-sync-period=10s \
  --tls-cert-file=/etc/kubernetes/ssl/kube-controller-manager.pem \
  --tls-private-key-file=/etc/kubernetes/ssl/kube-controller-manager-key.pem \
  --use-service-account-credentials=true \
  --alsologtostderr=true \
  --logtostderr=false \
  --log-dir=/var/log/kubernetes \
  --v=2"
EOF

# 创建controller-manager systemctl启动配置文件
tee kube-controller-manager.service << 'EOF'
[Unit]
Description=Kubernetes Controller Manager
Documentation=https://github.com/kubernetes/kubernetes
[Service]
EnvironmentFile=-/etc/kubernetes/kube-controller-manager.conf
ExecStart=/usr/local/bin/kube-controller-manager $KUBE_CONTROLLER_MANAGER_OPTS
Restart=on-failure
RestartSec=5
[Install]
WantedBy=multi-user.target
EOF

# 复制配置文件到对应目录
cp kube-controller-manager*.pem /etc/kubernetes/ssl/
cp kube-controller-manager.kubeconfig /etc/kubernetes/
cp kube-controller-manager.conf /etc/kubernetes/
cp kube-controller-manager.service /usr/lib/systemd/system/
rsync -vaz kube-controller-manager.pem k8s-master2:/etc/kubernetes/ssl/
rsync -vaz kube-controller-manager.kubeconfig kube-controller-manager.conf k8s-master2:/etc/kubernetes/
rsync -vaz kube-controller-manager.service k8s-master2:/usr/lib/systemd/system/

```

## 7.2. 所有master节点启动服务
```shell
# 启动controller-manager服务
systemctl daemon-reload  &&systemctl enable kube-controller-manager && systemctl start kube-controller-manager && systemctl status kube-controller-manager

# 查看集群组件的状态
kubectl get componentstatuses
#Warning: v1 ComponentStatus is deprecated in v1.19+
#NAME                 STATUS      MESSAGE                                                                                        #ERROR
#scheduler            Unhealthy   Get "https://127.0.0.1:10259/healthz": #dial tcp 127.0.0.1:10259: connect: connection refused
#controller-manager   Healthy     ok
#etcd-1               Healthy     {"health":"true"}
#etcd-0               Healthy     {"health":"true"}

```




# 8. 部署kube-scheduler组件
## 8.1. 准备配置文件
```shell
cd /data/work

tee kube-scheduler-csr.json << 'EOF'
{
    "CN": "system:kube-scheduler",
    "hosts": [
      "127.0.0.1",
      "192.168.11.71",
      "192.168.11.72",
      "192.168.11.73"
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
      {
        "C": "CN",
        "ST": "Hubei",
        "L": "Wuhan",
        "O": "system:kube-scheduler",
        "OU": "system"
      }
    ]
}
EOF

# 生成kube-scheduler证书
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-scheduler-csr.json | cfssljson -bare kube-scheduler

# 创建kube-scheduler的kubeconfig文件
kubectl config set-cluster kubernetes --certificate-authority=ca.pem --embed-certs=true --server=https://192.168.11.71:6443 --kubeconfig=kube-scheduler.kubeconfig

# 设置客户端认证参数
kubectl config set-credentials system:kube-scheduler --client-certificate=kube-scheduler.pem --client-key=kube-scheduler-key.pem --embed-certs=true --kubeconfig=kube-scheduler.kubeconfig

# 设置上下文参数
kubectl config set-context system:kube-scheduler --cluster=kubernetes --user=system:kube-scheduler --kubeconfig=kube-scheduler.kubeconfig

# 设置当前上下文
```shell
kubectl config use-context system:kube-scheduler --kubeconfig=kube-scheduler.kubeconfig


# 创建kube-scheduler的配置文件
```shell
tee kube-scheduler.conf << 'EOF'
KUBE_SCHEDULER_OPTS="--address=127.0.0.1 \
--kubeconfig=/etc/kubernetes/kube-scheduler.kubeconfig \
--leader-elect=true \
--alsologtostderr=true \
--logtostderr=false \
--log-dir=/var/log/kubernetes \
--v=2"
EOF

# 创建kube-scheduler服务启动文件
tee kube-scheduler.service << 'EOF'
[Unit]
Description=Kubernetes Scheduler
Documentation=https://github.com/kubernetes/kubernetes

[Service]
EnvironmentFile=-/etc/kubernetes/kube-scheduler.conf
ExecStart=/usr/local/bin/kube-scheduler $KUBE_SCHEDULER_OPTS
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

# 拷贝文件到对应位置，同时拷贝配置文件到master2
cp kube-scheduler*.pem /etc/kubernetes/ssl/
cp kube-scheduler.kubeconfig /etc/kubernetes/
cp kube-scheduler.conf /etc/kubernetes/
cp kube-scheduler.service /usr/lib/systemd/system/
rsync -vaz kube-scheduler*.pem k8s-master2:/etc/kubernetes/ssl/
rsync -vaz kube-scheduler.kubeconfig kube-scheduler.conf k8s-master2:/etc/kubernetes/
rsync -vaz kube-scheduler.service k8s-master2:/usr/lib/systemd/system/

```

## 8.2. master节点启动kube-schudler组件

```shell
# master节点启动服务
systemctl daemon-reload &&  systemctl enable kube-scheduler && systemctl start kube-scheduler && systemctl status kube-scheduler

# 查看各个组件的状态
kubectl get componentstatuses
#Warning: v1 ComponentStatus is deprecated in v1.19+
#NAME                 STATUS    MESSAGE             ERROR
#scheduler            Healthy   ok
#controller-manager   Healthy   ok
#etcd-1               Healthy   {"health":"true"}
#etcd-0               Healthy   {"health":"true"}

```

# 9. 部署kubelet组
## 9.1. 下载依赖的镜像

手动下载`pause`，以及`coredns`镜像，并重新打tag
<font size=4 color=Red>**注意：这里需要在各个`Node`节点上执行，因为`kubelet`是运行在各个`node`上的**</font>
```shell
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.2
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.2 k8s.gcr.io/pause:3.2
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.2
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.7.0
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.7.0 k8s.gcr.io/coredns:1.7.0
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.7.0
```

## 9.2. 准备kubelet配置文件

- <font size=4 color=Red>**注意：由于是kubelet配置文件，而kubelet是运行在Node节点上的，因此这里的address地址为node节点地址**</font>

```shell
cd /data/work

# 创建kubelet-bootstrap.kubeconfig
export BOOTSTRAP_TOKEN=$(awk -F "," '{print $1}' /etc/kubernetes/token.csv)
rm -rf kubelet-bootstrap.kubeconfig
kubectl config set-cluster kubernetes --certificate-authority=ca.pem --embed-certs=true --server=https://192.168.11.71:6443 --kubeconfig=kubelet-bootstrap.kubeconfig
kubectl config set-credentials kubelet-bootstrap --token=${BOOTSTRAP_TOKEN} --kubeconfig=kubelet-bootstrap.kubeconfig
kubectl config set-context default --cluster=kubernetes --user=kubelet-bootstrap --kubeconfig=kubelet-bootstrap.kubeconfig
kubectl config use-context default --kubeconfig=kubelet-bootstrap.kubeconfig
kubectl create clusterrolebinding kubelet-bootstrap --clusterrole=system:node-bootstrapper --user=kubelet-bootstrap

# 创建kubelet配置文件 address需要换成node的地址
tee kubelet.json << 'EOF'
{
  "kind": "KubeletConfiguration",
  "apiVersion": "kubelet.config.k8s.io/v1beta1",
  "authentication": {
    "x509": {
      "clientCAFile": "/etc/kubernetes/ssl/ca.pem"
    },
    "webhook": {
      "enabled": true,
      "cacheTTL": "2m0s"
    },
    "anonymous": {
      "enabled": false
    }
  },
  "authorization": {
    "mode": "Webhook",
    "webhook": {
      "cacheAuthorizedTTL": "5m0s",
      "cacheUnauthorizedTTL": "30s"
    }
  },
  "address": "192.168.11.73",
  "port": 10250,
  "readOnlyPort": 10255,
  "cgroupDriver": "systemd",
  "hairpinMode": "promiscuous-bridge",
  "serializeImagePulls": false,
  "featureGates": {
    "RotateKubeletServerCertificate": true
  },
  "clusterDomain": "cluster.local.",
  "clusterDNS": ["10.255.0.2"]
}
EOF

# 创建Kubelet服务启动文件
tee kubelet.service << 'EOF'
[Unit]
Description=Kubernetes Kubelet
Documentation=https://github.com/kubernetes/kubernetes
After=docker.service
Requires=docker.service
[Service]
WorkingDirectory=/var/lib/kubelet
ExecStart=/usr/local/bin/kubelet \
 --bootstrap-kubeconfig=/etc/kubernetes/kubelet-bootstrap.kubeconfig \
 --cert-dir=/etc/kubernetes/ssl \
 --kubeconfig=/etc/kubernetes/kubelet.kubeconfig \
 --config=/etc/kubernetes/kubelet.json \
 --network-plugin=cni \
 --pod-infra-container-image=k8s.gcr.io/pause:3.2 \
 --alsologtostderr=true \
 --logtostderr=false \
 --log-dir=/var/log/kubernetes \
 --v=2
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF


# 拷贝kubelet证书以及配置到node节点上
mkdir /etc/kubernetes/ssl -p

# 在master1上执行
scp kubelet-bootstrap.kubeconfig kubelet.json k8s-node1:/etc/kubernetes/
scp ca.pem k8s-node1:/etc/kubernetes/ssl/
scp  kubelet.service k8s-node1:/usr/lib/systemd/system/

```


## 9.3. 在node上启动kubelet服务

```shell
mkdir /var/lib/kubelet
mkdir /var/log/kubernetes
systemctl daemon-reload && systemctl enable kubelet &&  systemctl start kubelet && systemctl status kubelet

```

# 10. 部署kube-proxy组件
## 10.1. 准备配置文件

```shell
cd /data/work

tee kube-proxy-csr.json << 'EOF'
{
  "CN": "system:kube-proxy",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CN",
      "ST": "Hubei",
      "L": "Wuhan",
      "O": "k8s",
      "OU": "system"
    }
  ]
}
EOF

# 生成证书
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-proxy-csr.json | cfssljson -bare kube-proxy

创建kubeconfig配置文件
kubectl config set-cluster kubernetes --certificate-authority=ca.pem --embed-certs=true --server=https://192.168.11.71:6443 --kubeconfig=kube-proxy.kubeconfig
kubectl config set-credentials kube-proxy --client-certificate=kube-proxy.pem --client-key=kube-proxy-key.pem --embed-certs=true --kubeconfig=kube-proxy.kubeconfig
kubectl config set-context default --cluster=kubernetes --user=kube-proxy --kubeconfig=kube-proxy.kubeconfig
kubectl config use-context default --kubeconfig=kube-proxy.kubeconfig

# 创建kube-proxy配置文件
tee kube-proxy.yaml << 'EOF'
apiVersion: kubeproxy.config.k8s.io/v1alpha1
bindAddress: 192.168.11.73
clientConnection:
  kubeconfig: /etc/kubernetes/kube-proxy.kubeconfig
clusterCIDR: 192.168.11.0/24
healthzBindAddress: 192.168.11.73:10256
kind: KubeProxyConfiguration
metricsBindAddress: 192.168.11.73:10249
mode: "ipvs"
EOF

# 创建服务启动文件
tee kube-proxy.service << 'EOF'
[Unit]
Description=Kubernetes Kube-Proxy Server
Documentation=https://github.com/kubernetes/kubernetes
After=network.target

[Service]
WorkingDirectory=/var/lib/kube-proxy
ExecStart=/usr/local/bin/kube-proxy \
  --config=/etc/kubernetes/kube-proxy.yaml \
  --alsologtostderr=true \
  --logtostderr=false \
  --log-dir=/var/log/kubernetes \
  --v=2
Restart=on-failure
RestartSec=5
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF

# 拷贝配置文件到node节点上
scp kube-proxy.kubeconfig kube-proxy.yaml k8s-node1:/etc/kubernetes/
scp  kube-proxy.service k8s-node1:/usr/lib/systemd/system/

```

# 11. 部署calico组件

```shell
wget https://docs.projectcalico.org/v3.14/manifests/calico.yaml
kubectl apply -f calico.yaml

kubectl get pod -n kube-system

```

# 12. 部署CoreDNS组件

```shell
kubectl apply -f coredns.yaml
kubectl get pod -n kube-system

```

# 13. 查看集群状态

```shell
kubectl get nodes

```


# 14. 集群组件功能验证测试
## 14.1. 对系统用户kubernetes做授权

```shell
kubectl create clusterrolebinding kubernetes-kubectl --clusterrole=cluster-admin --user=kubernetes

```

## 14.2. 测试集群部署tomcat服务



























