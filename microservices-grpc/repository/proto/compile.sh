protoc repository.proto \
    --go_out=plugins=grpc:. \
    --grpc-gateway_out=logtostderr=true:. \
    --swagger_out=logtostderr=true:. \
    -I. -Igoogleapis