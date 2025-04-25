#!/bin/bash
# This script compiles the client and server for Linux with CGO disabled and static linking for the client.
gcc -o client client.c -static
gcc -o server server.c 
