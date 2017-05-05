#!/bin/bash
options=(app server)
if [ $# -lt 1 ];then
    echo -e "Error: need an argument, should be one of : ${options[@]}"
    echo -e "\"$0 app\" to build app.go\n\"$0 server\" to build a static file server."
    exit 1
fi
if [ $1 == 'app' ]; then
    go build -ldflags "-X main.version=`date +%Y_%m_%d_%H:%M:%S`" app.go
elif [ $1 == 'server' ]; then
    go build archive/file_server/file_server.go
else
    echo "invalid argument"
fi