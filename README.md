# Zonked

## Getting Started

`go run main.go` or use Docker somehow.

## Docker for Dev

Run `docker run --publish 8080:8080 --name test --rm zonked` from the command line. Visit the app at [localhost:8080](localhost:8080).

To rebuild the image, run `docker build -t zonked .`.

_built using Gin_
