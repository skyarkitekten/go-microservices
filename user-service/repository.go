package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/skyarkitekten/go-microservices/user-service/proto/user"
)

// Repository is an interface for UserRepository
type Repository interface {
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	GetByEmailAndPassword(*pb.User) (*pb.User, error)
	Create(*pb.User) error
}

// UserRepository is a repository for the User context
type UserRepository struct {
	db *gorm.DB
}

// Get returns a User from the repository by Id
func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll returns all Users from the repository
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetByEmailAndPassword returns a User from the repository by email and password
func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Create inserts a new User into the repository
func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
