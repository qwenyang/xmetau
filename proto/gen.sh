#!/usr/bin/env bash

outDir="../../../../"

protoc -I ./ --go_out ${outDir} --go-grpc_out ${outDir} common.proto 
protoc -I ./ --go_out ${outDir} --go-grpc_out ${outDir} unidao.proto 
