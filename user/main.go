package main

import (
	"github.com/micro/go-micro/v2"
	_"github.com/siliconvalley001/project1/user/model"


	handler1 "github.com/siliconvalley001/project1/user/handler"

	"fmt"
	ex "github.com/siliconvalley001/project1/user/proto"
)

func main() {

	// Create service
	srv:=micro.NewService(
		micro.Name("micro.user"),
		micro.Version("latest"),
		)
	srv.Init()
	// Register handler
	//pb.RegisterUserHandler(srv.Server(), new(handler.User))

	ex.RegisterUserHandler(srv.Server(),new(handler1.Handler))
	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
