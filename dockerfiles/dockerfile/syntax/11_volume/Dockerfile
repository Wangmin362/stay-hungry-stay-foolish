
FROM busybox

RUN mkdir /myvol
RUN echo "hello world" > /myvol/greeting
# 如果在docker run的时候没有指定 -v映射外部的宿主机的目录，那么就是数据集一个临时目录映射进来，容器删除后，数据仍旧会被删除
VOLUME /myvol

CMD ping baidu.com
