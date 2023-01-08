package main

import (
	"context"
	"dataAnalyzerFile/frequentationPB"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Hello(ctx context.Context, request *frequentationPB.FrequentationResponse) (*frequentationPB.FrequentationResponse, error) {
	response := &frequentationPB.FrequentationResponse.Value{
		commune_libellemin: "Hello "
	}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	frequentationPB.FrequentationResponse(server)

	s.Serve(lis)
}
