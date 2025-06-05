#!/bin/sh

if [ ! -z $dev   ]; then CompileDaemon -log-prefix=false -build="go build -o autoupdate-service ./openslides-autoupdate-service" -command="./autoupdate-service"; fi
if [ ! -z $tests ]; then go vet ./... && go test -test.short ./...; fi