#!/bin/bash

CONTAINERS=("glee1" "glee2" "glee3")

for container in ${CONTAINERS[@]} 
do
    if docker inspect -f '{{.State.Running}}' "$container" &>/dev/null; then
        echo "$container is running, restarting it..."
        docker restart "$container"
    else
        echo "$container is not running, starting it..."
        docker start "$container"
    fi
done