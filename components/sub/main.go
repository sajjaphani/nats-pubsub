package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
	"github.com/sajjaphani/nats-pubsub/components/common/stream"
	"github.com/sajjaphani/nats-pubsub/components/core/message"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL,
		// Decoding errors will be handled here
		nats.ErrorHandler(func(nc *nats.Conn, s *nats.Subscription, err error) {
			if s != nil {
				log.Printf("Async error in %q/%q: %v", s.Subject, s.Queue, err)
			} else {
				log.Printf("Async error outside subscription: %v", err)
			}
		}))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create JetStream Context
	js, err := stream.InitJetStream(nc)
	if err != nil {
		log.Fatal(err)
	}

	// Create a channel to listen for termination signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Subscribe to the topic until termination signal is received
	_, err = js.Subscribe(stream.StreamSubjectMessageCreated, func(m *nats.Msg) {
		var msg message.Message
		err = json.Unmarshal(m.Data, &msg)
		if err != nil {
			log.Printf("Message decode error: %v", err)
		} else {
			log.Printf("Received -> Id: %v - Text: %s\n", msg.Id, msg.Text)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	// Wait for termination signal
	<-signalCh
	log.Println("Termination signal received. Subscriber exiting...")
}
