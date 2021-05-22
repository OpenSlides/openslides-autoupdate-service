#!/bin/sh
if [ ! -z $1 ]
then
    export TEST_CASE=$1
fi

go test -race ./tests
