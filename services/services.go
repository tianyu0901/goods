package services

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/context"
)

var (
	rpcClient client.Client
	ctx       context.Context
)

const (
	goodsRpc   = "go.micro.api.v1.goods"
	goodsTopic = "topic.go.micro.api.v1.goods"
)

func InitService(service micro.Service) {

	ctx = metadata.NewContext(context.Background(), map[string]string{})

	rpcClient = service.Client()
}
