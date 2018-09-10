// consignment-service/main.go
package main

import (
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/skyarkitekten/go-microservices/consignment-service/proto/consignment"
	vesselProto "github.com/skyarkitekten/go-microservices/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	service := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", service.Client())

	service.Init()

	pb.RegisterShippingServiceHandler(service.Server(), &service{session, vesselClient})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
