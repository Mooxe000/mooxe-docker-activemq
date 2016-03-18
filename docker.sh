#!/usr/bin/env bash

cwd="$(cd -P -- "$(dirname -- "$0")" && pwd -P)"

set -e

docker run \
  --name activemq \
  --rm \
  -ti \
  -e 'ACTIVEMQ_MIN_MEMORY=512' \
  -e 'ACTIVEMQ_MAX_MEMORY=2048' \
  -v ${cwd}:/root/activemq \
  -p 8160:8161 \
  mooxe/activemq \
  /bin/bash

  # -v $(pwd)/activemq.xml:/opt/activemq/conf/activemq.xml \
  # -v $(pwd)/jetty.xml:/opt/activemq/conf/jetty.xml \
  # /app/init.py start
