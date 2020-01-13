package main

import (
	"blatt-4-consal/proto"
	"blatt-4-consal/services/userservice/microservice"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		//micro.Address(fmt.Sprintf(":%v", 3000)),
		micro.Context(context.TODO()),
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterUserHandler(service.Server(), new(microservice.User))
	fmt.Println("User Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
