#!/bin/sh

dockerd --storage-driver=vfs --log-level=error &

# Close Dockerd savely on exit
DOCKERD_PID=$!
trap 'kill $DOCKERD_PID' EXIT INT TERM ERR

RETRY=0
MAX=10
until docker info >/dev/null 2>&1
do
  if [ "$RETRY" -ge "$MAX" ]
  then
    echo "Dockerd setup error"
    exit 1
  fi
  sleep 1
  RETRY=$((RETRY + 1))
  echo "Waiting for dockerd $RETRY/$MAX"
done

echo "Started dockerd"

CATCH=0

# Run Linters & Tests
go vet ./... || CATCH=1
go test -timeout 60s -race ./... || CATCH=1
gofmt -l . || CATCH=1
golint -set_exit_status ./... || CATCH=1

exit $CATCH