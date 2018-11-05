FROM golang:1.10

ARG YANDEX_API_KEY=foo

ARG VERSION

ENV YANDEX_API_KEY=${YANDEX_API_KEY}
ENV VERS=$VERSION

WORKDIR /go/src/github.com/aeternas/SwadeshNess
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN ECHO $VERS
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM alpine:latest

ARG VERSION

ARG YANDEX_API_KEY=foo

ENV YANDEX_API_KEY=${YANDEX_API_KEY}
ENV VERSION=${VERSION}

RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/aeternas/SwadeshNess .

EXPOSE 8080

CMD ["./SwadeshNess"]
