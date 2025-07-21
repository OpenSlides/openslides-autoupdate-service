#!/bin/sh
# Timeout argument to be used in GitHub Workflows

if [ "$APP_CONTEXT" = "dev" ]
then
    if [ -n "$DEV_TIMEOUT" ]
    then
        echo "Go build and wait for $DEV_TIMEOUT seconds"
        CompileDaemon -log-prefix=false -build="go build" -command="./openslides-autoupdate-service" &

        # Sleep for 15 seconds, then check if build executable exists
        sleep "$DEV_TIMEOUT"

        echo "Done sleeping"
        ls -a
        if [ -f "./openslides-autoupdate-service" ]; then echo "Found go executable"; fi
        echo "Couldn't find go executable"
        exit 1
    else
        CompileDaemon -log-prefix=false -build="go build" -command="./openslides-autoupdate-service"
    fi
fi
if [ "$APP_CONTEXT" = "tests" ]; then sleep inf; fi