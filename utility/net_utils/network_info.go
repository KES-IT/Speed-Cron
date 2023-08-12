package net_utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcache"
	"net"
)

type uNetworkInfo struct{}

var NetworkInfo = &uNetworkInfo{}

// GetMacAddress 获取Mac地址
func (u *uNetworkInfo) GetMacAddress() (internalIP, macAddress string) {

	// GitHub测试状态
	githubTestStatus := gcache.MustGet(context.Background(), "GitHubTestStatus").Bool()
	if githubTestStatus {
		return "172.0.0.1", "f0:2f:74:4f:61:ce"
	}

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
