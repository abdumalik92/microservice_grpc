package main

import (
	"context"
	desc "github.com/abdumalik92/microservice_grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
	nodeID  = 12
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect server: %v", err)
	}
	defer conn.Close()

	c := desc.NewNoteV1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: nodeID})
	if err != nil {
		log.Fatalf("failed to get note by id: %v", err)
	}
	log.Printf("%+v", r.GetNote())
}
