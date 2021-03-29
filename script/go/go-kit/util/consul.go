package util

import (
	"log"

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

func RegService() error {
	reg := consulapi.AgentServiceRegistration{}
	reg.ID = "kkk"
	reg.Name = "kkk"
	reg.Address = "192.168.1.13"
	reg.Port = 8001
	reg.Check = &consulapi.AgentServiceCheck{
		HTTP:     "http://192.168.1.13:8001/health",
		Interval: "5s",
	}
	return ConsulClient.Agent().ServiceRegister(&reg)
}

func UnregService() error {
	return ConsulClient.Agent().ServiceDeregister("kkk")
}
