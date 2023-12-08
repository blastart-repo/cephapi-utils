package cautils

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectgRPC(address, port string) (*grpc.ClientConn, error) {
	serverAddress := fmt.Sprintf("%s:%s", address, port)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

/*
type Connection struct {
	grpcConn *grpc.ClientConn
}

// NewConn initializes a new connection and automatically connects to gRPC.
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

	c.grpcConn = conn
	return nil
}

func (c *Connection) Close() {
    if c.grpcConn != nil {
        c.grpcConn.Close()
    }
}

*/
