#! /bin/bash

# 文件类型 => conf 文件  ; api 文件
if [ ! "$1" ]; then
  printf "缺少proto文件类型参数, 例如 ./proto.sh conf \n"
  exit 0
fi

if [ "$1" == "conf" ]; then
  # 编译配置protobuf文件
  INTERNAL_PROTO_FILES=$(find ./pkg/conf -name *.proto)
  protoc --proto_path=./pkg/conf \
  	       --proto_path=./third_party \
   	       --go_out=paths=source_relative:./pkg/conf \
  	       $INTERNAL_PROTO_FILES
elif [ "$1" == "api" ]; then
  # 编译api
  API_PROTO_FILES=$(find api -name *.proto)
  protoc --proto_path=. \
  	       --proto_path=./third_party \
   	       --go_out=paths=source_relative:. \
   	       --go-http_out=paths=source_relative:. \
   	       --go-grpc_out=paths=source_relative:. \
   	       --openapi_out==paths=source_relative:. \
   	       --go-errors_out=paths=source_relative:. \
  	       $API_PROTO_FILES
fi
