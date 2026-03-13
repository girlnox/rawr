#!/bin/bash

mkdir -p exports/

echo "Building rawr-amd64..."

export GOARCH=amd64
export GOOS=linux

go build main.go
mv main exports/rawr-amd64 

echo "Building rawr-aarch64..."

export GOARCH=arm64
export GOOS=linux

go build main.go
mv main exports/rawr-aarch64

echo "Adding export time info"

date > exports/date.txt
cat exports/date.txt

echo "All done!!!"
