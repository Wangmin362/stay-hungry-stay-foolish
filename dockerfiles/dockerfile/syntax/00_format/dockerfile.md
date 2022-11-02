### Dockerfile格式说明

- 1、`Dockerfile`中的指令不区分大小写，但是官方约定使用**大写**指令，以便一眼就能看出来哪些是`Dockerfile`指令，哪些是构建参数
- 2、Dockerfile中的注释应该以`#`开头，在`docker build`的时候所有注释会被删除，然后再进行构建
- 3、`Dockerfile`中的注释不支持换行符
