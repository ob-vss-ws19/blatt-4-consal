package main

import (
	"blatt-4-consal/proto"
	"blatt-4-consal/services/showservice/microservice"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("show"),
		micro.Version("latest"),
		//micro.Address(fmt.Sprintf(":%v", 3003)),
		micro.Context(context.TODO()),
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterShowHandler(service.Server(), new(microservice.Show))
	fmt.Println("Show Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
