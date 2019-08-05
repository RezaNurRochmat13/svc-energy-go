#!/usr/bin/env bash

#Compile golang files to binary executable
env GOOS=linux CGO_ENABLED=1 go build -a -ldflags="-s -w" -installsuffix cgo -o main .
