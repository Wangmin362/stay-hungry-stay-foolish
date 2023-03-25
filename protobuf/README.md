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

## 字段编号分配原则

- 1、为了保持版本的兼容性，一旦字段的编号开始使用，就不应该更改编号
  - 主要原因是，gRPC的消息序列化的时候会使用编号代替字段的key，如果数据生产方把gRpc数据放入到kafka，并且在消息还没有被消费之前，我们把字段的编号换了，此时数据的消费方就不能正确的序列化数据，因此一旦字段的编号确定，不建议修改字段编号
- 2、`1-15`号编号采用一个字节编号，而`16-2048`编号采用2个字节编码，因此建议把常用字段放到`1-15`编号编码

## 枚举

- 1、枚举必须有一个零值，建议使用`XXX_UNSPECIFIED = 0`来作为枚举的第一个值

## 其它proto依赖导入


# Protoc工具
## 参数

- `-I --proto_path <dir>`： 指定需要编译的`proto`文件，如果不指定，那么默认编译当前目录下的所有`protoc`文件
- 

## golang开发

