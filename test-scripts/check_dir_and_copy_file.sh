#!/bin/bash

# 检查并在glee1中创建/expr目录和/expr/data目录
docker exec -it glee1 bash -c "if [ ! -d '/expr' ]; then mkdir /expr; fi"
docker exec -it glee1 bash -c "if [ ! -d '/expr/data' ]; then mkdir /expr/data; fi"

# 检查并在glee2中创建/expr目录和/expr/data目录
docker exec -it glee2 bash -c "if [ ! -d '/expr' ]; then mkdir /expr; fi"
docker exec -it glee2 bash -c "if [ ! -d '/expr/data' ]; then mkdir /expr/data; fi"

# 检查并在glee3中创建/expr目录和/expr/data目录
docker exec -it glee3 bash -c "if [ ! -d '/expr' ]; then mkdir /expr; fi"
docker exec -it glee3 bash -c "if [ ! -d '/expr/data' ]; then mkdir /expr/data; fi"


# 定义要复制的文件路径
SOURCE_FILE="../bin/glee-etcd"

# 定义目标容器
CONTAINERS=("glee1" "glee2" "glee3")

# 遍历每个容器,执行复制
for container in ${CONTAINERS[@]}  
do
  docker cp $SOURCE_FILE $container:/expr/
done


docker exec -it glee1 chmod +x /expr/glee-etcd
docker exec -it glee2 chmod +x /expr/glee-etcd
docker exec -it glee3 chmod +x /expr/glee-etcd