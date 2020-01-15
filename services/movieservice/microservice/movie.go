package microservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
)

type Movie struct {
}

var movies = make(map[string]bool)

func (mv *Movie) AddMovie(context context.Context, req *proto.MovieRequest, res *proto.Response) error {
	if _, exists := movies[req.MovieTitle]; exists {
		return makeFailedResponse(res, fmt.Sprintf("#ADD_MOVIE_FAIL: Movie %s does exist", req.MovieTitle))
	}
	movies[req.MovieTitle] = true // value type is not specified
	fmt.Println(req.MovieTitle)
	return makeResponse(res, fmt.Sprintf("#ADD_MOVIE: Movie %s added", req.MovieTitle))
}

func (mv *Movie) DeleteMovie(ctx context.Context, req *proto.MovieRequest, res *proto.Response) error {
	if _, exists := movies[req.MovieTitle]; !exists {
		return makeFailedResponse(res, fmt.Sprintf("#DELETE_MOVIE_FAIL: Movie %s doesn't exist yet", req.MovieTitle))
	}
	//create Show service client and delete corresponding shows
	deleteCorrespondingShows(req.MovieTitle)
	delete(movies, req.MovieTitle)
	return makeResponse(res, fmt.Sprintf("#DELETE_MOVIE: User %s deleted successfully", req.MovieTitle))
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

func makeResponse(res *proto.Response, message string) error {
	res.Success = true
	res.Message = message
	return nil
}

func makeFailedResponse(res *proto.Response, message string) error {
	res.Success = false
	res.Message = message
	return nil
}
