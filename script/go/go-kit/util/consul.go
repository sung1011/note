package util

import (
	"log"
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
	reg := consulapi.AgentServiceRegistration{}
	SvcID = genSvcID()
	SvcAddress = "192.168.1.13"
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
