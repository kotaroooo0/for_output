#!/bin/bash

set -e

export DOCKER_BUILDKIT=0

docker pull test:stage1 || true &
docker pull test:stage2 || true &
wait

docker build -t test:latest --cache-from=test:stage1,test:stage2 .

docker tag $(docker images --filter 'label=stage=1' -q | head -n 1) test:stage1
docker tag $(docker images --filter 'label=stage=2' -q | head -n 1) test:stage2

docker push test:latest &
docker push test:stage1 &
docker push test:stage2 &
wait

# ダメな例
docker build -t test:latest .
docker build -t test:stage1 --target stage1 --cache-from test:stage1 .
docker build -t test:stage2 --target stage2 --cache-from test:stage2 .


export DOCKER_BUILDKIT=1

docker build -t test:latest --cache-from=test:stage1,test:stage2 --build-arg BUILDKIT_INLINE_CACHE=1  .

docker build -t test:stage1 --target stage1 --cache-from=test:stage1 --build-arg BUILDKIT_INLINE_CACHE=1  . &
docker build -t test:stage2 --target stage2 --cache-from=test:stage2 --build-arg BUILDKIT_INLINE_CACHE=1  . &
wait

docker push test:latest &
docker push test:stage1 &
docker push test:stage2 &
wait


# ダメな例
