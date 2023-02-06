#!/bin/sh

# 使用docker配置文件
cp server-docker.yaml server.yaml

# 默认值初始化，让 docker 可以单独运行
CMD_ARGS=${CMD_ARGS:="server"}
RUN_ARGS=${RUN_ARGS:=""}

 # web
if [ "${CMD_ARGS}" = "server" ]; then
    HOST=${HOST:=""}
    PORT=${PORT:="50001"}
else
    echo "app run fail..."
fi

# 环境变量与配置文件显示
env
cat ${CMD_ARGS}.yaml

# 容器启动命令
./app ${CMD_ARGS} ${RUN_ARGS}
