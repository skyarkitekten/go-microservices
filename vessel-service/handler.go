package main

import (
	pb "github.com/skyarkitekten/go-microservices/vessel-service/proto/vessel"

	"gopkg.in/mgo.v2"

	"golang.org/x/net/context"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	err := repo.Create(req)
	if err != nil {
		return err
	}

	res.Vessel = req
	res.Create = true
	return nil
}
