package server

import (
	"review-service/internal/conf"
	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewRegistrar, NewGRPCServer, NewHTTPServer)


func NewRegistrar(conf *conf.Registry) registry.Registrar {
	// new consul client
	c := api.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	// new reg with consul client
	reg := consul.New(client, consul.WithHealthCheck(true))
	return reg
} 