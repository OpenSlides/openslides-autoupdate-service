#!/bin/bash

# Executes all linters. Should errors occur, CATCH will be set to 1, causing an erroneous exit code.

echo "########################################################################"
echo "###################### Run Linters #####################################"
echo "########################################################################"

# Parameters
while getopts "lscp" FLAG; do
    case "${FLAG}" in
    l) LOCAL=true ;;
    s) SKIP_BUILD=true ;;
    c) SKIP_CONTAINER_UP=true ;;
    p) PERSIST_CONTAINERS=true ;;
    *) echo "Can't parse flag ${FLAG}" && break ;;
    esac
done

# Setup
IMAGE_TAG=openslides-autoupdate-tests
CATCH=0
DOCKER_EXEC="docker exec autoupdate-test"

# Safe Exit
trap 'if [ -z "$PERSIST_CONTAINERS" ] && [ -z "$SKIP_CONTAINER_UP" ]; then docker stop autoupdate-test && docker rm autoupdate-test' EXIT

# Optionally build image
if [ -z "$SKIP_BUILD" ]; then make build-tests; fi

# Execution
if [ -z "$LOCAL" ]
then
    # Container Mode
    if [ -z "$SKIP_CONTAINER_UP" ]; then docker run -d -t --name autoupdate-test ${IMAGE_TAG}; fi
    eval "$DOCKER_EXEC go vet ./..."
    eval "$DOCKER_EXEC golint -set_exit_status ./..."
    eval "$DOCKER_EXEC gofmt -l ."
else
    # Local Mode
    go vet ./...
    golint -set_exit_status ./...
    gofmt -l -s -w .
fi