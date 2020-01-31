package creds

import (
	"context"

	clusterapi "github.com/caspr-io/caspr/api/cluster"
	"github.com/caspr-io/caspr/internal/utils"
)

type CredentialsFetcher struct {
	ServiceAddress string
}

func (cf *CredentialsFetcher) FetchCredentials(id string) ([]byte, error) {
	grpcConn := utils.DialGrpc(cf.ServiceAddress)

	client := clusterapi.NewClusterServiceClient(grpcConn)

	clusterDetails, err := client.GetClusterDetails(context.Background(), &clusterapi.ClusterId{Id: id})
	if err != nil {
		return nil, err
	}

	return clusterDetails.GetCredentials(), nil
}
