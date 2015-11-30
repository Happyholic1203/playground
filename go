#!/bin/bash

if [ -z $1 ]; then
    echo "Usage: $0 <mysql|sa|sqlalchemy|git|tmux|bash|mq>" >&2
    exit 0
fi

CONTAINER_NAME=playground
IMAGE=happyholic1203/playground:latest
SCRIPTS_DIR=`pwd`/`dirname $0`/scripts
README=`pwd`/`dirname $0`/README.md

docker run -it --rm \
-v $SCRIPTS_DIR:/scripts \
-v $README:/root/README.md \
--name=$CONTAINER_NAME \
$IMAGE /scripts/init $1
