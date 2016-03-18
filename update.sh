#!/usr/bin/env bash

cwd="$(cd -P -- "$(dirname -- "$0")" && pwd -P)"

set -e

npm run pullSub

sed -ie 's/webcenter\/openjdk-jre:8/mooxe\/java8/g' \
  ./activemq/Dockerfile

rm -rf ./activemq/Dockerfilee
