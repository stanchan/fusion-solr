package main

import (
	"log"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"github.com/stanchan/fusion-solr/handler"
	"github.com/stanchan/fusion-solr/subscriber"
)

func main() {
	// optionally setup command line usage
	cmd.Init()

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
