package goutils

import (
	"errors"
	"fmt"
	"net/rpc"

	"github.com/hashicorp/consul/api"
)

var docker_ip string

func SetDockerIP(addr string) {
	docker_ip = addr
}

func GetConsulClient() (client *api.Client, err error) {

	if docker_ip == "" {
		err = errors.New("Please set the Docker IP address.")
		return
	}

	config := api.DefaultConfig()
	config.Address = docker_ip

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
		err = errors.New(fmt.Sprintf("Error: Can't get catalog in: %s", docker_ip))
		return
	}

	services, _, err := catalog.Service(serviceName, "", nil)
	if HasErrorAndPrintStack(err) {
		return
	}

	if len(services) == 0 {
		err = errors.New(fmt.Sprintf("Error: %s has no service: %s", docker_ip, serviceName))
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
