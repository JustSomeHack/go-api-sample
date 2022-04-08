# go-api-sample
Sample Go API Project

[![Build Status](https://drone.onebytedata.net/api/badges/JustSomeHack/go-api-sample/status.svg)](https://drone.onebytedata.net/JustSomeHack/go-api-sample)

## Usage

Run all tests `go test -p 1 -coverprofile=coverage.txt ./...`

Run benchmarks against the API `cd internal/controllers && go test -benchmem -run="^#" -bench . && cd ../..`

Run benchmarks against the Database `cd internal/services && go test -benchmem -run="^#" -bench . && cd ../..`

Install swagger spec generate tool `go install github.com/swaggo/swag/cmd/swag@latest`

Generate swagger spec `swag init -d cmd/server`

Build API `go build -v -a -o build/docker/go-api-sample cmd/server/main.go`

## Runtime Environment Variables

`CONNECTION_STRING="postgresql://root@cockroachdb:26257/animals?sslmode=disable"`

