#!/bin/bash
set -ex
#SCRIPT_DIR=$(dirname $(readlink -e $0))
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DOCKER_HOME=${SCRIPT_DIR}/../
PROJECT_HOME=${SCRIPT_DIR}/../../
. ${SCRIPT_DIR}/image

pushd ${PROJECT_HOME}/web
rm -rf web/build
npm run build
cp -r build ${DOCKER_HOME}/web/build
popd

pushd ${PROJECT_HOME}
export GOOS=linux
export GOARCH=amd64
rm -rf jarvis_master_${GOOS}_${GOARCH}.bin
go clean && go build -o jarvis_master_${GOOS}_${GOARCH}.bin
cp jarvis_master_linux_amd64.bin ${DOCKER_HOME}/
popd

echo "jarvis master build done"
