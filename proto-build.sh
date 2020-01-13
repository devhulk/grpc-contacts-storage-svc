#!/bin/bash
protoc -I $(pwd)/internal/storage storage.proto --go_out=plugins=grpc:storage
