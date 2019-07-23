FROM golang:1.11.9-alpine3.9
WORKDIR /go/src/app
RUN apk add --no-cache git gcc musl-dev
ENV GO111MODULE=on
EXPOSE 8090
COPY main.go home.html go.mod ./ 
RUN go build main.go
CMD ["./main"]