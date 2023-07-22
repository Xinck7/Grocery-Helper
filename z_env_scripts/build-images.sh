#!/bin/bash

DOCKER_BUILDKIT=1

TOP=$(git rev-parse --show-toplevel)
cd $TOP

#api
cd api && docker build -t grocery-api:test . && cd ../

#Frontend
# cd frontend && docker build -t grocery_frontend:test . && cd ../
