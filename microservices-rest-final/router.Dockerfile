FROM golang:1.11.9-alpine3.9 as BUILD

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app
ENV GO111MODULE=on
COPY ./ ./

WORKDIR /app/router
RUN go build .

FROM alpine

COPY --from=BUILD /app/router/router /app/router
CMD ["/app/router"]
