package rpcserver

import (
	"context"
	"log"
	"net"

	dB "mobileapps/jobsserver/firebasedb"
	pb "mobileapps/jobsserver/protos/gen/jobslist"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedJobsServiceServer
}

func (s *server) JobsList(ctx context.Context, request *pb.EmptyRequest) (*pb.JobResponse, error) {
	resp := dB.GetJobs()
	return &pb.JobResponse{Jobs: resp}, nil
}

func Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterJobsServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
