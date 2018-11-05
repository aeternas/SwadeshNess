FROM golang:1.10

ARG VERS

ENV VER $VERS

WORKDIR /go/src/github.com/aeternas/SwadeshNess
COPY . .

RUN echo $VER
RUN go get -d -v ./...
RUN go install -v ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM alpine:latest

ENV VERSION $VER

RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/aeternas/SwadeshNess .

EXPOSE 8080

CMD ["./SwadeshNess"]
