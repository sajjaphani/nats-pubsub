# Pub/Sub with NATS.io

This repository is a [Go](https://golang.org/) project for experimentation on [NATS.io](https://nats.io/). It contains the following components:

- [server](components/server/): Starts NATS server
- [publisher](components/pub/): Implementation of a Go client that acts as a publisher
- [subscriber](components/sub/): Implementation of a Go client that acts as a subscriber

## Setup

1. Seart the server
2. Start the subscriber
3. Start the publisher

Please refer to the following sections for instructions on how to start each of these individual components.

**Note:** If the publisher is started before the subscriber, the messages will be dropped until a subscriber is available to consume them.

## Server

Run the following command to start the server:

```sh
cd components/server
go run main.go server.go
```

## Publisher

Run the following command to start the publisher:

```sh
cd components/pub
go run main.go
```

## Subscriber

Run the following command to start the subscriber:

```sh
cd components/sub
go run main.go
```
