package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/skyarkitekten/go-microservices/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	//contact the server and printout response
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})

	if err != nil {
		log.Fatalf("could not even: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
