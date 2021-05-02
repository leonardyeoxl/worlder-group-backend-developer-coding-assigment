# README

## Build common docker image

```sh
$ cd common_build/alpine
$ docker build -t common-grpc-alpine .
```

## Docker compose build

```sh
$ docker-compose build
```

## Run docker compose

```sh
$ docker-compose up
```

## Generate protobuf go file
```sh
$ SRC_DIR=./
$ DST_DIR=$SRC_DIR
$ protoc \
  -I=$SRC_DIR \
  --go_out=plugins=grpc:$DST_DIR \
  $SRC_DIR/data/data.proto
```

## Test API

```sh
$ curl -X GET 'localhost:8080/data?ID1=2&ID2=A'
```

```sh
$ curl -X GET 'localhost:8080/data?start_timestamp=1619954581&end_timestamp=1619954585'
```

```sh
$ curl -X GET 'localhost:8080/data?ID1=2&ID2=A&start_timestamp=1619954581&end_timestamp=1619954585'
```