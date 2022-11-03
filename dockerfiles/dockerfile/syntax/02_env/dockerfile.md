### ENV指令

- 可以使用`\`对`$`符号进行转义，从而使得特殊字符串不被`docker build`替换为`ENV`环境变量
- 环境变量可以使用`$ENV`，或者`${ENV}`的方式引用，后者主要用于解决环境变量有空格的情况
- `ENV`声明的环境变量可以通过`docker run -e K1=V1`的方式修改
  - 注意：除了`CMD`以及`ENTRYPOINT`指令会再运行的时候去取环境变量，其余指令由于再`docker build`的过程中就已经确定了值，是无法更改的
    - 想要在`docker run`的过程中修改变量值，应该使用`ARG`指令
- `ENV`指令可以被使用再如下指令中：
  - `ADD`
  - `COPY`
  - `ENV`
  - `EXPOSE`
  - `FROM`
  - `LABEL`
  - `STOPSIGNAL`
  - `USER`
  - `VOLUME`
  - `WORKDIR`
  - `ONBUILD`

#### 转义对于ENV变量的引用，使其为一个普通字符串

```dockerfile
FROM busybox
ENV FOO=/bar
# 如果定义了FOO环境变量，就使用mmmd替换FOO环境变量
RUN echo 111${FOO:+mmmd}111 # RUN echo 111mmmd111
# 如果没有定义FOO环境变量，就使用sdjj替换FOO环境变量
RUN echo 111${FOO:-sdjj}222 # RUN echo 111/bar222

# 可以再一个ENV指令中定义多个环境变量
ENV K1=V1 K2=V2 K3=V3

# 环境变量也可以使用这种方式定义，但是这种方式一次只能定义一个环境变量
ENV K4 V4

RUN echo $K1 > abc.txt

CMD ["/bin/sh", "-c", "ping $K1"]
```

#### ENV的默认值功能

```dockerfile
FROM busybox
ENV FOO=/bar
# 如果定义了FOO环境变量，就使用mmmd替换FOO环境变量
RUN echo 111${FOO:+mmmd}111 # RUN echo 111mmmd111
# 如果没有定义FOO环境变量，就使用sdjj替换FOO环境变量
RUN echo 111${FOO:-sdjj}222 # RUN echo 111/bar222
```

#### 查看镜像的环境变量

```dockerfile
FROM busybox
ENV FOO=/bar
# 如果定义了FOO环境变量，就使用mmmd替换FOO环境变量
RUN echo 111${FOO:+mmmd}111 # RUN echo 111mmmd111
# 如果没有定义FOO环境变量，就使用sdjj替换FOO环境变量
RUN echo 111${FOO:-sdjj}222 # RUN echo 111/bar222

# 可以再一个ENV指令中定义多个环境变量
ENV K1=V1 K2=V2 K3=V3

CMD ["/bin/sh", "-c", "ping baidu.com"]
```

```shell
root@remote-code:~# docker image inspect f53a7b0eb20f
[
    {
        "Id": "sha256:f53a7b0eb20f9ae40fae29efa43c6c59287e8ba053ea63a113ec1eb66d7fb606",
        "RepoTags": [],
        "RepoDigests": [],
        "Parent": "sha256:59ad81a436b70f6bfb4f494e555695bbd295f5c504be3e5f86e9ed8ce109df40",
        "Comment": "",
        "Created": "2022-11-03T06:06:57.56858298Z",
        "Container": "316865a714e77424041f4a8eb42a9e1095b98890979f7dc9798646c210d9ae3e",
        "ContainerConfig": {
            "Hostname": "316865a714e7",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "FOO=/bar",
                "K1=V1",
                "K2=V2",
                "K3=V3"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "#(nop) ",
                "CMD [\"/bin/sh\" \"-c\" \"ping baidu.com\"]"
            ],
            "Image": "sha256:59ad81a436b70f6bfb4f494e555695bbd295f5c504be3e5f86e9ed8ce109df40",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {}
        },
        "DockerVersion": "20.10.7",
        "Author": "",
        "Config": {
            "Hostname": "",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "FOO=/bar",
                "K1=V1",
                "K2=V2",
                "K3=V3"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "ping baidu.com"
            ],
            "Image": "sha256:59ad81a436b70f6bfb4f494e555695bbd295f5c504be3e5f86e9ed8ce109df40",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": null
        },
        "Architecture": "amd64",
        "Os": "linux",
        "Size": 1239820,
        "VirtualSize": 1239820,
        "GraphDriver": {
            "Data": {
                "MergedDir": "/var/lib/docker/overlay2/acd0f7f830d6add8b11df23f71dd78cded2b16672defe6dee7b0548c8e9568b5/merged",
                "UpperDir": "/var/lib/docker/overlay2/acd0f7f830d6add8b11df23f71dd78cded2b16672defe6dee7b0548c8e9568b5/diff",
                "WorkDir": "/var/lib/docker/overlay2/acd0f7f830d6add8b11df23f71dd78cded2b16672defe6dee7b0548c8e9568b5/work"
            },
            "Name": "overlay2"
        },
        "RootFS": {
            "Type": "layers",
            "Layers": [
                "sha256:01fd6df81c8ec7dd24bbbd72342671f41813f992999a3471b9d9cbc44ad88374"
            ]
        },
        "Metadata": {
            "LastTagTime": "0001-01-01T00:00:00Z"
        }
    }
]
```

#### 通过docker run -e K1=UUUUIIII修改环境变量

```dockerfile
FROM busybox
ENV FOO=/bar
# 如果定义了FOO环境变量，就使用mmmd替换FOO环境变量
RUN echo 111${FOO:+mmmd}111 # RUN echo 111mmmd111
# 如果没有定义FOO环境变量，就使用sdjj替换FOO环境变量
RUN echo 111${FOO:-sdjj}222 # RUN echo 111/bar222

