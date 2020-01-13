package main

import (
	"blatt-4-consal/proto"
	"blatt-4-consal/services/cinemahallservice/microservice"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {
	//Create a new Service. Include name, version, address and context
	service := micro.NewService(
		micro.Name("cinemahall"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", 3002)),
		micro.Context(context.TODO()), //needed
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterCinemahallHandler(service.Server(), new(microservice.Cinemahall))
	fmt.Println("Cinemahall Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
