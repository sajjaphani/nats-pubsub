# Pub/Sub with NATS.io

This repository is a [Go](https://golang.org/) project for experimentation on [NATS.io](https://nats.io/). It contains the following components:

- [server](components/server/): Starts NATS server
- [publisher](components/pub/): Implementation of a Go client that acts as a publisher
- [subscriber](components/sub/): Implementation of a Go client that acts as a subscriber

The server component will run the NATS server with the JetStream subsystem enabled. The [structured messages]((components/core/message/message.go)) are sent and received using the standard JSON encoding over JetStream.

## Setup

1. Seart the server
2. Start the subscriber
3. Start the publisher

Please refer to the following sections for instructions on how to start each of these individual components.


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
