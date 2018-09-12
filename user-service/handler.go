package main

import (
	pb "github.com/skyarkitekten/go-microservices/user-service/proto/user"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (s *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (s *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (s *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := s.tokenService.Encode(user)
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)
	err = s.repo.Create(req)
	if err != nil {
		return err
	}

	token, err := s.tokenService.Encode(req)
	if err != nil {
		return err
	}

	res.User = req
	res.Token = &pb.Token{Token: token}

	return nil
}

func (s *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
