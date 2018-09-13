// consignment-service/main.go
package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"golang.org/x/net/context"

	"github.com/micro/go-micro"
	pb "github.com/skyarkitekten/go-microservices/consignment-service/proto/consignment"
	userService "github.com/skyarkitekten/go-microservices/user-service/proto/user"
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

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, res interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth metadata found in request")
		}

		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		authClient := userService.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, res)
		return err
	}
}
