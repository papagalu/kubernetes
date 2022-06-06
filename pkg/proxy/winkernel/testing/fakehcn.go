package testing

import (
	"github.com/Microsoft/hcsshim/hcn"
)

const (
	testHostName      = "test-hostname"
	macAddress        = "00-11-22-33-44-55"
	clusterCIDR       = "192.168.1.0/24"
	destinationPrefix = "192.168.2.0/24"
	providerAddress   = "10.0.0.3"
	guid              = "123ABC"
)

type fakeHCN struct{}

func newFakeHCN() *fakeHCN {
	return &fakeHCN{}
}

func (HCN fakeHCN) getNetworkByName(networkName string) (*hcn.HostComputeNetwork, error) {
	return &hcn.HostComputeNetwork{
		Id: guid,
	}, nil
}
