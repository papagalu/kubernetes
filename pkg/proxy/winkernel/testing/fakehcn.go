package testing

import (
	"github.com/Microsoft/hcsshim/hcn"
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

type FakeHCN struct {
	endpoints     []*hcn.HostComputeEndpoint
	loadbalancers []*hcn.HostComputeLoadBalancer
}

func NewFakeHCN() *FakeHCN {
	return &FakeHCN{}
}

func (HCN FakeHCN) getNetworkByName(networkName string) (*hcn.HostComputeNetwork, error) {
	return &hcn.HostComputeNetwork{
		Id:   guid,
		Name: networkName,
	}, nil

}

func (HCN FakeHCN) listEndpointsOfNetwork(networkId string) ([]hcn.HostComputeEndpoint, error) {
	var endpoints []hcn.HostComputeEndpoint
	for _, ep := range HCN.endpoints {
		if ep.HostComputeNetwork == networkId {
			endpoints = append(endpoints, *ep)
		}
	}
	return endpoints, nil
}

func (HCN FakeHCN) getEndpointByID(endpointId string) (*hcn.HostComputeEndpoint, error) {
	endpoint := &hcn.HostComputeEndpoint{}
	for _, ep := range HCN.endpoints {
		if ep.Id == endpointId {
			endpoint.Id = endpointId
			endpoint.Name = ep.Name
			endpoint.HostComputeNetwork = ep.HostComputeNetwork
			endpoint.Health = ep.Health
			endpoint.IpConfigurations = ep.IpConfigurations
		}
	}
	return endpoint, nil
}

func (HCN FakeHCN) listEndpoints() ([]hcn.HostComputeEndpoint, error) {

	var endpoints []hcn.HostComputeEndpoint
	for _, ep := range HCN.endpoints {
		endpoints = append(endpoints, *ep)
	}
	return endpoints, nil
}

func (HCN FakeHCN) getEndpointByName(endpointName string) (*hcn.HostComputeEndpoint, error) {
	endpoint := &hcn.HostComputeEndpoint{}
	for _, ep := range HCN.endpoints {
		if ep.Name == endpointName {
			endpoint.Id = ep.Id
			endpoint.Name = endpointName
			endpoint.HostComputeNetwork = ep.HostComputeNetwork
			endpoint.Health = ep.Health
			endpoint.IpConfigurations = ep.IpConfigurations
		}
	}
	return endpoint, nil
}

func (HCN FakeHCN) listLoadBalancers() ([]hcn.HostComputeLoadBalancer, error) {
	var loadbalancers []hcn.HostComputeLoadBalancer
	for _, lb := range HCN.loadbalancers {
		loadbalancers = append(loadbalancers, *lb)
	}
	return loadbalancers, nil
}

func (HCN FakeHCN) getLoadBalancerByID(loadBalancerId string) (*hcn.HostComputeLoadBalancer, error) {
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

func (HCN FakeHCN) createEndpoint(endpoint *hcn.HostComputeEndpoint) (*hcn.HostComputeEndpoint, error) {
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

func (HCN FakeHCN) createLoadBalancer(loadbalancer *hcn.HostComputeLoadBalancer) (*hcn.HostComputeLoadBalancer, error) {
	newLoadBalancer := &hcn.HostComputeLoadBalancer{
		Id:                   loadbalancer.Id,
		HostComputeEndpoints: loadbalancer.HostComputeEndpoints,
		SourceVIP:            loadbalancer.SourceVIP,
		Flags:                loadbalancer.Flags,
	}

	HCN.loadbalancers = append(HCN.loadbalancers, newLoadBalancer)

	return newLoadBalancer, nil
}
