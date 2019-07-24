FROM golang:1.11.9-alpine3.9 as build
RUN apk add git musl-dev
WORKDIR /app
ADD . .
ENV GOFLAGS=-mod=vendor
RUN CGO_ENABLED=0 go build -o mini .

FROM alpine
COPY --from=build /app/mini /app/mini 
CMD ["/app/mini"]