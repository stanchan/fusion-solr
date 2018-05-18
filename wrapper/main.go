package main

import (
	"log"

	"context"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"github.com/stanchan/fusion-solr/handler"
	"github.com/stanchan/fusion-solr/subscriber"
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[Log Wrapper] Before serving request method: %v", req.Method())
		err := fn(ctx, req, rsp)
		log.Printf("[Log Wrapper] After serving request")
		return err
	}
}

func logSubWrapper(fn server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, req server.Message) error {
		log.Printf("[Log Sub Wrapper] Before serving publication topic: %v", req.Topic())
		err := fn(ctx, req)
		log.Printf("[Log Sub Wrapper] After serving publication")
		return err
	}
}

func main() {
	// optionally setup command line usage
	cmd.Init()

	md := server.DefaultOptions().Metadata
	md["datacenter"] = "local"

	server.DefaultServer = server.NewServer(
		server.WrapHandler(logWrapper),
		server.WrapSubscriber(logSubWrapper),
		server.Metadata(md),
	)

	// Initialise Server
	server.Init(
		server.Name("go.micro.srv.fusion"),
	)

	// Register Handlers
	server.Handle(
		server.NewHandler(
			new(handler.Fusion),
		),
	)

	// Register Subscribers
	if err := server.Subscribe(
		server.NewSubscriber(
			"topic.go.micro.srv.fusion",
			new(subscriber.Fusion),
		),
	); err != nil {
		log.Fatal(err)
	}

	if err := server.Subscribe(
		server.NewSubscriber(
			"topic.go.micro.srv.fusion",
			subscriber.Handler,
		),
	); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
