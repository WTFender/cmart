#!/bin/bash
go test -v || { echo -e 'STOP\tGo tests failed, stopped build' ; exit 1; }
echo '=== BUILD'
GOOS=linux   GOARCH=amd64 go build -o ./bin/cmart.elf ./cmd/cmart/ && echo -e 'ok\t./bin/cmart.elf'
GOOS=windows GOARCH=amd64 go build -o ./bin/cmart.exe ./cmd/cmart/ && echo -e 'ok\t./bin/cmart.exe'
GOOS=darwin  GOARCH=amd64 go build -o ./bin/cmart.app ./cmd/cmart/ && echo -e 'ok\t./bin/cmart.app'
echo '=== DONE'