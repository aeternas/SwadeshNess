FROM golang:1.10

ARG YANDEX_API_KEY=foo

ENV YANDEX_API_KEY=${YANDEX_API_KEY}

WORKDIR /go/src/github.com/aeternas/SwadeshNess
COPY . .

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

FROM alpine:latest

COPY --from=0 /go/src/github.com/aeternas/SwadeshNess .

CMD ["./SwadeshNess"]
