#!/usr/bin/env bash

cwd="$(cd -P -- "$(dirname -- "$0")" && pwd -P)"

imgName="amqconf"
homeDir="/root/${imgName}"

set -e

docker run \
  --name ${imgName} \
  --rm \
  -ti \
  -v ${cwd}/${imgName}/bin:${homeDir}/bin \
  -v ${cwd}/${imgName}/res:${homeDir}/res \
  -v ${cwd}/${imgName}/dist:${homeDir}/dist \
  -v ${cwd}/${imgName}/conf:${homeDir}/conf \
  -e GOPATH=/go:${homeDir} \
  mooxe/base \
  /usr/bin/env bash -lc " \
    /root/${imgName}/bin/${imgName} --conf /root/${imgName}/conf/amqconf.json Dev
  "
