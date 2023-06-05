package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create a channel to listen for termination signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	subject := "hello.nats"
	// Subscribe to the topic until termination signal is received
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		log.Println("Received:", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Wait for termination signal
	<-signalCh
	log.Println("Termination signal received. Subscriber exiting...")
}
