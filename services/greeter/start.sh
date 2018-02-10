#! /bin/bash
go build .
./greeter --registry_address=$REGISTRY_ADDRESS
