
#### 基本示例代码

```dockerfile
FROM busybox

ENV NAME=DAVID
LABEL K1=V1 K2=V2 K3=V3
LABEL kk1=vv1 \
      kk2=vv2

# 这种写法也会直接解析变量
LABEL NAME1=$NAME
# 双引号可以正常解析环境变量
LABEL NAME2="$NAME"
# 单引号则不会解析环境变量，而是会当成一个普通的字符串
LABEL NAME3='$NAME'
```

```shell
root@remote-code:~# docker image inspect fd6b69b8757f
[
    {
        "Id": "sha256:fd6b69b8757fc1cb7a40942d36bb5c55512b5c2fdc722cca9ecfd4f85e7e69a9",
        "RepoTags": [],
        "RepoDigests": [],
        "Parent": "sha256:6903c5db3847ef5880a78b1d4b54c9166b70bd0721090991f1f531a1abda6e79",
        "Comment": "",
        "Created": "2022-11-03T01:35:52.191762956Z",
        "Container": "7d00ea5837bd198ad53b1ba1ba0e2434c43a96d81e910de3eb13d392294e6c2c",
        "ContainerConfig": {
            "Hostname": "7d00ea5837bd",
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
                "NAME=DAVID"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "#(nop) ",
                "LABEL NAME3=$NAME"
            ],
            "Image": "sha256:6903c5db3847ef5880a78b1d4b54c9166b70bd0721090991f1f531a1abda6e79",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {
                "K1": "V1",
                "K2": "V2",
                "K3": "V3",
                "NAME1": "DAVID",
                "NAME2": "DAVID",
                "NAME3": "$NAME",
                "kk1": "vv1",
                "kk2": "vv2"
            }
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
                "NAME=DAVID"
            ],
            "Cmd": [
                "sh"
            ],
            "Image": "sha256:6903c5db3847ef5880a78b1d4b54c9166b70bd0721090991f1f531a1abda6e79",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {
                "K1": "V1",
                "K2": "V2",
                "K3": "V3",
                "NAME1": "DAVID",
                "NAME2": "DAVID",
                "NAME3": "$NAME",
                "kk1": "vv1",
                "kk2": "vv2"
            }
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
