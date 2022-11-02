
### FROM
- 语法：
  - `FROM [--platform=<platform>] <image> [AS <name>]`
  - `FROM [--platform=<platform>] <image>[:<tag>] [AS <name>]`
  - `FROM [--platform=<platform>] <image>[@<digest>] [AS <name>]`
- `AS <name>`是`Dockerfile`多阶段构建的语法
- `FROM`指令用于指定当前镜像是基于哪个`ROOTFS`构建的，一个有效的`Dockerfile`必须以`FROM`指令开始
  - 注意：`ARG`指定是可以出现在`FROM`指定之前的，用于构建`FROM`指令的参数
- 一个`Dockerfile`中可以出现多个`FROM`镜像，以便一次性创建多个镜像
  - 注意：`docker build`在遇到`FROM`指令的时候会把之前的构建状态清空
- 可以通过 `--platform=$BUILDPLATFORM`来指定镜像可以运行的平台，默认情况下会被设置为构建镜像的平台
  - 可选的值有：`linux/amd64, linux/arm64, windows/amd64`
- 

#### FROM --platform举例

```dockerfile

```
