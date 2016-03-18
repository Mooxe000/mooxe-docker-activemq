#!/usr/bin/env bash

cwd="$(cd -P -- "$(dirname -- "$0")" && pwd -P)"
ID="2"
ImgName="activemq"

set -e

docker run \
  --name activemq_${ID} \
  --rm \
  -ti \
  -e 'ACTIVEMQ_MIN_MEMORY=512' \
  -e 'ACTIVEMQ_MAX_MEMORY=2048' \
  -v ${cwd}/amqconf/GoSource/amqconf/dist/${ImgName}.xml:/opt/${ImgName}/conf/${ImgName}.xml \
  -p 8162:8161 \
  -p 61617:61616 \
  -p 61614:61613 \
  mooxe/${ImgName} \
  /bin/bash
