## kubernetes v1.23.10 一个Master,一个Node

[参考文档](https://www.cnblogs.com/fengdejiyixx/p/16576021.html)
> 安装单节点集群的目的: 方便调试K8S源码，理解其设计思路，因此多节点集群没有意义，反而增加了调试的复杂度

### 规划

- K8S环境规划：
  - Pod网段：`10.0.0.0/16`
  - Service网段：`10.255.0.0/16`
- 实验环境规划：
  - 操作系统：`Centso7.9`
  - 配置：1GB, 2vCpu, 100G硬盘

| 集群角色 |      IP       |    主机名    |                         安装组件                         |
| ------- | ------------- | ----------- | ------------------------------------------------------- |
| 控制节点 | 192.168.11.71 | k8s-master1 | api-server, controller-manager, scheduler, etcd, docker |
| 工作节点 | 192.168.11.72 | k8s-node1   | kubelet, kube-proxy, docker, calico, coredns            |

### 环境准备
### 修改Ip地址 todo 改为前缀匹配
sed -i 's/IPADDR=192.168.11.11/IPADDR=192.168.11.72/g' /etc/sysconfig/network-scripts/ifcfg-ens33
hostnamectl set-hostname node2
reboot

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
crontab -e
* */1 * * * /usr/sbin/ntpdate	cn.pool.ntp.org
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


tee /etc/docker/daemon.json << 'EOF'
{
"registry-mirrors":["https://rsbud4vc.mirror.aliyuncs.com","https://registry.docker-cn.com","https://docker.mirrors.ustc.edu.cn","https://dockerhub.azk8s.cn","http://hub-mirror.c.163.com","http://qtid6917.mirror.aliyuncs.com", "https://rncxm540.mirror.aliyuncs.com"],
"exec-opts": ["native.cgroupdriver=systemd"]
}
EOF
systemctl daemon-reload && systemctl restart docker && systemctl status docker


# master1上操作
mkdir /data/work -p
cd /data/work/
wget -O cfssl_linux-amd64 https://ghproxy.com/https://github.com/cloudflare/cfssl/releases/download/v1.6.3/cfssl_1.6.3_linux_amd64
wget -O cfssljson_linux-amd64 https://ghproxy.com/https://github.com/cloudflare/cfssl/releases/download/v1.6.3/cfssljson_1.6.3_linux_amd64
wget -O cfssl-certinfo_linux-amd64 https://ghproxy.com/https://github.com/cloudflare/cfssl/releases/download/v1.6.3/cfssl-certinfo_1.6.3_linux_amd64
#把文件变成可执行权限
chmod +x *
mv cfssl_linux-amd64 /usr/local/bin/cfssl
mv cfssljson_linux-amd64 /usr/local/bin/cfssljson
mv cfssl-certinfo_linux-amd64 /usr/local/bin/cfssl-certinfo

# 生成CA证书请求文件
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
cfssl gencert -initca ca-csr.json  | cfssljson -bare ca

# 生成CA证书文件
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


# 生成ETCD证书
tee etcd-csr.json << 'EOF'
{
  "CN": "etcd",
  "hosts": [
    "127.0.0.1",
    "192.168.11.71",
    "192.168.11.72",
    "192.168.11.73",
    "192.168.11.74"
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
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes etcd-csr.json | cfssljson  -bare etcd

# 部署ETCD
etcdVersion=3.4.13
wget -O etcd-linux-amd64 https://ghproxy.com/https://github.com/etcd-io/etcd/releases/download/v${etcdVersion}/etcd-v${etcdVersion}-linux-amd64.tar.gz
tar -zxvf etcd-v3.4.13-linux-amd64.tar.gz
cp -ar etcd-v3.4.13-linux-amd64/etcd* /usr/local/bin
chmod +x /usr/local/bin/etcd*

scp -r etcd-v3.4.13-linux-amd64/etcd* k8s-master2:/usr/local/bin
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

# master1 master2上执行
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

# master1 以及 master2上启动服务
systemctl daemon-reload && systemctl enable etcd.service && systemctl start etcd.service

# 查看etcd集群
ETCDCTL_API=3 && /usr/local/bin/etcdctl --write-out=table --cacert=/etc/etcd/ssl/ca.pem --cert=/etc/etcd/ssl/etcd.pem --key=/etc/etcd/ssl/etcd-key.pem --endpoints=https://192.168.11.71:2379,https://192.168.11.72:2379 endpoint health
#+----------------------------+--------+-------------+-------+
#|          ENDPOINT          | HEALTH |    TOOK     | ERROR |
#+----------------------------+--------+-------------+-------+
#| https://192.168.11.71:2379 |   true | 29.053448ms |       |
#| https://192.168.11.72:2379 |   true | 29.973864ms |       |
#+----------------------------+--------+-------------+-------+

# 下载Kubernetes二进制版本,官方地址:https://github.com/kubernetes/kubernetes/tree/master/CHANGELOG
wget https://storage.googleapis.com/kubernetes-release/release/v1.23.13/kubernetes-server-linux-amd64.tar.gz
tar zxvf kubernetes-server-linux-amd64.tar.gz
cd  kubernetes/server/bin/
cp kube-apiserver kube-controller-manager kube-scheduler kubectl /usr/local/bin/
rsync -vaz kube-apiserver kube-controller-manager kube-scheduler kubectl k8s-master2:/usr/local/bin/
scp kubelet kube-proxy k8s-node1:/usr/local/bin/
cd /data/work/
mkdir -p /etc/kubernetes/ssl
mkdir /var/log/kubernetes

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
    "192.168.11.74",
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

# 启动kube-apiserver master1 以及 master2同时执行
systemctl daemon-reload && systemctl enable kube-apiserver && systemctl start kube-apiserver

# 查看kube-apiserver启动状态
systemctl status kube-apiserver
# 查看api-server授权状态
curl --insecure https://192.168.11.71:6443/
#{
#  "kind": "Status",
#  "apiVersion": "v1",
#  "metadata": {},
#  "status": "Failure",
#  "message": "Unauthorized",
#  "reason": "Unauthorized",
#  "code": 401
#}
# 上面看到 401，这个是正常的的状态，还没认证

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

# kubectl 命令补全
yum install -y bash-completion
source /usr/share/bash-completion/bash_completion
source <(kubectl completion bash)
kubectl completion bash > ~/.kube/completion.bash.inc
source '/root/.kube/completion.bash.inc'
source $HOME/.bash_profile


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
      "192.168.11.73",
      "192.168.11.74"
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


# 启动controller-manager服务
systemctl daemon-reload  &&systemctl enable kube-controller-manager && systemctl start kube-controller-manager && systemctl status kube-controller-manager








