#!/usr/bin/env bash

docker rm edu-mysql

USER_ID=$(id -u)
GROUP_ID=$(id -g)

docker run -d \
  --platform=linux/amd64 \
  --user "$USER_ID":"$GROUP_ID" \
  --name edu-mysql \
  -e MYSQL_ROOT_PASSWORD=test \
  -e MYSQL_DATABASE=test  \
  -p3306:3306  \
  mysql:5.7.35 --character-set-server=utf8 --collation-server=utf8_general_ci


## connect
## docker exec -it edu-mysql mysql -u root -p -P3306
##  delete from Scores;
##  delete from Students;