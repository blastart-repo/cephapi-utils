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

func NewAdminClient(clusterName string) (*admin.API, error) {
	resp, err := GetClusterInfo(clusterName)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	c, err := admin.New(resp.GetEndpointurl(), resp.GetAccesskey(), resp.GetAccesssecret(), nil)
	if err != nil {
		return nil, err
	}
	return c, nil

}

func GetClusterInfo(clusterName string) (*proto.Cluster, error) {
	conn, err := ConnectgRPC()
	if err != nil {
		return nil, err
	}
	client := proto.NewClusterServiceClient(conn)

	clr, err := client.GetCluster(context.Background(), &proto.ClusterIn{Clustername: clusterName})
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

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
