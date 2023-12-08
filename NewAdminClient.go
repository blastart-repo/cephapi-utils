package cautils

import (
	"context"
	"github.com/blastart-repo/cephapi-utils/proto"
	"github.com/ceph/go-ceph/rgw/admin"
	"google.golang.org/grpc"
)

func NewAdminClient(address, port, clusterName string) (*admin.API, error) {
	conn, err := ConnectgRPC(address, port)
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

	adm, err := admin.New(clr.GetEndpointurl(), clr.GetAccesskey(), clr.GetAccesssecret(), nil)
	if err != nil {
		return nil, err
	}

	return adm, nil

}
