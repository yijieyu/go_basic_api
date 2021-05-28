#!/bin/bash
set -ex
app=app_api
task=$1
go build -mod vendor -tags "jsoniter" -o $app

if [ -n "$task" ]; then
  ./$app
fi
