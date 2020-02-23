#!/bin/bash

go build .
./core_api &
PID=$!
while inotifywait -e create -e moved_to -r -q -q .; do
  echo "reload"
  kill $PID
  sleep 1
  go build .
  ./core_api &
  PID=$!
done

