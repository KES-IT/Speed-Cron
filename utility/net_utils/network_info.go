package net_utils

import (
	"fmt"
	"net"
)

type uNetworkInfo struct{}

var NetworkInfo = &uNetworkInfo{}

// GetMacAddress 获取Mac地址
func (u *uNetworkInfo) GetMacAddress() (internalIP, macAddress string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, iFace := range interfaces {
		if iFace.Flags&net.FlagLoopback == 0 && iFace.Flags&net.FlagUp != 0 {
			addressList, err := iFace.Addrs()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			for _, addr := range addressList {
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), iFace.HardwareAddr.String()
				}
			}
		}
	}
	return
}
