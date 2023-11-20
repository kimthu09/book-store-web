#!/bin/bash
mkdir -p ./storage
chmod 777 -R  ./storage
gin --appPort 8080 --immediate