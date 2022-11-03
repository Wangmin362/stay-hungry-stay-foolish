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
