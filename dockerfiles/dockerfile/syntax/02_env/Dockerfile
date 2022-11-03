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
