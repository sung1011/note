package util

import (
	"log"
	"net"
	"strconv"

	"github.com/google/uuid"
	consulapi "github.com/hashicorp/consul/api"
)

var ConsulClient *consulapi.Client

func init() {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"

	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
}

var SvcName string
var SvcPort string
var SvcID string
var SvcAddress string

func RegService() error {
	localIP, err := getLocalIPv4s()
	if err != nil {
		log.Fatal(err)
	}
	reg := consulapi.AgentServiceRegistration{}
	SvcID = genSvcID()
	SvcAddress = localIP[0]
	// SvcAddress = "127.0.0.1"

	reg.ID = SvcID
	reg.Name = SvcName
	reg.Address = SvcAddress
	sp, _ := strconv.Atoi(SvcPort)
	reg.Port = sp
	reg.Tags = []string{"primary"}
	reg.Check = &consulapi.AgentServiceCheck{
		HTTP:     "http://" + SvcAddress + ":" + SvcPort + "/health",
		Interval: "5s",
	}
	return ConsulClient.Agent().ServiceRegister(&reg)
}

func UnregService() error {
	return ConsulClient.Agent().ServiceDeregister(SvcID)
}

func genSvcID() string {
	return uuid.New().String()
}

func getLocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, nil
}
