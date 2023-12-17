#!/usr/bin/env bash

docker rm edu-mysql

UID=$(id -u)
GID=$(id -g)

docker run --platform linux/amd64 -d \
  --user $UID:$GID \
  --name edu-mysql \
  -e MYSQL_ROOT_PASSWORD=test \
  -e MYSQL_DATABASE=test  \
  -p13306:3306  \
  mysql:5.7.35 --character-set-server=utf8 --collation-server=utf8_general_ci
