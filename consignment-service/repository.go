package main

import (
	pb "github.com/skyarkitekten/go-microservices/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	dbName               = "shippy"
	cosignmentCollection = "consignments"
)

// Repository is the interface for the ConsignmentRepository
type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

// ConsignmentRepository is the repository for consignments
type ConsignmentRepository struct {
	session *mgo.Session
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(cosignmentCollection)
}
