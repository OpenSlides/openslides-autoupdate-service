#!/bin/bash

# Executes all tests. Should errors occur, CATCH will be set to 1, causing an erroneous exit code.

echo "########################################################################"
echo "###################### Run Tests and Linters ###########################"
echo "########################################################################"

# Parameters
while getopts "p" FLAG; do
    case "${FLAG}" in
    p) PERSIST_CONTAINERS=true ;;
    *) echo "Can't parse flag ${FLAG}" && break ;;
    esac
done

# Setup
IMAGE_TAG=openslides-autoupdate-tests
LOCAL_PWD=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
CATCH=0

# Execution
if [ "$(docker images -q $IMAGE_TAG)" = "" ]; then make build-tests || CATCH=1; fi
docker run -d --name autoupdate-test ${IMAGE_TAG} || CATCH=1
docker exec autoupdate-test go vet ./... || CATCH=1
docker exec autoupdate-test go test -test.short ./... || CATCH=1

# Linters
bash "$LOCAL_PWD"/run-lint.sh -b || CATCH=1

if [ -z "$PERSIST_CONTAINERS" ]; then docker stop autoupdate-test && docker rm autoupdate-test || CATCH=1; fi

exit $CATCH