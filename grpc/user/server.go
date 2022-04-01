package user

import (
	"context"
	"errors"
	r "github.com/cheeeasy2501/go-user-data-service/internal/repository"
	"github.com/cheeeasy2501/go-user-data-service/pkg/db"
)

type GRPCServer struct {
	UnimplementedUserServiceServer
}

//TODO create NewGRPCServer method and insert user repo
func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {}

// GetUserData ...
func (s *GRPCServer) GetUserData(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	cnf := db.NewConfig()
	mysql := db.NewMysql(cnf)
	err := mysql.Open()
	if err != nil {
		return nil, err
	}
	userRepo := r.NewUserRepo(mysql)
	response := &GetUserResponse{}

	id := req.GetId()

	if id <= 0 {
		return response, errors.New("Invalid id")
	}

	user, err := userRepo.GetById(id)

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
