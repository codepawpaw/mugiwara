#! /bin/sh
set -e

export GOARCH="amd64"
export GOOS="linux"
export CGO_ENABLED=0

go build -o dist/go-mysql-crud .

docker build -t go-mysql-crud .
