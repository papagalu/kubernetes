package testing

import (
	"github.com/Microsoft/hcsshim/hcn"
	"k8s.io/kubernetes/pkg/proxy/winkernel"
)

const (
	guid = "123ABC"
)

type HCN interface {
	getNetworkByName(networkName string) (*hcn.HostComputeNetwork, error)
	listEndpointsOfNetwork(networkId string) ([]hcn.HostComputeEndpoint, error)
	getEndpointByID(endpointId string) (*hcn.HostComputeEndpoint, error)
	listEndpoints() ([]hcn.HostComputeEndpoint, error)
	getEndpointByName(endpointName string) (*hcn.HostComputeEndpoint, error)
	listLoadBalancers() ([]hcn.HostComputeLoadBalancer, error)
	getLoadBalancerByID(loadBalancerId string) (*hcn.HostComputeLoadBalancer, error)
	createEndpoint(endpoint *hcn.HostComputeEndpoint) (*hcn.HostComputeEndpoint, error)
	createLoadBalancer(loadbalancer *hcn.HostComputeLoadBalancer) (*hcn.HostComputeLoadBalancer, error)
}

type fakeHCN struct {
	endpoints     []*hcn.HostComputeEndpoint
	loadbalancers []*hcn.HostComputeLoadBalancer
}

func newFakeHCN() *fakeHCN {
	return &fakeHCN{}
}

func (HCN fakeHCN) getNetworkByName(networkName string) (*hcn.HostComputeNetwork, error) {
	return &hcn.HostComputeNetwork{
		Id:   guid,
		Name: networkName,
		Type: winkernel.NETWORK_TYPE_OVERLAY,
	}, nil
}

func (HCN fakeHCN) listEndpointsOfNetwork(networkId string) ([]hcn.HostComputeEndpoint, error) {
	return nil, nil
}

func (HCN fakeHCN) getEndpointByID(endpointId string) (*hcn.HostComputeEndpoint, error) {
	ipConfig := &hcn.IpConfig{
		IpAddress: epIpAddress,
	}
	return &hcn.HostComputeEndpoint{
		Id:               endpointId,
		IpConfigurations: []hcn.IpConfig{*ipConfig},
		MacAddress:       epMacAddress,
		SchemaVersion: hcn.SchemaVersion{
			Major: 2,
			Minor: 0,
		},
	}, nil
}

func (HCN fakeHCN) listEndpoints() ([]hcn.HostComputeEndpoint, error) {

	var endpoints []hcn.HostComputeEndpoint
	for _, ep := range HCN.endpoints {
		endpoints = append(endpoints, *ep)
	}
	return endpoints, nil
}

func (HCN fakeHCN) getEndpointByName(endpointName string) (*hcn.HostComputeEndpoint, error) {
	loadbalancer := &hcn.HostComputeLoadBalancer{}
	for _, lb := range HCN.loadbalancers {
		if lb.Id == loadBalancerId {
			loadbalancer.Id = loadBalancerId
			loadbalancer.Flags = lb.Flags
			loadbalancer.HostComputeEndpoints = lb.HostComputeEndpoints
			loadbalancer.SourceVIP = lb.SourceVIP
		}
	}
	return loadbalancer, nil
}

func (HCN fakeHCN) listLoadBalancers() ([]hcn.HostComputeLoadBalancer, error) {
	var loadbalancers []hcn.HostComputeLoadBalancer
	for _, lb := range HCN.loadbalancers {
		loadbalancers = append(loadbalancers, *lb)
	}
	return loadbalancers, nil
}

func (HCN fakeHCN) getLoadBalancerByID(loadBalancerId string) (*hcn.HostComputeLoadBalancer, error) {
	loadbalancer := &hcn.HostComputeLoadBalancer{}
	for _, lb := range HCN.loadbalancers {
		if lb.Id == loadBalancerId {
			loadbalancer.Id = loadBalancerId
			loadbalancer.Flags = lb.Flags
			loadbalancer.HostComputeEndpoints = lb.HostComputeEndpoints
			loadbalancer.SourceVIP = lb.SourceVIP
		}
	}
	return loadbalancer, nil
}

func (HCN fakeHCN) createEndpoint(endpoint *hcn.HostComputeEndpoint) (*hcn.HostComputeEndpoint, error) {
	newEndpoint := &hcn.HostComputeEndpoint{
		Id:                 endpoint.Id,
		Name:               endpoint.Name,
		HostComputeNetwork: guid,
		IpConfigurations:   endpoint.IpConfigurations,
		MacAddress:         endpoint.MacAddress,
		Flags:              endpoint.Flags,
		SchemaVersion:      endpoint.SchemaVersion,
	}

	HCN.endpoints = append(HCN.endpoints, newEndpoint)

	return newEndpoint, nil
}

func (HCN fakeHCN) createLoadBalancer(loadbalancer *hcn.HostComputeLoadBalancer) (*hcn.HostComputeLoadBalancer, error) {
	newLoadBalancer := &hcn.HostComputeLoadBalancer{
		Id:                   loadbalancer.Id,
		HostComputeEndpoints: loadbalancer.HostComputeEndpoints,
		SourceVIP:            loadbalancer.SourceVIP,
		Flags:                loadbalancer.Flags,
	}

	HCN.loadbalancers = append(HCN.loadbalancers, newLoadBalancer)

	return newLoadBalancer, nil
}
