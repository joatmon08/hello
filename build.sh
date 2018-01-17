#!/bin/bash

GOOS=linux GOARCH=amd64 go build main.go
docker build -t joatmon08/hello:latest .
