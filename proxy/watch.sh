#!/bin/bash

go build .
./proxy &
PID=$!
while inotifywait -e create -e moved_to -r -q -q .; do
  echo "reload"
  kill $PID
  sleep 1
  go build .
  ./proxy &
  PID=$!
done

