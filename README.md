# drone-circleci

[![Build Status](https://cloud.drone.io/api/badges/wesleimp/drone-circleci/status.svg)](https://cloud.drone.io/wesleimp/drone-circleci)
[![Go Report Card](https://goreportcard.com/badge/github.com/wesleimp/drone-circleci)](https://goreportcard.com/report/github.com/wesleimp/drone-circleci)

Drone plugin for trigger circleci builds. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```sh
$ go build
```

## Docker

Build the Docker image with the following commands:

```sh
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-circleci

$ docker build --rm -t wesleimp/drone-circleci -f docker/Dockerfile .
```

## Usage

Execute from the working directory:

```sh
$ docker run --rm \
    -e PLUGIN_USER=octocat \
    -e PLUGIN_REPO=hello-world \
    -e PLUGIN_BRANCH=master \
    -e PLUGIN_TOKEN=key-85c40a295fbd2dd0dfd47c38d8ae11d5e \
    wesleimp/drone-circleci
```
