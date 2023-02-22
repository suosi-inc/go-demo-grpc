# 安装

```
# 通过 requirements 安装
pip install -r requirements.txt

# 直接安装
pip install grpcio
pip install grpcio-tools
```

# 编译代码

进入到 python_client 根目录, 执行 `python -m grpc_tools.protoc -I../../protobuf --python_out=./pb --pyi_out=./pb --grpc_python_out=./pb ../../protobuf/*.proto"`

# 注意

> python 生成代码之后, 有少许的包依赖会出错, 需要手动更正。原因是默认生成的代码默认包在根目录, 且未找到 python_package 的选项
