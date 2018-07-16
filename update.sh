#!/bin/bash

docker pull $1/market-prod:latest
if docker stop market-app; then docker rm market-app; fi
docker run -d -p 8077:8077 --name market-app $1/market-prod
if docker rmi $(docker images --filter "dangling=true" -q --no-trunc); then :; fi