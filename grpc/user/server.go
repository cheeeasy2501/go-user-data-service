package user

import (
	"context"
	"errors"
	r "github.com/cheeeasy2501/go-user-data-service/internal/repository"
)

type GRPCServer struct {
	UnimplementedUserServiceServer
	ur *r.UserRepository
}

func NewGRPCServer(ur *r.UserRepository) *GRPCServer {
	return &GRPCServer{
		ur: ur,
	}
}
func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {}

// GetUserData ...
func (s *GRPCServer) GetUserData(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	response := &GetUserResponse{}

	id := req.GetId()

	if id <= 0 {
		return response, errors.New("Invalid id")
	}

	user, err := s.ur.GetById(id)

	if err != nil {
		return response, err
	}

	response.Id = user.Id
	response.Email = user.Email
	response.Password = user.Password
	response.FirstName = user.FirstName
	response.LastName = user.LastName
	response.Active = user.Active

	return response, nil
}
