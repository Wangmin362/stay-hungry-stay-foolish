# Protobuf语法

- 以下划线开头的字段会把所有的下划线删除，并使用大写的`X`代替
- 关于optional字段的讨论
  - [implementing_proto3_presence](https://github.com/protocolbuffers/protobuf/blob/main/docs/implementing_proto3_presence.md)
  - [field_presence](https://github.com/protocolbuffers/protobuf/blob/main/docs/field_presence.md)
  - optional在protoc3.12版本作为实验性质被开放出来，并且已经在protoc 3.15版本设置为默认

# 约定 && 规范
## 协议

- 1、协议应该使用`proto3`版本的协议,`proto2`版本协议太老了

## 命名规范

- 1、实体类、服务名、rpc方法应遵从驼峰命名，并且大写开头
- 2、实体类字段应该遵从小写下划线命名
- 3、实体应该使用单数
- 4、实体和服务名应该是有关联的
  - 譬如：实体为`Book`，那么服务名应该命名为：`BookService`

### RPC方法命名

- 1、对于实体的标准操作，譬如查询，列表查询，创建，更新和删除都应该使用标准方法
  - 标准方法有：`Get, List, Create, Update, Delate`