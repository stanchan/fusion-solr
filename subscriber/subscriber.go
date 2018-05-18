package subscriber

import (
	"log"

	"context"
	fusion "github.com/stanchan/fusion-solr/proto/fusion"
)

type Fusion struct{}

func (e *Fusion) Handle(ctx context.Context, msg *fusion.Message) error {
	log.Print("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *fusion.Message) error {
	log.Print("Function Received message: ", msg.Say)
	return nil
}
