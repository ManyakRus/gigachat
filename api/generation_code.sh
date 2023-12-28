#!/bin/bash 

ADAPT_PATH=./

clear
protoc --go_out=${ADAPT_PATH}/grpc_proto --go_opt=paths=import --go-grpc_out=${ADAPT_PATH}/grpc_proto --go-grpc_opt=paths=import ${ADAPT_PATH}gigachat.proto
