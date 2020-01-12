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

var movies = make(map[string]bool)

func (mv *Movie) AddMovie(context context.Context, req *proto.MovieRequest, res *proto.Response) error {

	if _, exists := movies[req.MovieTitle]; exists {
		res.Success = true
		res.Message = fmt.Sprintf("#ADD_MOVIE_FAIL: Movie %s does exist", req.MovieTitle)
		return nil
	}
	movies[req.MovieTitle] = true // value type is not specified
	res.Success = true
	res.Message = fmt.Sprintf("#ADD_MOVIE: Movie %s added", req.MovieTitle)
	return nil
}

func (mv *Movie) DeleteMovie(ctx context.Context, req *proto.MovieRequest, res *proto.Response) error {
	if _, exists := movies[req.MovieTitle]; !exists {
		res.Success = true
		res.Message = fmt.Sprintf("#DELETE_MOVIE_FAIL: Movie %s doesn't exist yet", req.MovieTitle)
		return nil
	}
	//create Show service client and delete corresponding shows
	deleteCorrespondingShows(req.MovieTitle)
	delete(movies, req.MovieTitle)
	res.Success = true
	res.Message = fmt.Sprintf("#DELETE_MOVIE: User %s deleted successfully", req.MovieTitle)
	return nil
}

func deleteCorrespondingShows(movieTitle string) {
	var client client.Client
	showService := proto.NewShowService("show", client)

	res, err := showService.GetShows(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("#DELETE_MOVIE_ERROR: %s", err)
	}
	// Iterate through DATA struct (shows) and call delete show
	for _, show := range res.Value {
		if movieTitle == show.Movie {
			_, err := showService.DeleteShow(context.TODO(), &proto.ShowRequest{Id: show.Id})
			if err != nil {
				fmt.Printf("#DELETE_USER_ERROR: %s", err)
			}
		}
	}
}

func (mv *Movie) GetMovies(context context.Context, req *proto.Request, res *proto.MovieResponse) error {
	for movie := range movies {
		//only key used. Value remains unused
		res.Value = append(res.Value, &proto.MovieRequest{MovieTitle: movie})
	}
	return nil
}

//Start Service for movie class
func StartMovieService(context context.Context, port int64) {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("movie"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(context), //needed
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
