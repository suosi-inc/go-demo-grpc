# go-demo-grpc

> 服务端启动, 加参数 `server`

> 客户端启动, 加参数 `client user:get`

## 开发准备

> protoc.exe，windows平台编译受限，很难自己手动编译，直接下载一个，地址：`https://github.com/protocolbuffers/protobuf`

> protoc-gen-go, 使用命令 `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

> 编译代码 `protoc ./protobuf/*.proto --proto_path="./protobuf" --go_out="." --go-grpc_out="."`

## TODO

- [x] 服务端-客户端通信
  - [x] 身份校验(appId, appKey模式)
- [x] 项目结构分层
- [x] 加入配置文件
- [x] 加入 di
- [x] 多命令入口
  - [x] ~~多服务入口, 配置麻烦, git提交无法维护, CI麻烦~~
- [ ] go mod 父子级(解决 protobuf 之间的依赖)
  - [ ] 线上的 protobuf 如何交互
- [x] client 放在 di 中
  - [ ] 连接池
  - [ ] 服务端引用其他的服务
- [x] ci 部署
- [ ] 存储介质交互
  - [ ] mysql
  - [ ] es
  - [ ] redis
- [ ] 错误日志入库
  - [ ] 区分系统
  - [ ] 区分节点
- [ ] 性能分析工具

### 项目结构

- cmd: 命令行入口
- config: 配置文件
- protobuf: 定义的消息, 属于公用的, 可提供给其他项目
- internal
  - pkg: 公用包
    - log
    - di
    - setup: 有公用的部分, 因此放在 pkg
  - server: rpc 服务
    - service: 服务
    - msg: 错误信息
    - config: 配置
  - client: 单独的客户端sdk, 依赖于 protobuf
    - config: 配置
    - command: 实例代码, 无实际意义

## protobuf 语法

### service

> 提供模板服务, 入参和出参均为 message

```
service UserService {
  rpc Get(GetUserRequest) returns (GetUserResponse) {};
  rpc Search(SearchUserRequest) returns (SearchUserResponse) {};
  rpc Add(AddUserRequest) returns (AddUserResponse) {};
  rpc Edit(EditUserRequest) returns (EditUserResponse) {};
  rpc Remove(RemoveUserRequest) returns (RemoveUserResponse) {};
}
```

### message

> 通信的基础介质

```
// 格式: `类型 字段名 = <同级中不重复的数字>;`
// message 可以嵌套
// 引用外部的 protobuf, 使用 import, 对于非项目中的, 路径需要写全称
// optional 表示可选参数
// repeated 表示这种类型的列表

import "google/protobuf/timestamp.proto";

message SearchUserRequest {
  message CreatedTimeRange {
    google.protobuf.Timestamp start = 1;
    google.protobuf.Timestamp end = 2;
  }

  optional string name = 1;
  optional UserStatus status = 2;
  optional CreatedTimeRange createdTimeRange = 3;
  repeated string menuList = 4;
  int64 page = 5;
  int64 size = 6;
}
```


