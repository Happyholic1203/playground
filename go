#!/bin/bash

if [ -z $1 ]; then
    echo "Usage: $0 <mysql|sa|sqlalchemy|tmux|bash>" >&2
    exit 0
fi

CONTAINER_NAME=playground
IMAGE=happyholic1203/playground:latest
SCRIPTS_DIR=`pwd`/`dirname $0`/scripts

docker run -it --rm \
-v $SCRIPTS_DIR:/scripts \
--name=$CONTAINER_NAME \
$IMAGE /scripts/init $1
