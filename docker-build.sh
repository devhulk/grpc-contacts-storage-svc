#!/bin/bash

# rm all <none> images
# docker rmi -f $(docker images -f "dangling=true" -q)

# Official Golang Image Build
docker rm -f storage-service && docker rmi -f storage-service
docker build -t storage-service .; docker run -d -p 5540:5540 --name=storage-service storage-service

# Minimal build using Docker Scratch Image
# docker rm -f storage-service-scratch && docker rmi -f storage-service-scratch
# docker build -t storage-service-scratch -f Dockerfile.scratch .; docker run -d -p 5540:5540 --name=storage-service-scratch storage-service-scratch