# 可以再一个ENV指令中定义多个环境变量
ENV K1=V1 K2=V2 K3=V3

RUN echo $K1 > abc.txt

CMD ["/bin/sh", "-c", "ping $K1"]
```

```shell
root@remote-code:~# docker inspect nervous_gagarin
[
    {
        "Id": "61e3e97f3f5c5a42949ce34695aa2ed6eb7a4df427390d6b1f6bdb663c70aca4",
        "Created": "2022-11-03T06:15:26.039268965Z",
        "Path": "/bin/sh",
        "Args": [
            "-c",
            "ping $K1"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 2489548,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2022-11-03T06:15:26.408498928Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:cad766a0b2bd319d33c4a73807dabe328e35b42043923e49a523788a8542b0f5",
        "ResolvConfPath": "/var/lib/docker/containers/61e3e97f3f5c5a42949ce34695aa2ed6eb7a4df427390d6b1f6bdb663c70aca4/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/61e3e97f3f5c5a42949ce34695aa2ed6eb7a4df427390d6b1f6bdb663c70aca4/hostname",
        "HostsPath": "/var/lib/docker/containers/61e3e97f3f5c5a42949ce34695aa2ed6eb7a4df427390d6b1f6bdb663c70aca4/hosts",
        "LogPath": "/var/lib/docker/containers/61e3e97f3f5c5a42949ce34695aa2ed6eb7a4df427390d6b1f6bdb663c70aca4/61e3e97f3f5c5a42949ce34695aa2ed6eb7a4df427390d6b1f6bdb663c70aca4-json.log",
        "Name": "/nervous_gagarin",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {},
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": true,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "CgroupnsMode": "host",
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/9216cae36ef583f6b2169ede00cc22fff47ca8061ba892e5f7b0f1bb45db5188-init/diff:/var/lib/docker/overlay2/09c6b38f7d02f23bd02bfdf82afd409b9d6400dbf3a8a64bf4f5c374aad56fc1/diff:/var/lib/docker/overlay2/acd0f7f830d6add8b11df23f71dd78cded2b16672defe6dee7b0548c8e9568b5/diff",
                "MergedDir": "/var/lib/docker/overlay2/9216cae36ef583f6b2169ede00cc22fff47ca8061ba892e5f7b0f1bb45db5188/merged",
                "UpperDir": "/var/lib/docker/overlay2/9216cae36ef583f6b2169ede00cc22fff47ca8061ba892e5f7b0f1bb45db5188/diff",
                "WorkDir": "/var/lib/docker/overlay2/9216cae36ef583f6b2169ede00cc22fff47ca8061ba892e5f7b0f1bb45db5188/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "61e3e97f3f5c",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "K1=BAIDU.COM",
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "FOO=/bar",
                "K2=V2",
                "K3=V3"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "ping $K1"
            ],
            "Image": "cad766a0b2bd",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {}
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "d7472a57144ad3f228a451ea705d562f25e600be3b14aaf19a7236074b123e51",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {},
            "SandboxKey": "/var/run/docker/netns/d7472a57144a",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "6738017d6535d84131012eea4c47e2472ff9654a293d2b83404836eae29b4a61",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.5",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:05",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "e98e42084af6ecf413d2e4e4587280405fa4338f40307a2df0bee4047ded3d3e",
                    "EndpointID": "6738017d6535d84131012eea4c47e2472ff9654a293d2b83404836eae29b4a61",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.5",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:05",
                    "DriverOpts": null
                }
            }
        }
    }
]

root@remote-code:~# docker run -d --rm -e K1=BAIDU.COM cad766a0b2bd
61e3e97f3f5c5a42949ce34695aa2ed6eb7a4df427390d6b1f6bdb663c70aca4
root@remote-code:~# docker ps
CONTAINER ID   IMAGE               COMMAND                  CREATED         STATUS         PORTS                                                                                                                                                                                                                                                                                                                                                                                                                                      NAMES
61e3e97f3f5c   cad766a0b2bd        "/bin/sh -c 'ping $K…"   2 seconds ago   Up 2 seconds 


root@remote-code:~# docker logs -f nervous_gagarin
PING BAIDU.COM (110.242.68.66): 56 data bytes # 这里是再运行时取了真正的K1环境变量，实际上这里并没有再构建过程中执行，只有再docker run的时候才会被执行
64 bytes from 110.242.68.66: seq=0 ttl=50 time=43.316 ms
64 bytes from 110.242.68.66: seq=1 ttl=50 time=40.355 ms
64 bytes from 110.242.68.66: seq=2 ttl=50 time=41.482 ms
64 bytes from 110.242.68.66: seq=3 ttl=50 time=41.357 ms
64 bytes from 110.242.68.66: seq=4 ttl=50 time=40.864 ms
64 bytes from 110.242.68.66: seq=5 ttl=50 time=40.455 ms
^C
root@remote-code:~# docker exec nervous_gagarin cat abc.txt
V1  # 可以看到，这里实际上就是V1，这是再docker build过程中就决定了的，无法更改
```
