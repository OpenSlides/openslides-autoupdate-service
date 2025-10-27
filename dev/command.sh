#!/bin/sh
# Timeout argument to be used in GitHub Workflows

if [ "$APP_CONTEXT" = "dev" ]; then exec CompileDaemon -log-prefix=false -build="go build" -directory="../" -build-dir="./" -command="./openslides-autoupdate-service"; fi
if [ "$APP_CONTEXT" = "tests" ]; then sleep inf; fi
