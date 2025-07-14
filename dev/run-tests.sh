#!/bin/bash

# Executes all tests. Should errors occur, CATCH will be set to 1, causing an erroneous exit code.

echo "########################################################################"
echo "###################### Run Tests and Linters ###########################"
echo "########################################################################"

# Parameters
while getopts "s" FLAG; do
    case "${FLAG}" in
    s) SKIP_BUILD=true ;;
    *) echo "Can't parse flag ${FLAG}" && break ;;
    esac
done

# Setup
IMAGE_TAG=openslides-autoupdate-tests
LOCAL_PWD=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Safe Exit
trap 'docker stop autoupdate-test && docker rm autoupdate-test' EXIT

# Execution
if [ -z "$SKIP_BUILD" ]; then make build-tests; fi
docker run --privileged -t ${IMAGE_TAG} --name autoupdate-test ./dev/container-tests.sh

# Linters
bash "$LOCAL_PWD"/run-lint.sh -s -c