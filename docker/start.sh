#!/bin/bash

case $1 in
  "run")
    shift
    ./avalon "$@"
    ;;
  "build")
    cp /go/src/avalon/avalon /binary
    cp /go/src/avalon/start.sh /binary
    chown -R $(id -u):$(id -u) /binary
    ;;
  *)
    echo "usage: $0 [run|build]"
    exit 1
    ;;
esac