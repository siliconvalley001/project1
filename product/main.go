package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/opentracing/opentracing-go"
	"github.com/siliconvalley001/project1/product/common"
	"github.com/siliconvalley001/project1/product/config"
	"github.com/siliconvalley001/project1/product/handler"
	"log"
	ex "github.com/siliconvalley001/project1/product/proto"



)

func main() {
	//配置中心
	consulconfig,err:=config.GetConsulConfig("127.0.0.1",8500,"/lizhengda/project/config")
	if err!=nil{
		log.Fatal(err)
	}
	//注册中心
	registry:=consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs=[]string{
				"127.0.0.1:8500",
			}

		},
	)
	//链路追踪
	t,io,err:=common.NewTrace("micro.product.jaeger","127.0.0.1:6831")
	if err!=nil{
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)




	// Create service
	srv := micro.NewService(
		micro.Name("product"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Config(consulconfig),
		//绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),

	)

	// Register handler
	ex.RegisterProductHandler(srv.Server(), new(handler.Product))

	// Run service
	if err := srv.Run(); err != nil {

	}
}
