package user

import (
	"context"
	m "github.com/cheeeasy2501/go-user-data-service/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GRPCClient struct {
	Conn   *grpc.ClientConn
	Client UserServiceClient
}

func NewClient() *GRPCClient {
	return &GRPCClient{}
}

func (c *GRPCClient) Open() {
	//TODO CHECK LIB OR METHOD LAZYLOAD FOR GRPC CONN.
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Connection not open!", err)
	}

	c.Conn = conn
	c.Client = NewUserServiceClient(c.Conn)
}

func (c *GRPCClient) Close() {
	err := c.Conn.Close()
	if err != nil {
		log.Println("Connection not close!", err)
	}
}

func (c *GRPCClient) GetUserData(ctx context.Context, id uint64) (*m.User, error) {
	req := &GetUserRequest{
		Id: id,
	}

	data, err := c.Client.GetUserData(ctx, req)
	if err != nil {
		return nil, err
	}

	return &m.User{
		Id:        data.GetId(),
		Email:     data.GetEmail(),
		Password:  data.GetPassword(),
		FirstName: data.GetFirstName(),
		LastName:  data.GetLastName(),
		Active:    data.GetActive(),
	}, nil
}
