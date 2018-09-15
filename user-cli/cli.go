package main

import (
	"log"
	"os"

	"github.com/micro/go-micro"
	microClient "github.com/micro/go-micro/client"
	pb "github.com/skyarkitekten/go-microservices/user-service/proto/auth"
	"golang.org/x/net/context"
)

func main() {

	srv := micro.NewService(
		micro.Name("shippy.user-cli"),
		micro.Version("latest"),
	)

	srv.Init()

	client := pb.NewAuthClient("shippy.user", microClient.DefaultClient)

	name := "User McUserface"
	email := "user@mcuserface.com"
	password := "user123"
	company := "Acme"

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})

	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not get all users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error %v", email, err)
	}
	log.Printf("Authenticated with token: %s \n", authResponse.Token)

	os.Exit(0)
}
