package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sajjaphani/nats-pubsub/components/common/stream"
	"github.com/sajjaphani/nats-pubsub/components/core/message"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create JetStream Context
	js, err := stream.InitJetStream(nc)
	if err != nil {
		log.Fatal(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	// Create a channel to listen for termination signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Publish messages every second until termination signal is received
	for {
		select {
		case <-signalCh:
			log.Println("Termination signal received. Publisher exiting...")
			return
		default:
			m := message.NewMessage("Hello NATS!")
			msg, err := json.Marshal(m)
			if err != nil {
				log.Println(err)
				continue
			}

			_, err = js.Publish(stream.StreamSubjectMessageCreated, msg)
			if err != nil {
				log.Println("Publish error:", err)
			} else {
				log.Println("Published:", m.String())
			}

			time.Sleep(5 * time.Second)
		}
	}
}
