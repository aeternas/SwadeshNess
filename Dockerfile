FROM golang:1.12

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

RUN apk --no-cache add ca-certificates curl
COPY --from=0 /go/src/github.com/aeternas/SwadeshNess .

HEALTHCHECK CMD curl -sSk http://localhost:8079/dev/v1/groups || exit 1

CMD ["./SwadeshNess"]
