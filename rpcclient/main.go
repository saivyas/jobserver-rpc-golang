package main

import (
	"context"
	"fmt"
	"log"

	pb "mobileapps/jobsserver/protos/gen/jobslist"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJobsServiceClient(conn)

	r, err := c.JobsList(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.GetJobs())

}
