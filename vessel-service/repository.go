package main

import (
	"gopkg.in/mgo.v2/bson"

	pb "github.com/skyarkitekten/go-microservices/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

// Repository is an interface for VesselRepository
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	Close()
}

// VesselRepository is a repository for the Vessel context
type VesselRepository struct {
	session *mgo.Session
}

// FindAvailable finds one vessel that meets the specification for shipping a container
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)

	if err != nil {
		return nil, err
	}

	return vessel, nil
}

// Create inserts a new vessel into the repository
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// Close closes the session to the datastore
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
