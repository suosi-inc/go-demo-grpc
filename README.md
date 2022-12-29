# go-demo-grpc

## 准备

> protoc.exe，windows平台编译受限，很难自己手动编译，直接下载一个，地址：https://github.com/protocolbuffers/protobuf

> 编译代码 `protoc ./protobuf/*.proto --proto_path="./protobuf" --go_out="." --go-grpc_out="."`

## 设计

- [x] 服务端-客户端通信
- [x] 项目结构分层
- [ ] 加入配置文件
- [ ] 加入 di
- [ ] 多命令入口
  - [x] ~~多服务入口, 配置麻烦, git提交无法维护, CI麻烦~~
- [ ] go mod 父子级
- [ ] client 放在 di 中
- [ ] ci 部署

### 项目结构

- client: 客户端文件, 属于公用的, 可提供给其他项目
- cmd: 命令行入口
- config: 配置文件
- protobuf: 定义的消息, 属于公用的, 可提供给其他项目
- internal
  - pkg: 公用包
    - log
    - di
  - rpc: rpc 服务
    - service: 服务
    - msg: 错误信息



