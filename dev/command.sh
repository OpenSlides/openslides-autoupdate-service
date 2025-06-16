#!/bin/sh

if [ ! -z $dev   ]; then CompileDaemon -log-prefix=false -build="go build" -command="./openslides-autoupdate-service"; fi
if [ ! -z $tests ]; then go vet ./... && go test -test.short ./...; fi