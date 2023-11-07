#!/bin/bash
# set -euo pipefail

# 定义容器名称
CONTAINERS=("glee1" "glee2" "glee3")

# 清空etcd数据目录
for container in ${CONTAINERS[@]}
do 
  docker exec $container rm -rf /expr/data/*
done