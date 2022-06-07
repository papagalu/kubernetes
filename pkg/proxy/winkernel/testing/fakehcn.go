package testing

import (
	"github.com/Microsoft/hcsshim/hcn"
)

const (
	guid = "123ABC"
)

type HCN interface {
	GetNetworkByName(networkName string) (*hcn.HostComputeNetwork, error)
	ListEndpointsOfNetwork(networkId string) ([]hcn.HostComputeEndpoint, error)
	GetEndpointByID(endpointId string) (*hcn.HostComputeEndpoint, error)
	ListEndpoints() ([]hcn.HostComputeEndpoint, error)
	GetEndpointByName(endpointName string) (*hcn.HostComputeEndpoint, error)
	ListLoadBalancers() ([]hcn.HostComputeLoadBalancer, error)
	GetLoadBalancerByID(loadBalancerId string) (*hcn.HostComputeLoadBalancer, error)
	CreateEndpoint(endpoint *hcn.HostComputeEndpoint) (*hcn.HostComputeEndpoint, error)
	CreateLoadBalancer(loadbalancer *hcn.HostComputeLoadBalancer) (*hcn.HostComputeLoadBalancer, error)
}

type FakeHCN struct {
	endpoints     []*hcn.HostComputeEndpoint
	loadbalancers []*hcn.HostComputeLoadBalancer
}

func NewFakeHCN() *FakeHCN {
	return &FakeHCN{}
}

func (HCN FakeHCN) GetNetworkByName(networkName string) (*hcn.HostComputeNetwork, error) {
	return &hcn.HostComputeNetwork{
		Id:   guid,
		Name: networkName,
		Type: "overlay",
	}, nil

}

func (HCN FakeHCN) ListEndpointsOfNetwork(networkId string) ([]hcn.HostComputeEndpoint, error) {
	var endpoints []hcn.HostComputeEndpoint
	for _, ep := range HCN.endpoints {
		if ep.HostComputeNetwork == networkId {
			endpoints = append(endpoints, *ep)
		}
	}
	return endpoints, nil
}

func (HCN FakeHCN) GetEndpointByID(endpointId string) (*hcn.HostComputeEndpoint, error) {
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

func (HCN FakeHCN) ListEndpoints() ([]hcn.HostComputeEndpoint, error) {

	var endpoints []hcn.HostComputeEndpoint
	for _, ep := range HCN.endpoints {
		endpoints = append(endpoints, *ep)
	}
	return endpoints, nil
}

func (HCN FakeHCN) GetEndpointByName(endpointName string) (*hcn.HostComputeEndpoint, error) {
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

func (HCN FakeHCN) ListLoadBalancers() ([]hcn.HostComputeLoadBalancer, error) {
	var loadbalancers []hcn.HostComputeLoadBalancer
	for _, lb := range HCN.loadbalancers {
		loadbalancers = append(loadbalancers, *lb)
	}
	return loadbalancers, nil
}

func (HCN FakeHCN) GetLoadBalancerByID(loadBalancerId string) (*hcn.HostComputeLoadBalancer, error) {
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

func (HCN FakeHCN) CreateEndpoint(endpoint *hcn.HostComputeEndpoint) (*hcn.HostComputeEndpoint, error) {
	newEndpoint := &hcn.HostComputeEndpoint{
		Id:                 endpoint.Id,
		Name:               endpoint.Name,
		HostComputeNetwork: guid,
		IpConfigurations:   endpoint.IpConfigurations,
		MacAddress:         endpoint.MacAddress,
		Flags:              hcn.EndpointFlagsNone,
		SchemaVersion:      endpoint.SchemaVersion,
	}

	HCN.endpoints = append(HCN.endpoints, newEndpoint)

	return newEndpoint, nil
}

func (HCN FakeHCN) CreateRemoteEndpoint(endpoint *hcn.HostComputeEndpoint) (*hcn.HostComputeEndpoint, error) {
	newEndpoint := &hcn.HostComputeEndpoint{
		Id:                 endpoint.Id,
		Name:               endpoint.Name,
		HostComputeNetwork: guid,
		IpConfigurations:   endpoint.IpConfigurations,
		MacAddress:         endpoint.MacAddress,
		Flags:              hcn.EndpointFlagsRemoteEndpoint | endpoint.Flags,
		SchemaVersion:      endpoint.SchemaVersion,
	}

	HCN.endpoints = append(HCN.endpoints, newEndpoint)

	return newEndpoint, nil
}

func (HCN FakeHCN) CreateLoadBalancer(loadbalancer *hcn.HostComputeLoadBalancer) (*hcn.HostComputeLoadBalancer, error) {
	newLoadBalancer := &hcn.HostComputeLoadBalancer{
		Id:                   loadbalancer.Id,
		HostComputeEndpoints: loadbalancer.HostComputeEndpoints,
		SourceVIP:            loadbalancer.SourceVIP,
		Flags:                loadbalancer.Flags,
	}

	HCN.loadbalancers = append(HCN.loadbalancers, newLoadBalancer)

	return newLoadBalancer, nil
}
