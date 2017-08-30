#!/usr/bin/env bash

set -e

THRIFT_ARGS="-out . --gen go"
for F in $PATH_TO_THRIFT_FILES/*.thrift; do
    thrift ${THRIFT_ARGS} "$F" || $(echo "Failed to generate golang based on thrift file $F"; exit 1)
done
