#!/bin/bash

echo "Building application"
go build
kill -15 $(pidof bsm-be-gin)
#nohup ./bsm-be-gin > /dev/null 2>&1&
pm2 stop 0
pm2 start bsm-be-gin
echo "Deploy completed"