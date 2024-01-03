#!/bin/bash
mkdir -p ./storage
chmod 777 -R ./storage
go install github.com/codegangsta/gin@latest
gin --appPort 8080 --immediate
