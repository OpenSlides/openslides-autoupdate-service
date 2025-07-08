#!/bin/bash

# Executes all tests. Should errors occur, CATCH will be set to 1, causing an erroneous exit code.

echo "########################################################################"
echo "###################### Run Tests and Linters ###########################"
echo "########################################################################"

# Setup
IMAGE_TAG=openslides-autoupdate-tests

# Safe Exit
trap 'docker stop $(docker ps -a -q --filter ancestor=${IMAGE_TAG})' EXIT

# Execution
if [ "$(docker images -q $IMAGE_TAG)" = "" ]; then make build-test; fi
docker run --privileged -t ${IMAGE_TAG} ./dev/container-tests.sh