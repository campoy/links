protoc repository.proto \
    --go_out=plugins=grpc:. \
    -I. -Igoogleapis