# go_playground

This project holds experimentation with tools surround go projects, and has grown into a template that can be used to bootstrap new projects.

In particular, gRPC projects with a go server implementaiton. It also produces go, java, and npm packages that clients can consume to integrate with the server.

## Run
```sh
go run ./service
```

## Build
```sh
podman build .
```
or
```sh
docker build .
```

## Run via docker-compose
```sh
podman compose up
```
or
```sh
docker compose up
```

## cURL
```sh
buf curl \
  --schema proto \
  --data '{"name": "You"}' \
  http://localhost:8080/mpataki.go_playground.proto.greeting.v1.GreetingService/SayHello
```
