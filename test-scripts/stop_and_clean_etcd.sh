#!/bin/bash
# set -euo pipefail

# 定义容器名称
CONTAINERS=("glee1" "glee2" "glee3")

# 停止容器中的glee-etcd进程
for container in ${CONTAINERS[@]} 
do
  docker exec $container pkill glee-etcd
done 

# 解除端口绑定
for container in ${CONTAINERS[@]}
do
  # 检查2380端口
  PID=$(docker exec $container lsof -ti :2380)
  if [ ! -z "$PID" ]; then
    docker exec $container kill $PID
  fi
  
  # 检查2379端口
  PID=$(docker exec $container lsof -ti :2379)
  if [ ! -z "$PID" ]; then  
    docker exec $container kill $PID
  fi
done

# 清空etcd数据目录
for container in ${CONTAINERS[@]}
do 
  docker exec $container rm -rf /expr/data/*
done