#!/usr/bin/env bash

cwd="$(cd -P -- "$(dirname -- "$0")" && pwd -P)"

set -e

rm -rf ${cwd}/amqconf/res/activemq.xml

docker run \
  --name amqconf \
  --rm \
  -ti \
  -v ${cwd}/amqconf/res:/root/amqconf \
  mooxe/activemq \
  /usr/bin/env bash -lc ' \
    cp /opt/activemq/conf/activemq.xml \
      /root/amqconf/activemq.xml \
  '

  # -v $(pwd)/activemq.xml:/opt/activemq/conf/activemq.xml \
  # -v $(pwd)/jetty.xml:/opt/activemq/conf/jetty.xml \
  # /app/init.py start
