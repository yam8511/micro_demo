#! /bin/bash
go build .
./pusher --registry_address=$REGISTRY_ADDRESS
