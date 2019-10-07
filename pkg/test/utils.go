package test

import (
	"fmt"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
	"net"
	"strings"
)

const (
	ArgsFormat = "IgnoreUnknown=1;K8S_POD_NAMESPACE=%s;K8S_POD_NAME=%s;K8S_POD_INFRA_CONTAINER_ID=%s"
)

func GenerateIPAMResult(cniVersion string, ipConfig []string, routeConfig []string, dnsConfig []string) *current.Result {
	ipamResult := &current.Result{}
	if cniVersion != "" {
		ipamResult.CNIVersion = cniVersion
	} else {
		ipamResult.CNIVersion = "0.3.1"
	}
	ipamResult.IPs = parseIPs(ipConfig)
	ipamResult.Routes = parseRoute(routeConfig)
	ipamResult.DNS = types.DNS{Nameservers: dnsConfig}
	return ipamResult
}

func parseRoute(routeConfig []string) []*types.Route {
	routes := make([]*types.Route, 0)
	for _, rt := range routeConfig {
		route := strings.Split(rt, ",")
		_, dst, _ := net.ParseCIDR(strings.Trim(route[0], " "))
		routeCfg := &types.Route{Dst: *dst}
		if len(route) == 2 {
			gw := net.ParseIP(strings.Trim(route[1], " "))
			routeCfg.GW = gw
		}
		routes = append(routes, routeCfg)
	}
	return routes
}

func parseIPs(ips []string) []*current.IPConfig {
	ipConfigs := make([]*current.IPConfig, 0)
	for _, ipc := range ips {
		configs := strings.Split(ipc, ",")
		addr := strings.Trim(configs[0], " ")
		gw := strings.Trim(configs[1], " ")
		version := strings.Trim(configs[2], " ")
		ipConfigs = append(ipConfigs, parseIPConfig(addr, gw, version))
	}
	return ipConfigs
}
func parseIPConfig(ipAddress string, gw string, version string) *current.IPConfig {
	ip, ipv4Net, _ := net.ParseCIDR(ipAddress)
	ipv4Net.IP = ip
	ipConfig := &current.IPConfig{Version: version, Address: *ipv4Net}
	if gw != "" {
		gateway := net.ParseIP(gw)
		ipConfig.Gateway = gateway
	}
	return ipConfig
}

func GenerateCNIArgs(podName string, podNamespace string, podInfraContainerID string) string {
	return fmt.Sprintf(ArgsFormat, podNamespace, podName, podInfraContainerID)
}

type DummyOVSConfigError struct {
	error
	timeout   bool
	temporary bool
}

func (e *DummyOVSConfigError) Timeout() bool {
	return e.timeout
}

func (e *DummyOVSConfigError) Temporary() bool {
	return e.temporary
}

func NewDummyOVSConfigError(errMsg string, temporary bool, timeout bool) *DummyOVSConfigError {
	return &DummyOVSConfigError{error: fmt.Errorf(errMsg), timeout: timeout, temporary: temporary}
}