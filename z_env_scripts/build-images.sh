#!/bin/bash

DOCKER_BUILDKIT=1

TOP=$(git rev-parse --show-toplevel)
cd $TOP

#api
# cd api && docker build -t grocery-api:test . && cd ../
cd api && docker build -f Dockerfile.prod -t grocery-api:prod . && cd ../

#Frontend
cd frontend && docker build -t grocery-frontend:live . && cd ../
