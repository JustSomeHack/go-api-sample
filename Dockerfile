FROM alpine:latest

RUN apk --no-cache add ca-certificates

EXPOSE 8080

COPY build/docker/go-api-sample /usr/local/bin/go-api-sample
