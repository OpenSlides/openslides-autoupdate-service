#!/bin/bash

# Executes all tests. Should errors occur, CATCH will be set to 1, causing an erroneous exit code.

echo "########################################################################"
echo "###################### Run Tests and Linters ###########################"
echo "########################################################################"

# Parameters
PERSIST_CONTAINERS=$1

# Setup
IMAGE_TAG=openslides-autoupdate-tests
CATCH=0

# Execution
if [ "$(docker images -q $IMAGE_TAG)" = "" ]; then make build-test || CATCH=1; fi
docker run ${IMAGE_TAG} || CATCH=1

if [ -z "$PERSIST_CONTAINERS" ]; then docker stop "$(docker ps -a -q --filter ancestor=${IMAGE_TAG} --format="{{.ID}}")" || CATCH=1; fi

exit $CATCH