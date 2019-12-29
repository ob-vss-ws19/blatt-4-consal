package movieservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type Movie struct {
	movies map[string]bool
}

func (mv *Movie) AddMovie(ctx context.Context, req *proto.MovieRequest, rsp *proto.Response) error {
	return nil
}

func (mv *Movie) DeleteMovie(ctx context.Context, req *proto.MovieRequest, rsp *proto.Response) error {
	return nil
}

func (mv *Movie) GetMovies(ctx context.Context, req *proto.Request, rsp *proto.MovieResponse) error {
	return nil
}

//Start Service for movie class
func StartMovieService() {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("movie"),
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
