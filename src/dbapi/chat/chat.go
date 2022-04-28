package chat

import (
	"log"
	"golang.org/x/net/context"
)

type Server struct {
	
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message,error) {

	log.Printf("Received Message body from client carName: %s", message.carName)
	log.Printf("Received Message body from client carModule: %s", message.carModule)
	log.Printf("Received Message body from client carYear: %s", message.carYear)

	return &Message {carName: "Cassandra", carModule: "Fiat", carYear: "1983"}, nil
}