FROM golang:1.11.9-alpine3.9 as BUILD

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
COPY ./ ./

WORKDIR /app/web
RUN go build .

FROM alpine

COPY --from=BUILD /app/web/web /app/web
WORKDIR /app
COPY ./web/home.html /app/home.html
CMD ["./web"]
