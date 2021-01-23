package main

import (
	"github.com/siliconvalley001/project1/cart/common"
	"github.com/siliconvalley001/project1/cart/handler"
	pb "github.com/siliconvalley001/project1/cart/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"log"
	res "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2/registry"
	_"github.com/siliconvalley001/project1/cart/model"
	"github.com/opentracing/opentracing-go"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"




)
var QPS=100

func main() {
	//配置中心
	consulConfig,err:=common.GetConsulConfig("127.0.0.1",8500,"")
	if err!=nil{
		log.Fatal(err)
	}
	//注册中心
	reg:=res.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"127.0.0.1:8500",
		}
		
	})
	//链路追踪
	open,io,err:=common.NewTrace("micro.cart","127.0.0.1:6831")
	if err!=nil{
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(open)
	// Create service
	srv := micro.NewService(
		micro.Name("cart"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8087"),
		micro.Registry(reg),
		micro.Config(consulConfig),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	srv.Init()


	// Register handler
	pb.RegisterCartHandler(srv.Server(), new(handler.Cart))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
