package goe

import (
	"net"
)

type (
	device struct {
		Network network
	}
	network struct{}
)

var (
	Device device
)

func (network) Mac() []string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	var macAddrArr []string
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrArr = append(macAddrArr, macAddr)
	}
	return macAddrArr
}
