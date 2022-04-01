package user

import (
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
