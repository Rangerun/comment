package data

import (
	"context"
	v1 "review-b/api/review/v1"

	"review-b/internal/conf"

	consul "github.com/go-kratos/consul/registry"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/hashicorp/consul/api"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDiscovery, NewReviewServiceClient, NewData, NewBusinessRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	// 嵌入一个client 端调用服务端的回复评价接口
	rc v1.ReviewClient
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rc v1.ReviewClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rc: rc, log: log.NewHelper(logger)}, cleanup, nil
}


func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := api.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	client, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	// new dis with consul client
	dis := consul.New(client)
	return dis
}


func NewReviewServiceClient(d registry.Discovery) v1.ReviewClient{
	conn, err := grpc.DialInsecure(
        context.Background(),
        //grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithEndpoint("discovery:///review.service"),
		grpc.WithDiscovery(d),
    )

	if err != nil {
		panic(err)
	}
	return v1.NewReviewClient(conn)
}
