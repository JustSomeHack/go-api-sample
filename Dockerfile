FROM golang:latest as builder

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN mkdir /src

COPY . /src/

WORKDIR /src

RUN swag init

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o build/docker/go-api-sample

FROM alpine:latest

RUN apk --no-cache add ca-certificates

EXPOSE 8080

COPY --from=builder /src/build/docker/go-api-sample /usr/local/bin/

CMD [ "/usr/local/bin/go-api-sample" ]