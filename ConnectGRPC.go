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
