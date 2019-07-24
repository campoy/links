FROM golang:1.12
WORKDIR /app
ENV GOFLAGS=-mod=vendor
ADD . .
RUN CGO_ENABLED=0 go build -o links .

FROM alpine
WORKDIR /app
COPY --from=0 /app/links /app/links
COPY --from=0 /app/home.html /app/home.html
CMD ["/app/links"]
