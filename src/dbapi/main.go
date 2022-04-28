package main

import(
	"log"
	"net"
	"google.golang.org/grpc"
)


func main() {

	lis, err := net.Listern("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}	
(
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s", err)
	}
	
}
