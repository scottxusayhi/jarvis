#!/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
. ${SCRIPT_DIR}/image

docker run -it -p 2999:2999 $IMAGE --mysql-addr=10.1.10.99:3306
