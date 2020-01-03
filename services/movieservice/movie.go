package movieservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

type Movie struct {
}

//initialize a map using built in function make
var movies = make(map[string]bool)

func (mv *Movie) AddMovie(ctx context.Context, req *proto.MovieRequest, rsp *proto.Response) error {
	if _, ok := movies[req.MovieTitle]; ok {
		rsp.Success = false;
		rsp.Message = fmt.Sprintf("Movie %s does already exist", req.MovieTitle)
		return nil
	}

	movies[req.MovieTitle] = true //value type is not specified
	rsp.Success = true
	rsp.Message = fmt.Sprintf("New Movie %s added", req.MovieTitle)
	return nil
}

func deleteCorrespondingShows(movieTitle string) {
	var client client.Client
	show := proto.NewShowService("showing", client)

	rsp, err := show.GetShows(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	//Iterate through DATA struc (shows) and call delete show
	for _, v := range rsp.Value {
		if movieTitle == v.Movie {
			_, err := show.DeleteShow(context.TODO(), &proto.ShowRequest{Id: v.Id})
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}
	}
}
func (mv *Movie) DeleteMovie(ctx context.Context, req *proto.MovieRequest, rsp *proto.Response) error {
	if _, ok := movies[req.MovieTitle]; !ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Movie %s does not exist0", req.MovieTitle)
		return nil
	}
	//create Show service client and delete corresponding shows
	deleteCorrespondingShows(req.MovieTitle)
	delete(movies, req.MovieTitle)
	rsp.Success = true
	rsp.Message = fmt.Sprintf("Movie %s was deleted", req.MovieTitle)
	return nil
}

func (mv *Movie) GetMovies(ctx context.Context, req *proto.Request, rsp *proto.MovieResponse) error {
	for k := range movies {
		//only key used. Value remains unused
		rsp.Value = append(rsp.Value, &proto.MovieRequest{MovieTitle: k})
	}
	return nil
}

//Start Service for movie class
func StartMovieService() {
	//Create a new Service. Add name address and context
	var port int32 = 8081
	service := micro.NewService(
		micro.Name("movie"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(nil), //needed
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
