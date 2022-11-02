
### Dockerfile注释

```dockerfile
FROM busybox

RUN echo hello \
    world

# dockerfile中的注释没有任何影响，这两种写法是等价的
# docker build在构建的过程中会把注释直接删除，然后再进行构建
# dockerfile中的注释必须单独一样，注释前不能有docker指令，否则无法正常进行构建
RUN echo hello \
    # sdfsfsdf
    world

# 下面的写法是有问题的，注释必须单独一行，注释不能放在正常语句的后面
#RUN echo hello \
#    dockerfile # sdfsfsdf
#    world
```
