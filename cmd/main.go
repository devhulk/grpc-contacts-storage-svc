package main

import (
	"context"
	"log"
	"net"

	db "storage-service/internal/db"

	spb "storage-service/internal/storage"

	"google.golang.org/grpc"
)

const (
	port = ":5540"
)

type server struct {
	spb.UnimplementedStorageServiceServer
}

func (s *server) ReadDB(ctx context.Context, req *spb.ReadDBRequest) (*spb.ReadDBReply, error) {
	log.Printf("Recieved request to read contacts : %s", req.Message)
	// contacts = append(contacts, req.Contact)
	contacts := db.ReadData()
	return &spb.ReadDBReply{Contacts: contacts}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	spb.RegisterStorageServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve ContactServer: %v", err)
	}

}
