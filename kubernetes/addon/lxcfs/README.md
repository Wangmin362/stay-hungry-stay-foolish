# LXCFS 问题汇总

<font size=5 color=Red>用于解决K8S pod获取CPU以及内存资源为宿主机资源的问题</font>
- 参考视频
  - [K8S资源视图隔离](https://www.huweihuang.com/kubernetes-notes/resource/lxcfs/lxcfs.html)

## 使用建议

- 1、应用程序的镜像的基础镜像的版本应该尽可能使使用5.10.x以上的内核版本，否则在应用程序中执行某些命令，依然无法正确获取容器的资源使用

## 角色

角色 | 用途 | 其他 
----|------|----
[lxcfs](https://linuxcontainers.org/lxcfs/) | lxcfs 是一个开源的 FUSE 的用户态文件系统，用来实现来支持LXC容器| 容器内部获得正确的限制的 cpu、内存等信息
[lxcfs-admission-webhook](https://github.com/denverdino/lxcfs-admission-webhook)| 按需拦截 pod 的创建， patch lxcfs 的 volume | 代码有些岁数了，有一定的优化空间
PODS | 需要获取正确数据的程序| nginx、jave 等
---

## 测试的 pod

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-lxcfs
spec:
  replicas: 1
  selector:
    matchLabels:
      app: testlxcfs
  template:
    metadata:
      labels:
        app: testlxcfs
    spec:
      containers:
        - name: alpine
          #image: 172.30.3.150/devops/ubuntu:20.04
          image: 172.30.3.150/devops/alpine:3.16
          command: ["/bin/sh"]
          args: ["-c", "sleep 1d"]
          imagePullPolicy: Always
          resources:
            requests:
              memory: "256Mi"
              cpu: "0.2"
            limits:
              memory: "1024Mi"
              cpu: "0.5"

```
---


## 正常情况

```shell
### 1 cpu
# kubectl exec -it test-lxcfs-5b449777dd-fl7m9 -- cat /proc/cpuinfo | grep processor | wc -l
1
### 1G memory
# kubectl exec -it test-lxcfs-5b449777dd-fl7m9 -- cat /proc/meminfo  | grep MemTotal
MemTotal:        1048576 kB

```
---


## 问题

> 小概率发生

### lxcfs 异常
- 已经正常运行的 pod 不能获取 cpu、内存，即使 lxcfs 恢复也是一样的情况
```shell
### lxcfs 异常
# kubectl exec -it test-lxcfs-5b449777dd-fl7m9 -- cat /proc/cpuinfo 
cat: can't open '/proc/cpuinfo': Socket not connected
command terminated with exit code 
# kubectl exec -it test-lxcfs-5b449777dd-fl7m9 -- cat /proc/meminfo 
cat: can't open '/proc/meminfo': Socket not connected
command terminated with exit code 1
### lxcfs 恢复正常
# kubectl exec -it test-lxcfs-5b449777dd-fl7m9 -- cat /proc/cpuinfo 
cat: can't open '/proc/cpuinfo': Socket not connected
command terminated with exit code 1
# kubectl exec -it test-lxcfs-5b449777dd-fl7m9 -- cat /proc/meminfo
cat: can't open '/proc/meminfo': Socket not connected
command terminated with exit code 1
### 重建 pod
# kubectl exec -it test-lxcfs-5b449777dd-bzgqm -- cat /proc/cpuinfo | grep processor | wc -l
1
# kubectl exec -it test-lxcfs-5b449777dd-bzgqm -- cat /proc/meminfo  | grep MemTotal
MemTotal:        1048576 kB

```

- 新建 pod 不能正常启动，直到 lxcfs 恢复

```shell
# kubectl get pods -w 
NAME                          READY   STATUS    RESTARTS   AGE
test-lxcfs-5b449777dd-r5g68   0/1     Pending   0          0s
test-lxcfs-5b449777dd-r5g68   0/1     Pending   0          0s
test-lxcfs-5b449777dd-r5g68   0/1     Init:0/1   0          0s
test-lxcfs-5b449777dd-r5g68   0/1     Init:0/1   0          1s
test-lxcfs-5b449777dd-r5g68   0/1     PodInitializing   0          2s
test-lxcfs-5b449777dd-r5g68   0/1     RunContainerError   0          4s
test-lxcfs-5b449777dd-r5g68   0/1     RunContainerError   1          5s
test-lxcfs-5b449777dd-r5g68   0/1     CrashLoopBackOff    1          17s
test-lxcfs-5b449777dd-r5g68   0/1     RunContainerError   2          38s
test-lxcfs-5b449777dd-r5g68   0/1     RunContainerError   3          51s
test-lxcfs-5b449777dd-r5g68   0/1     CrashLoopBackOff    3          52s
test-lxcfs-5b449777dd-r5g68   1/1     Running             4          90s

### describe 信息
Error: failed to start container "alpine": Error response from daemon: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:76: mounting "/var/lib/lxcfs/proc/loadavg" to rootfs at "/proc/loadavg" caused: mount through procfd: not a directory: unknown: Are you trying to mount a directory onto a file (or vice-versa)? Check if the specified host path exists and is the expected type

### Running 之后能够获取正确的信息
# kubectl exec -it test-lxcfs-5b449777dd-r5g68 -- cat /proc/cpuinfo | grep processor | wc -l
1
# kubectl exec -it test-lxcfs-5b449777dd-r5g68 -- cat /proc/meminfo  | grep MemTotal
MemTotal:        1048576 kB
```

### lxcfs-admission-webhook 异常
> pod 将不会被 path lxcfs 相关的 volume，pod 获取的信息和宿主机一致；需要 lxcfs-admission-webhook 恢复正常只用重建 pod

### lxcfs、lxcfs-admission-webhook 启动优先级
> 需要确保启动优先级高于其他 pod(服务器断电，重启等情况)
