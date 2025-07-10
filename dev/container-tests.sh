#!/bin/sh

dockerd --storage-driver=vfs --log-level=error &

# Close Dockerd savely on exit
DOCKERD_PID=$!
trap 'kill $DOCKERD_PID' EXIT INT TERM ERR

RETRY=0
until docker info >/dev/null 2>&1; do
  if [ "$RETRY" -ge 10 ]; then
    echo "Dockerd setup error"
    exit 1
  fi
  echo "Waiting for dockerd"
  sleep 1
  RETRY=$(tries + 1)
done

echo "Started dockerd"

CATCH=0

# Run Linters & Tests
go vet ./... || CATCH=1
go test -timeout 60s -race ./... || CATCH=1
gofmt -l . || CATCH=1
golint -set_exit_status ./... || CATCH=1

exit $CATCH