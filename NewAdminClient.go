package cautils

import (
	"context"
	"fmt"

	"github.com/blastart-repo/cephapi-utils/proto"
	"github.com/ceph/go-ceph/rgw/admin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Cluster struct {
	ClusterName  string `json:"cluster_name" gorm:"unique,primaryKey"`
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
	EndpointURL  string `json:"endpoint_url" gorm:"unique"`
}

type NewClient struct {
	client *grpc.ClientConn
}

type SrvConnData struct {
	srvName string
	srvPort string
}

func (c *NewClient) NewAdminClient(clusterName string) (*admin.API, error) {
	resp, err := c.GetClusterInfo(clusterName)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	client, err := admin.New(resp.GetEndpointurl(), resp.GetAccesskey(), resp.GetAccesssecret(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil

}

func (c *NewClient) GetClusterInfo(clusterName string) (*proto.Cluster, error) {
	client := proto.NewClusterServiceClient(c.client)

	clr, err := client.GetCluster(context.Background(), &proto.ClusterIn{Clustername: clusterName})
	if err != nil {
		return nil, err
	}

	return clr, nil
}

func ConnectgRPC() (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("data-service:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
