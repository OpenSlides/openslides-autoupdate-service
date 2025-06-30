#!/bin/bash

# Executes all linters. Should errors occur, CATCH will be set to 1, causing an erroneous exit code.

echo "########################################################################"
echo "###################### Run Linters #####################################"
echo "########################################################################"

# Parameters
while getopts "lbp" FLAG; do
    case "${FLAG}" in
    l) LOCAL=true ;;
    b) BUILD=true ;;
    p) PERSIST_CONTAINERS=true ;;
    *) echo "Can't parse flag ${FLAG}" && break ;;
    esac
done

# Setup
IMAGE_TAG=openslides-autoupdate-tests
CATCH=0
DOCKER_EXEC="docker exec autoupdate-test"

# Optionally build image
if [ -n "$BUILD" ]
then
    if [ "$(docker images -q $IMAGE_TAG)" = "" ]; then make build-tests || CATCH=1; fi
fi

# Execution
if [ -z "$LOCAL" ]
then
    # Container Mode
    docker run -d -t --name autoupdate-test ${IMAGE_TAG} || CATCH=1
    eval "$DOCKER_EXEC go vet ./..." || CATCH=1
    eval "$DOCKER_EXEC golint -set_exit_status ./..." || CATCH=1
    eval "$DOCKER_EXEC gofmt -l ." || CATCH=1
else
    # Local Mode
    go vet ./... || CATCH=1
    golint -set_exit_status ./... || CATCH=1
    gofmt -l -s -w . || CATCH=1
fi

if [ -z "$PERSIST_CONTAINERS" ] && [ -n "$BUILD" ]; then docker stop autoupdate-test && docker rm autoupdate-test || CATCH=1; fi

exit $CATCH