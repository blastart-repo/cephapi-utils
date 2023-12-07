package cautils

import (
	"context"
	"errors"
	"fmt"

	"github.com/blastart-repo/cephapi-utils/proto"
	"github.com/ceph/go-ceph/rgw/admin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct {
	Address  string
	Port     string
	grpcConn *grpc.ClientConn
}

func NewConn(address, port string) *Connection {
	return &Connection{
		Address: address,
		Port:    port,
	}
}

func (c *Connection) ConnectgRPC() error {
	serverAddress := fmt.Sprintf("%s:%s", c.Address, c.Port)
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	c.grpcConn = conn
	return nil
}

func (c *Connection) GetClusterInfo(clusterName string) (*proto.Cluster, error) {
	if c.grpcConn == nil {
		return nil, errors.New("gRPC connection not initialized")
	}

	client := proto.NewClusterServiceClient(c.grpcConn)

	clr, err := client.GetCluster(context.Background(), &proto.ClusterIn{Clustername: clusterName})
	if err != nil {
		return nil, err
	}

	return clr, nil
}

func (c *Connection) NewAdminClient(clusterName string) (*admin.API, error) {
	if c.grpcConn == nil {
		return nil, errors.New("gRPC connection not initialized")
	}

	resp, err := c.GetClusterInfo(clusterName)
	if err != nil {
		return nil, err
	}

	client, err := admin.New(resp.GetEndpointurl(), resp.GetAccesskey(), resp.GetAccesssecret(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Connection) Close() error {
	if c.grpcConn != nil {
		return c.grpcConn.Close()
	}
	return nil
}
