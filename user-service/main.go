package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
	pb "github.com/skyarkitekten/go-microservices/user-service/proto/auth"
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

	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
