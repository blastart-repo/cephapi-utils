package cautils

import (
	"context"
	"fmt"
	"github.com/blastart-repo/cephapi-utils/proto"
	"github.com/ceph/go-ceph/rgw/admin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct {
	Address string
	Port    string
}

func NewConn(address, port string) *Connection {
	return &Connection{
		Address: address,
		Port:    port,
	}
}

func (c *Connection) GetClusterInfo(clusterName string) (*proto.Cluster, error) {
	conn, err := c.ConnectgRPC()
	if err != nil {
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(conn)

	client := proto.NewClusterServiceClient(conn)

	clr, err := client.GetCluster(context.Background(), &proto.ClusterIn{Clustername: clusterName})
	if err != nil {
		return nil, err
	}

	return clr, nil
}

func (c *Connection) ConnectgRPC() (*grpc.ClientConn, error) {
	serverAddress := fmt.Sprintf("%s:%s", c.Address, c.Port)
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Connection) NewAdminClient(clusterName string) (*admin.API, error) {
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
