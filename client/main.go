package main

import (
	"log"

	pb "github.com/budiuno/go-simple-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	people := []string{"John", "Paul", "Doe"}
	names := &pb.NameList{
		Names: people,
	}

	// callSayHello(client)
	//callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callSayHelloBiDirectionalStream(client, names)

}
