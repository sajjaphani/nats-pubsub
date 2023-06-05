package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	// Publish messages every second until termination signal is received
	for {
		select {
		case <-signalCh:
			log.Println("Termination signal received. Publisher exiting...")
			return
		default:
			msg := "Hello NATS!"
			err = nc.Publish(subject, []byte(msg))
			if err != nil {
				log.Println("Publish error:", err)
			} else {
				log.Println("Published:", msg)
			}

			time.Sleep(5 * time.Second)
		}
	}
}
