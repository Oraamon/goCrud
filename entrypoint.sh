#!/bin/bash

echo "Building the Go application..."
go build -o bin ./cmd/api/main.go

echo "Running the Go application..."
./bin
