package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
	pb "github.com/skyarkitekten/go-microservices/user-service/proto/user"
)

func main() {

	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pubsub := srv.Server().Options().Broker

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, pubsub})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
