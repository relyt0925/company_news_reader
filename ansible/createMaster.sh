#!/usr/bin/env bash

#./createMasters.sh 7001 5 redis-cluster-m- jamespedwards42/alpine-redis:3.2
readonly STARTING_PORT=${1:-0}
readonly NUM_MASTERS=${2:-0}
readonly NAME_PREFIX=${3:-"redis-cluster-m-"}
readonly IMAGE=${4:-"jamespedwards42/alpine-redis:3.2"}
for ((port = STARTING_PORT, endPort = port + NUM_MASTERS; port < endPort; port++)) do
  name="$NAME_PREFIX$port"
  docker -H tcp://192.168.70.103:2400 run \
    -d \
    --name "$name" \
    -p "$port:$port"/tcp \
    --net=host \
    --restart "always" \
      "$IMAGE" \
        --appendfsync everysec \
        --appendonly yes \
        --auto-aof-rewrite-percentage 20 \
        --cluster-enabled yes \
        --cluster-node-timeout 60000 \
        --cluster-require-full-coverage no \
        --logfile "$name".log \
        --port "$port" \
        --protected-mode no \
        --repl-diskless-sync yes \
        --save ''''
done