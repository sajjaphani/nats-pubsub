package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := runServer(nil)
	log.Println("NATS server is running")

	// Wait for termination signals to gracefully shut down the server
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	s.WaitForShutdown()
	log.Println("NATS server is stopped")
}
