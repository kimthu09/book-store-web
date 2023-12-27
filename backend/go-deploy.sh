#!/bin/bash

echo "Building application"
go build
kill -15 $(pidof propq-search-service-v3-gin)
#nohup ./propq-search-service-v3-gin > /dev/null 2>&1&
pm2 stop 0
pm2 start propq-search-service-v3-gin
echo "Deploy completed"