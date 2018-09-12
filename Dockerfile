FROM golang:1.10

ARG YANDEX_API_KEY=foo

ENV YANDEX_API_KEY=${YANDEX_API_KEY}

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get -d -v ./...
RUN go install -v ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main

FROM alpine:latest

COPY --from=0 /app .

CMD ["/app/main"]
