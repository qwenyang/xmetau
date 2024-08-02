#!/usr/bin/env bash

protoDir="./"
outDir="./cgi"

# 编译自定义的proto   
protoc -I ${protoDir}/ ${protoDir}/*.proto \
	--go_out ${outDir} \
	--go_opt paths=source_relative \
	--go-grpc_out ${outDir} \
	--go-grpc_opt paths=source_relative \
	--go-grpc_opt require_unimplemented_servers=false \
	--grpc-gateway_out ${outDir} \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
