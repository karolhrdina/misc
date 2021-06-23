#!/bin/bash

# NOTE:
# This is poor man's protobuf generation.
# 
# Although good enough for a PoC like this, a better way would be to use 
# something like `https://github.com/bufbuild/buf` where using config yaml's
# you can descriptively control what and how is generated without the need
# for the nitty-gritty details `protoc` and `protoc-gen-go-grpc` :)
  
# check that protoc-gen-go is installed
GO_PROTOC_GEN=$(go env GOPATH)/bin/protoc-gen-go
if [[ ! -x ${GO_PROTOC_GEN} ]]; then
    echo -e "ERROR: Can not find ${GO_PROTOC_GEN}"
    exit 1
fi

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." || exit 64 ; pwd -P)"
export BASE_DIR
cd "$BASE_DIR"

export PBGO_DIR="${PBGO_DIR:-"${BASE_DIR}/pb.go"}"

rm -rf "${PBGO_DIR}"
mkdir "${PBGO_DIR}"

protoc -I proto/ --go_out=pb.go --go-grpc_out=pb.go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  proto/port-domain.proto
