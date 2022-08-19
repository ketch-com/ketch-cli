# Hacking

## Get the code

```shell
$ git clone https://github.com/ketch-com/ketch-cli
$ cd ketch-cli/
```

## Getting the development dependencies

```shell
go get -u ./...
```

## Building

You can build this project using Go:

```shell
go build ./...
```

You can also produce Linux binaries suitable for Docker Compose using the following:

```shell
./scripts/build.sh
```

## Distribution

The docker containers produced by this repository are contained in the `docker` folder.

## Updating dependencies

To update the dependencies, run the following:

```shell
rm go.sum
go get -u -t ./...
```

