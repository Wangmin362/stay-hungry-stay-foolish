FROM busybox

RUN echo $HOME
RUN /bin/sh -c 'echo $HOME'
RUN sh -c  'echo $HOME'

# 由于 exec 格式的RUN命令，并不会调用shell解释执行，因此不会解析HOME环境变量，而是当作一个普通的字符串
RUN ["echo", "$HOME"]

# 可以手动调用shell解释器执行命令
RUN ["/bin/sh","-c","echo $HOME"]
