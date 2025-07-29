#!/bin/bash

# Executes all linters. Should errors occur, CATCH will be set to 1, causing an erroneous exit code.

echo "########################################################################"
echo "###################### Run Linters #####################################"
echo "########################################################################"

# Parameters
while getopts "ls" FLAG; do
    case "${FLAG}" in
    l) LOCAL=true ;;
    s) SKIP_SETUP=true ;;
    *) echo "Can't parse flag ${FLAG}" && break ;;
    esac
done

# Setup
CONTAINER_NAME="autoupdate-tests"
IMAGE_TAG=openslides-autoupdate-tests
DOCKER_EXEC="docker exec ${CONTAINER_NAME}"

# Safe Exit
trap 'if [ -z "$LOCAL" ] && [ -z "$SKIP_SETUP" ]; then docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME; fi' EXIT

# Execution
if [ -z "$LOCAL" ]
then
    # Setup
    if [ -z "$SKIP_SETUP" ]
    then
        make build-tests >/dev/null 2>&1
        docker run -d --name "${CONTAINER_NAME}" "${IMAGE_TAG}"
    fi

    # Container Mode
    eval "$DOCKER_EXEC go vet ./..."
    eval "$DOCKER_EXEC golint -set_exit_status ./..."
    eval "$DOCKER_EXEC gofmt -l ."
else
    # Local Mode
    go vet ./...
    golint -set_exit_status ./...
    gofmt -l -s -w .
fi