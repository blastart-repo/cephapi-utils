package cautils

import (
	"context"
	"github.com/blastart-repo/cephapi-utils/cluster"
	"github.com/ceph/go-ceph/rgw/admin"
	"google.golang.org/grpc"
)

func NewAdminClient(address, port, clusterName string, conn *grpc.ClientConn) (*admin.API, error) {
	client := cluster.NewClusterServiceClient(conn)
	clr, err := client.GetCluster(context.Background(), &cluster.ClusterIn{Clustername: clusterName})
	if err != nil {
		return nil, err
	}

	adm, err := admin.New(clr.GetEndpointurl(), clr.GetAccesskey(), clr.GetAccesssecret(), nil)
	if err != nil {
		return nil, err
	}

	return adm, nil
}
