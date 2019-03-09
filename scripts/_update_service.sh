#!/usr/bin/env bash

existingContainerId=`docker ps -qf "name=chat-service"`

if [ -z "$existingContainerId" ]
then
    echo "No previous container..."
else
    echo "Copying certificates..."
    docker cp $existingContainerId:/root/.local/share/certmagic /home/service/certmagic

    echo "Killing previous container..."
    docker kill chat-service
    docker rm chat-service
fi

mkdir -p certmagic
docker build -t chat-service /home/service/
docker run -d --name chat-service -p 443:443 -p 80:80 -it chat-service