#!/bin/bash

#@ToDo: use APP_ROOT env variable, after app dockerize.

protoc -I/usr/local/include \
		-I../ \
		-I../api/protobuf/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--grpc-gateway_out=logtostderr=true,paths=source_relative:../pkg/pb \
		--swagger_out=allow_merge=true,merge_file_name=api:../api/swagger \
		--go_out=plugins=grpc,paths=source_relative:../pkg/pb calendar.proto