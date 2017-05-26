#!/bin/bash

export IMPORT_PATH=$GOPATH/src:.
export GENERATOR="gofast_out"
export OUTPUT_DIR="."
export PROTO_FILES="./*.proto"

export OPTIONS_API="Mmodels.proto=github.com/otsimo/otsimopb,Mmessages.proto=github.com/otsimo/otsimopb,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types"

protoc --proto_path=$IMPORT_PATH --${GENERATOR}=plugins=grpc:${OUTPUT_DIR} $PROTO_FILES
protoc --proto_path=$IMPORT_PATH --${GENERATOR}=${OPTIONS_API},plugins=grpc:.. ./v2/*.proto

