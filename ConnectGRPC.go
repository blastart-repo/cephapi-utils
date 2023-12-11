package cautils

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct {
	GrpcConn *grpc.ClientConn
}

// NewConn initializes a new connectionGRPC and automatically connects to gRPC.
func NewConn(address, port string) (*Connection, error) {
	conn := &Connection{}

	if err := conn.ConnectgRPC(address, port); err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Connection) ConnectgRPC(address, port string) error {
	serverAddress := fmt.Sprintf("%s:%s", address, port)
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	c.GrpcConn = conn
	return nil
}

func (c *Connection) Close() error {
	if c.GrpcConn != nil {
		err := c.GrpcConn.Close()
		if err != nil {
			return fmt.Errorf("error closing gRPC connection: %w", err)
		}
	}
	return nil
}
