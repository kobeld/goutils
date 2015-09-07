package goutils

import (
	"errors"
	"fmt"
	"net/rpc"

	"github.com/hashicorp/consul/api"
)

const (
	CONSUL_ADDRESS = "172.17.42.1:8500"
)

func GetConsulClient() (client *api.Client, err error) {
	config := api.DefaultConfig()
	config.Address = CONSUL_ADDRESS

	client, err = api.NewClient(config)
	if HasErrorAndPrintStack(err) {
		return
	}

	return
}

func GetConsulServiceAddress(serviceName string) (addr string, err error) {

	client, err := GetConsulClient()
	if HasErrorAndPrintStack(err) {
		return
	}

	catalog := client.Catalog()
	if catalog == nil {
		err = errors.New(fmt.Sprintf("Error: Can't get catalog in: %s", CONSUL_ADDRESS))
		return
	}

	services, _, err := catalog.Service(serviceName, "", nil)
	if HasErrorAndPrintStack(err) {
		return
	}

	if len(services) == 0 {
		err = errors.New(fmt.Sprintf("Error: %s has no service: %s", CONSUL_ADDRESS, serviceName))
		return
	}

	addr = fmt.Sprintf("%s:%+v", services[0].ServiceAddress, services[0].ServicePort)

	return
}

func GetServiceTcpClient(serviceName string) (client *rpc.Client, err error) {
	serviceAddress, err := GetConsulServiceAddress(serviceName)
	if HasErrorAndPrintStack(err) {
		return
	}

	client, err = rpc.Dial("tcp", serviceAddress)
	if HasErrorAndPrintStack(err) {
		return
	}

	return
}
