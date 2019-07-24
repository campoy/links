FROM golang:1.11.9-alpine3.9 as BUILD

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
COPY ./ ./

WORKDIR /app/repository
RUN go build ./cmd/server

FROM alpine

COPY --from=BUILD /app/repository/server /app/server
CMD ["/app/server"]
