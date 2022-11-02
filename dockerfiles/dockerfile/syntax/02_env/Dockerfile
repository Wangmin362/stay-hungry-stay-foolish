FROM busybox
ENV FOO=/bar
# 如果定义了FOO环境变量，就使用mmmd替换FOO环境变量
RUN echo 111${FOO:+mmmd}111 # RUN echo 111mmmd111
# 如果没有定义FOO环境变量，就使用sdjj替换FOO环境变量
RUN echo 111${FOO:-sdjj}222 # RUN echo 111/bar222
