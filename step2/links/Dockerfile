FROM golang:1.11.9-alpine3.9 as BUILD

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
COPY ./ ./

RUN go build .

FROM alpine

COPY --from=BUILD /app/monolith /app/monolith
WORKDIR /app
COPY home.html /app/home.html
CMD ["./monolith"]
