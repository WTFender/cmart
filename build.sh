#!/bin/bash
GOOS=linux GOARCH=ppc64 go build -o ./bin/cmart.elf ./cmd/cmart/
GOOS=windows GOARCH=amd64 go build -o ./bin/cmart.exe ./cmd/cmart/
GOOS=darwin GOARCH=amd64 go build -o ./bin/cmart.app ./cmd/cmart/