#!/bin/bash
DOCKER_BUILDKIT=1
#Frontend
# cd frontend && docker build -t grocery_frontend:test . && cd ../

#backend
cd backend && docker build -t grocery_backend:test . && cd ../
docker compose restart backend
