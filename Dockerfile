FROM alpine:latest

RUN apk --no-cache add ca-certificates

EXPOSE 8080

ADD build/docker/go-api-sample /app/

ENTRYPOINT [ "/app/go-api-sample" ]
