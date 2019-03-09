#!/usr/bin/env bash
SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"

echo "Getting service ip..."
cd $SCRIPTPATH
cd ../terraform
serviceIp=`terraform output service-ip`

echo "Building service..."
cd $SCRIPTPATH
cd ../chat-service/service
go build -o main .

echo "Uploading files..."
cd $SCRIPTPATH
scp -r ../chat-service/service/main service@$serviceIp:/home/service
scp -r ../chat-service/Dockerfile service@$serviceIp:/home/service

echo "Spinning up docker..."
ssh service@$serviceIp 'bash -s' < _update_service.sh
