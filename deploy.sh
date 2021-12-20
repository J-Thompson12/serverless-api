#!/bin/bash

echo -e "\n+++++ Starting deployment +++++\n"

tfswitch 1.0.0

rm -rf ./bin

echo "+++++ build go packages +++++"

go test ./...
env GOOS=linux GOARCH=amd64 go build -o ./bin/api
cd ../..

echo "+++++ hello module +++++"
cd infrastructure
terraform init -input=false
terraform apply -input=false -auto-approve
cd ../

echo -e "\n+++++ Deployment done +++++\n"