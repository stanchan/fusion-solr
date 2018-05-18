package handler

import (
	"log"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	fusion "github.com/stanchan/fusion/proto/fusion"

	"context"
)

type Fusion struct{}

func (e *Fusion) Call(ctx context.Context, req *fusion.Request, rsp *fusion.Response) error {
	md, _ := metadata.FromContext(ctx)
	log.Printf("Received Fusion.Call request with metadata: %v", md)
	rsp.Msg = server.DefaultOptions().Id + ": Hello " + req.Name
	return nil
}

func (e *Fusion) Stream(ctx context.Context, stream server.Stream) error {
	log.Print("Executing streaming handler")
	req := &fusion.StreamingRequest{}

	// We just want to receive 1 request and then process here
	if err := stream.Recv(req); err != nil {
		log.Printf("Error receiving streaming request: %v", err)
		return err
	}

	log.Printf("Received Fusion.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Printf("Responding: %d", i)

		if err := stream.Send(&fusion.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (e *Fusion) PingPong(ctx context.Context, stream server.Stream) error {
	for {
		req := &fusion.Ping{}
		if err := stream.Recv(req); err != nil {
			return err
		}
		log.Printf("Got ping %v", req.Stroke)
		if err := stream.Send(&fusion.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
