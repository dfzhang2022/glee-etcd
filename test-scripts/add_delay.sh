#!/bin/bash

# 添加时延的容器
CONTAINERS=(glee1 glee2 glee3)

# 各容器对应的时延
DELAYS=(50ms 30ms 20ms)

for i in "${!CONTAINERS[@]}"
do
    # 获取容器的eth0接口
     eth0=$(docker inspect ${CONTAINERS[$i]} | jq -r '.[0].NetworkSettings.Networks.brige.EndpointID')
  #  eth0=$(docker inspect ${CONTAINERS[$i]} | jq -r '.[0].NetworkSettings.Networks.glee-net.EndpointID')
    eth0=$(docker inspect ${CONTAINERS[$i]} |grep -o '"EndpointID": "[^"]*"' | cut -d'"' -f4)
 #   echo ${eth0}
    # 在该接口上添加时延
    tc qdisc add dev ${eth0} root netem delay ${DELAYS[$i]}
done