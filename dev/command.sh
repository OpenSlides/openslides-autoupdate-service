#!/bin/sh
# Timeout argument to be used in GitHub Workflows

if [ "$APP_CONTEXT" = "dev"   ]
then
    if [ -n "$TIMEOUT_DEV" ]
    then
        timeout --preserve-status --signal=SIGINT 5s CompileDaemon -log-prefix=false -build="go build" -command="./openslides-autoupdate-service"
    else
        CompileDaemon -log-prefix=false -build="go build" -command="./openslides-autoupdate-service"
    fi
fi
if [ "$APP_CONTEXT" = "tests" ]; then sleep inf; fi