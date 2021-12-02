package main

import (
	"fmt"
	"net"
	"os"

	"github.com/xanthangum1/gorilla_microservice/currency/data"
	protos "github.com/xanthangum1/gorilla_microservice/currency/protos/currency"
	"github.com/xanthangum1/gorilla_microservice/currency/server"

	hclog "github.com/hashicorp/go-hclog"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("Unable to generate rates", "error", err)
		os.Exit(1)
	}

	//creates a new gRPC server, use WithInsercure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	c := server.NewCurrency(rates, log)

	// register the currencty server
	protos.RegisterCurrencyServer(gs, c)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	// listen for requests
	gs.Serve(l)
}
