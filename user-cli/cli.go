package main

import (
	"log"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/skyarkitekten/go-microservices/user-service/proto/user"
	"golang.org/x/net/context"
)

func main() {

	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", microClient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}
			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
