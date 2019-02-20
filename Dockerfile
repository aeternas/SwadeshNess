FROM golang:1.11

ARG VERS

ENV VER $VERS

WORKDIR /go/src/github.com/aeternas/SwadeshNess
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM alpine:latest

ARG VERS

ENV VERSION $VERS

RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/aeternas/SwadeshNess .

HEALTHCHECK --interval=10s --timeout=10s --retries=3 CMD curl -sS 127.0.0.1:8080/v1/groups || exit 1

CMD ["./SwadeshNess", "--https"]
