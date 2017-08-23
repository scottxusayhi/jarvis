#!/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
. ${SCRIPT_DIR}/image

if [ $# -lt 1 ]; then
    echo "Usage: $0 <env_file>"
    exit -1
fi

set -e
ENV_FILE=$1
python ${SCRIPT_DIR}/_check_env.py $ENV_FILE
shift
IMAGE=${IMAGE} ENV_FILE=${ENV_FILE} docker-compose -f ${SCRIPT_DIR}/docker-compose.yml up $*
