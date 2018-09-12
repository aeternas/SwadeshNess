FROM golang:1.10

ARG YANDEX_API_KEY=foo

ENV YANDEX_API_KEY=${YANDEX_API_KEY}

WORKDIR /go/src/github.com/aeternas/SwadeshNess
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

FROM golang:alpine

COPY --from=0 /go/src/github.com/aeternas/SwadeshNess .

CMD ["./SwadeshNess"]
