package main

import (
	"fmt"

	server "github.com/nats-io/nats-server/v2/server"
)

func defaultOptions() *server.Options {
	return &server.Options{
		Host:   "127.0.0.1",
		Port:   4222,
		NoLog:  false,
		NoSigs: false,
		Debug:  true,
		Trace:  true,
	}
}

func runServer(opts *server.Options) *server.Server {
	if opts == nil {
		opts = defaultOptions()
	}

	s, err := server.NewServer(opts)
	if err != nil || s == nil {
		panic(fmt.Sprintf("No NATS Server object returned: %v", err))
	}

	if !opts.NoLog {
		s.ConfigureLogger()
	}

	// Run server in Go routine.
	s.Start()

	return s
}
