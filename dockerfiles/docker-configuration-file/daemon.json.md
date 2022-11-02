```sh
cat >> /etc/docker/daemon.json <<EOF
{
    "registry-mirrors": ["https://registry.docker-cn.com"],
    "insecure-registries": ["172.30.3.149"],
    "exec-opts": ["native.cgroupdriver=systemd"],
    "log-driver": "json-file",
    "log-opts": {
        "max-size": "100m"
    },
    "storage-driver": "overlay2"
}
EOF
```
```shell
cat >> /etc/docker/daemon.json <<EOF
{
    "registry-mirrors": ["https://hub-mirror.c.163.com","https://mirror.baidubce.com"],
    "insecure-registries": ["172.30.3.149"],
    "exec-opts": ["native.cgroupdriver=systemd"],
    "log-driver": "json-file",
    "log-opts": {
        "max-size": "100m"
    },
    "storage-driver": "overlay2"
}
sh
```
