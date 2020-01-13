package main

import (
	"blatt-4-consal/proto"
	"blatt-4-consal/services/movieservice/microservice"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("movie"),
		micro.Version("latest"),
		//micro.Address(fmt.Sprintf(":%v", 3001)),
		micro.Context(context.TODO()), //needed
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterMovieHandler(service.Server(), new(microservice.Movie))
	fmt.Println("Movie Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
