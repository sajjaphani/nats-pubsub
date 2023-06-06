package stream

import (
	"github.com/nats-io/nats.go"
)

func InitJetStream(nc *nats.Conn) (nats.JetStreamContext, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	err = addStreamIfNotExists(js)
	if err != nil {
		return nil, err
	}

	return js, nil
}

func addStreamIfNotExists(js nats.JetStreamContext) error {
	_, err := js.StreamInfo(StreamName)
	if err != nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     StreamName,
			Subjects: []string{StreamSubject},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
