package main

import (
	"strings"
	"testing"

	pb "github.com/skyarkitekten/go-microservices/user-service/proto/auth"
)

var (
	user = &pb.User{
		Id:    "abc123",
		Email: "test@test.com",
	}
)

type MockRepo struct{}

func (repo *MockRepo) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	return users, nil
}

func (repo *MockRepo) Get(id string) (*pb.User, error) {
	var user *pb.User
	return user, nil
}

func (repo *MockRepo) Create(user *pb.User) error {
	return nil
}

func (repo *MockRepo) GetByEmail(email string) (*pb.User, error) {
	var user *pb.User
	return user, nil
}

func newInstance() Authable {
	repo := &MockRepo{}
	return &TokenService{repo}
}

func TestCanCreateToken(t *testing.T) {
	srv := newInstance()
	token, err := srv.Encode(user)
	if err != nil {
		t.Fail()
	}

	if token == "" {
		t.Fail()
	}

	if len(strings.Split(token, ".")) != 3 {
		t.Fail()
	}
}

func TestCanDecodeToken(t *testing.T) {
	srv := newInstance()
	token, err := srv.Encode(user)
	if err != nil {
		t.Fail()
	}
	claims, err := srv.Decode(token)
	if err != nil {
		t.Fail()
	}
	if claims.User == nil {
		t.Fail()
	}
	if claims.User.Email != user.Email {
		t.Fail()
	}
}

func TestThrowsErrorIfTokenInvalid(t *testing.T) {
	srv := newInstance()
	_, err := srv.Decode("nope.nope.nope")
	if err == nil {
		t.Fail()
	}
}
