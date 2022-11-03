
### 作用

- `EXPOSE`命令主要用于告知镜像的使用人员此镜像需要暴露的端口
- `EXPOSE`命令的意义在于约定和协同开发，并没有真正意义上暴露端口，真正暴露端口是通过`docker run -p`的方式暴露的
- `EXPOSE`声明暴露的端口只有当使用`docker run -p`暴露端口才有意义
- `EXPOSE`声明暴露的端口当使用`docker run -P`随机暴露端口时，会给每个端口随机分配一个端口


```dockerfile
FROM busybox

# 基本使用方式（没有声明协议，默认就是tcp协议）
EXPOSE 8001

# 设置该端口暴露那种协议
EXPOSE 8002/tcp

# 如果一个端口，须同时暴露两种协议，那么需要声明两次
EXPOSE 8003/tcp
EXPOSE 8003/udp

CMD ["/bin/bash", "-c", "ping baidu.com"]
```

```shell
root@remote-code:~#
root@remote-code:~# docker image inspect 28239f68d1ef
[
    {
        "Id": "sha256:28239f68d1ef1c79c30fe4fa386eb34c1246e94164542765c1b754cacfd786c9",
        "RepoTags": [],
        "RepoDigests": [],
        "Parent": "sha256:4c79485597f24102d906993230ddc1f3b1fdb9161ce6ca6bd70bbf52f5e971f6",
        "Comment": "",
        "Created": "2022-11-03T05:45:54.698544351Z",
        "Container": "c2d9b1f00c8291681a7ab11314a8f175b7806552be95dcac08f48ade76399740",
        "ContainerConfig": {
            "Hostname": "c2d9b1f00c82",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "8001/tcp": {},
                "8002/tcp": {},
                "8003/tcp": {},
                "8003/udp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "#(nop) ",
                "CMD [\"/bin/bash\" \"-c\" \"ping baidu.com\"]"
            ],
            "Image": "sha256:4c79485597f24102d906993230ddc1f3b1fdb9161ce6ca6bd70bbf52f5e971f6",
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
            "ExposedPorts": {
                "8001/tcp": {},
                "8002/tcp": {},
                "8003/tcp": {},
                "8003/udp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/bash",
                "-c",
                "ping baidu.com"
            ],
            "Image": "sha256:4c79485597f24102d906993230ddc1f3b1fdb9161ce6ca6bd70bbf52f5e971f6",
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

```shell
# 使用 -P 随机暴露端口时，docker会给没有给EXPOSE暴露的端口随机分配一个端口号暴露出来
root@remote-code:~# docker run -it --rm -P -d 5f804acd6cc3
root@remote-code:~# docker inspect cranky_hoover
[
    {
        "Id": "70b324d49c04d29c87230d96d60baedcc91ad82d72e9f96fb1540dd9cd49f6e9",
        "Created": "2022-11-03T05:54:28.238975461Z",
        "Path": "/bin/sh",
        "Args": [
            "-c",
            "ping baidu.com"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 2481666,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2022-11-03T05:54:28.636542726Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:5f804acd6cc3938ae1a7e0111ba5149177916766d9f790c78d22d0f7db0d178b",
        "ResolvConfPath": "/var/lib/docker/containers/70b324d49c04d29c87230d96d60baedcc91ad82d72e9f96fb1540dd9cd49f6e9/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/70b324d49c04d29c87230d96d60baedcc91ad82d72e9f96fb1540dd9cd49f6e9/hostname",
        "HostsPath": "/var/lib/docker/containers/70b324d49c04d29c87230d96d60baedcc91ad82d72e9f96fb1540dd9cd49f6e9/hosts",
        "LogPath": "/var/lib/docker/containers/70b324d49c04d29c87230d96d60baedcc91ad82d72e9f96fb1540dd9cd49f6e9/70b324d49c04d29c87230d96d60baedcc91ad82d72e9f96fb1540dd9cd49f6e9-json.log",
        "Name": "/cranky_hoover",
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
            "PublishAllPorts": true,
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
                "LowerDir": "/var/lib/docker/overlay2/9a54e17f4e593253b25d5ded937775a10e02fd7caaba95fa94922f56ccd28aa8-init/diff:/var/lib/docker/overlay2/acd0f7f830d6add8b11df23f71dd78cded2b16672defe6dee7b0548c8e9568b5/diff",
                "MergedDir": "/var/lib/docker/overlay2/9a54e17f4e593253b25d5ded937775a10e02fd7caaba95fa94922f56ccd28aa8/merged",
                "UpperDir": "/var/lib/docker/overlay2/9a54e17f4e593253b25d5ded937775a10e02fd7caaba95fa94922f56ccd28aa8/diff",
                "WorkDir": "/var/lib/docker/overlay2/9a54e17f4e593253b25d5ded937775a10e02fd7caaba95fa94922f56ccd28aa8/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "70b324d49c04",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "8001/tcp": {},
                "8002/tcp": {},
                "8003/tcp": {},
                "8003/udp": {}
            },
            "Tty": true,
            "OpenStdin": true,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "ping baidu.com"
            ],
            "Image": "5f804acd6cc3",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {}
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "143acbf2e744cf9b770a098b5d91fdeb27b231f06e31835dc8006c06d32fea3d",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "8001/tcp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "49158"
                    },
                    {
                        "HostIp": "::",
                        "HostPort": "49158"
                    }
                ],
                "8002/tcp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "49157"
                    },
                    {
                        "HostIp": "::",
                        "HostPort": "49157"
                    }
                ],
                "8003/tcp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "49156"
                    },
                    {
                        "HostIp": "::",
                        "HostPort": "49156"
                    }
                ],
                "8003/udp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "49154"
                    },
                    {
                        "HostIp": "::",
                        "HostPort": "49154"
                    }
                ]
            },
            "SandboxKey": "/var/run/docker/netns/143acbf2e744",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "ee7149eb3347d539f5a8c76f96bfef3865c765a113b677796611dd508f4d1af8",
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
                    "EndpointID": "ee7149eb3347d539f5a8c76f96bfef3865c765a113b677796611dd508f4d1af8",
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
```
