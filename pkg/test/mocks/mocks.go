//go:generate mockgen -destination ovsconfig_mock.go -package=mocks -mock_names OVSBridgeClient=MockOVSBridgeClient okn/pkg/ovs/ovsconfig OVSBridgeClient
//go:generate mockgen -destination ipam_mock.go -package=mocks okn/pkg/agent/cniserver/ipam IPAMDriver
//go:generate mockgen -destination ofclient_mock.go -package=mocks -mock_names Client=MockOFClient okn/pkg/agent/openflow Client

package mocks