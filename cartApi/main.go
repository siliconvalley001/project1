package main

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	"github.com/siliconvalley001/project1/cart/proto"
	"fmt"
	"net"
	"net/http"
	"github.com/micro/go-micro/v2/registry"
	res "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/siliconvalley001/project1/cartApi/common"

	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"log"
)

func main() {

	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "")
	if err != nil {
		log.Fatal(err)
	}
	//注册中心
	reg := res.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}

	})
	//链路追踪
	open, io, err := common.NewTrace("micro.cartapi", "127.0.0.1:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(open)

	hystrixstramHandler := hystrix.NewStreamHandler()
	hystrixstramHandler.Start()
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9096"), hystrixstramHandler)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Create service
	srv := micro.NewService(
		micro.Name("cartapi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		micro.Registry(reg),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(NewClientHystrixWrapper()),
		//负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
		micro.Config(consulConfig),
	)
	proto.NewCartService("micro.cartapi.service", srv.Client())

	srv.Init()

	// Register handlere

	//pb.RegisterCartApiHandler(srv.Server(), new(handler.CartApi))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

type clientWrap struct {
	client.Client
}

func (c *clientWrap) Call(ctx context.Context, resquset client.Request, response interface{}, opt ...client.CallOption) error {
	c.Client.Call(ctx, resquset, response, opt...)

	return hystrix.Do(resquset.Service()+"."+resquset.Endpoint(), func() error {
		//run 正常执行
		fmt.Println(resquset.Service() + "." + resquset.Endpoint())
		return c.Client.Call(ctx, resquset, response, opt...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}
func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrap{i}
	}
}
