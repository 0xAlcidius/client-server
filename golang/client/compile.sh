#!/bin/bash
# This script compiles the Go client for Linux with CGO disabled and static linking.
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o client client.go