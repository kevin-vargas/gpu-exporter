package consul

import (
	"errors"
	"strconv"

	"github.com/hashicorp/consul/api"
)

const (
	id = "nvidia_gpu_exporter"
)

var (
	tags = []string{"microservice", "golang"}
)

type discovery struct {
	*api.Client
	address        string
	port           int
	healthEndpoint string
	id             string
}

func New(
	consulAddress,
	address,
	healthEndpoint,
	port string,
) (*discovery, error) {

	conf := &api.Config{
		Address: consulAddress,
	}

	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	if len(port) == 0 {
		return nil, errors.New("invalid port")
	}

	portInt, err := strconv.Atoi(port[1:])
	if err != nil {
		return nil, err
	}
	return &discovery{
		healthEndpoint: healthEndpoint,
		address:        address,
		port:           portInt,
		Client:         client,
		id:             id,
	}, nil
}

func (d *discovery) Registry() error {
	serviceDefinition := &api.AgentServiceRegistration{
		ID:      d.id,
		Name:    d.id + "_ms",
		Port:    d.port,
		Address: d.address,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			Interval: "30s",
			Timeout:  "60s",
			HTTP:     d.healthEndpoint,
		},
	}
	if err := d.Agent().ServiceRegister(serviceDefinition); err != nil {
		return err
	}
	return nil
}
