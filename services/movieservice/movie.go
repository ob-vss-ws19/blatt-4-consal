package movieservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type Movie struct {}

func (mv *Movie) AddMovie(ctx context.Context, req *proto.MovieRequest, rsp *proto.MovieResponse) error {
	rsp.Greeting = "Hello" + req.Name
	return nil
}

func (mv *Movie) DeleteMovie(ctx context.Context, req *proto.MovieRequest, rsp *proto.MovieResponse) error {
	rsp.Greeting = "Hello" + req.Name
	return nil
}


//Start Service for movie class
func StartMoviesService() {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("cinema"),
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterMovieHandler(service.Server(), new(Movie))
	fmt.Println("Movie Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
