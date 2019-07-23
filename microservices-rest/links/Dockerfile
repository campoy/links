FROM golang:1.11.9-alpine3.9
WORKDIR /go/src/app
RUN apk add --no-cache git gcc musl-dev
ENV GO111MODULE=on
EXPOSE 8080
COPY main.go go.mod ./ 
# RUN go mod vendor
RUN go build main.go
CMD ["./main"]