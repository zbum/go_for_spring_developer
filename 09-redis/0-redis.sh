#!/usr/bin/env bash

docker stop edu-redis
docker rm edu-redis

docker run -d --name edu-redis \
  -p6379:6379  \
 redis