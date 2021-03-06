#!/bin/bash
set -e
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
. ${SCRIPT_DIR}/image

docker build -t ${IMAGE} \
  --build-arg branch=${GIT_BRANCH} \
  --build-arg commit=$(git rev-parse HEAD) \
  --build-arg buildtime="$(date +"%Y-%m-%d %T")"  \
  --build-arg owner="${IMAGE_OWNER}" ${SCRIPT_DIR}/../
#  --build-arg env_para="$(cat $SCRIPT_DIR/env-list.yaml)" \


echo -e "\n$IMAGE"
docker inspect --format "branch={{.Config.Labels.branch}} commit={{.Config.Labels.commit}} buildtime={{.Config.Labels.buildtime}} owner={{.Config.Labels.owner}}" $IMAGE
echo "env_para="
docker inspect --format "{{.Config.Labels.env_para}}" $IMAGE

