package main

import (
	"context"
	"log"
	"time"

	storage "storage-service/internal/storage"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:5540"
	defaultName = "Gerald"
	message     = "this is a read request"
)

func listContacts(ctx context.Context, conn *grpc.ClientConn) {

	s := storage.NewStorageServiceClient(conn)

	r, err := s.ReadDB(ctx, &storage.ReadDBRequest{
		Message: "Read request from test-client",
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Contact List: %v", r.Contacts)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Contact client did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	listContacts(ctx, conn)

}
