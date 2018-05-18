package main

import (
	"log"

	"context"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"github.com/stanchan/subscriber"

	fusion "github.com/micro/fusion/server/proto/fusion"
)

type Fusion struct{}

func (e *Fusion) Call(ctx context.Context, req *fusion.Request, rsp *fusion.Response) error {
	log.Print("Received Fusion.Call request")
	rsp.Msg = server.DefaultOptions().Id + ": Hello " + req.Name
	return nil
}

func (e *Fusion) Stream(ctx context.Context, req *example.StreamingRequest, stream fusion.Fusion_StreamStream) error {
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

func (e *Fusion) PingPong(ctx context.Context, stream fusion.Fusion_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Got ping %v", req.Stroke)
		if err := stream.Send(&fusion.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func main() {
	// optionally setup command line usage
	cmd.Init()

	// Initialise Server
	server.Init(
		server.Name("go.micro.srv.fusion"),
	)

	// Register Subscribers
	server.Subscribe(
		server.NewSubscriber(
			"topic.go.micro.srv.fusion",
			new(subscriber.Fusion),
		),
	)

	server.Subscribe(
		server.NewSubscriber(
			"topic.go.micro.srv.fusion",
			subscriber.Handler,
		),
	)

	// Register Handler
	fusion.RegisterFusionHandler(
		server.DefaultServer, new(Fusion),
	)

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
