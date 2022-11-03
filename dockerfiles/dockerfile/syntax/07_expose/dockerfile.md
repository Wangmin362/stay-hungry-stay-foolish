
### 作用

- `EXPOSE`命令主要用于告知镜像的使用人员此镜像需要暴露的端口
- `EXPOSE`命令的意义在于约定和协同开发，并没有真正意义上暴露端口，真正暴露端口是通过`docker run -p`的方式暴露的


```dockerfile
FROM busybox

MAINTAINER wangmin@skyguard.com.cn

CMD ["/bin/bash", "-c", "ping baidu.com"]
```

```shell
root@remote-code:~# docker image inspect 26c6351e5667
[
    {
        "Id": "sha256:26c6351e56673c6fe31377949b31b52c85e0421ab218956f67743d519bba069e",
        "RepoTags": [],
        "RepoDigests": [],
        "Parent": "sha256:be16a3b0186a99f348c50d971a289befd2b2b615fbfb2f286e834992848000d7",
        "Comment": "",
        "Created": "2022-11-03T01:42:09.376127798Z",
        "Container": "b77bfe116d6a5e5293301d2e1233d41039680077f8bc8807fb5eef305f37968d",
        "ContainerConfig": {
            "Hostname": "b77bfe116d6a",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
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
            "Image": "sha256:be16a3b0186a99f348c50d971a289befd2b2b615fbfb2f286e834992848000d7",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {}
        },
        "DockerVersion": "20.10.7",
        "Author": "wangmin@skyguard.com.cn",
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
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/bash",
                "-c",
                "ping baidu.com"
            ],
            "Image": "sha256:be16a3b0186a99f348c50d971a289befd2b2b615fbfb2f286e834992848000d7",
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
