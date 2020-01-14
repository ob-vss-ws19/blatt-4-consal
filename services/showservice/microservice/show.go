package microservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

type Show struct {
}

type ShowRequest struct {
	Movie      string
	Cinemahall string
}

var Shows = make(map[int32]*ShowRequest)
var Id int32 = 1

func (*Show) AddShow(context context.Context, req *proto.ShowRequest, res *proto.Response) error {
	if !doesMovieExist(req.Movie) {
		return makeFailedResponse(res, fmt.Sprintf("#SHOW_ADD_FAIL: Movie %s doesn't exist yet", req.Movie))
	}
	if !doesCinemahallExist(req.CinemaHall) {
		return makeFailedResponse(res, fmt.Sprintf("#SHOW_ADD_FAIL: Cinemahall %s doesn't exist yet", req.Movie))
	}
	Shows[Id] = &ShowRequest{Movie: req.Movie, Cinemahall: req.CinemaHall}
	makeResponse(res, fmt.Sprintf("#SHOW_ADD: New Show with ID %d in Cinema %s with Movie Title %s added", Id, req.CinemaHall, req.Movie))
	Id++
	return nil
}

func (*Show) DeleteShow(context context.Context, req *proto.ShowRequest, res *proto.Response) error {
	if _, exists := Shows[req.Id]; !exists {
		return makeFailedResponse(res, fmt.Sprintf("#SHOW_DELETE_FAIL: Show %d does not exist", req.Id))
	}
	// create Reservation service client and delete corresponding Reservations
	// delete cinemahall from map
	// deleteCorrespondingReservations(req.Id)
	delete(Shows, req.Id)
	return makeResponse(res, fmt.Sprintf("#DELETE_MOVIE: Show with id %d deleted successfully", req.Id))
}

func deleteCorrespondingReservations(showId int32) {
	var client client.Client
	reservationService := proto.NewReservationService("reservation", client)

	rsp, err := reservationService.GetReservations(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	//Iterate through DATA struc (reservations) and call delete reservation
	for _, v := range rsp.Value {
		if showId == v.Show {
			_, err := reservationService.DeleteReservation(context.TODO(), &proto.ReservationRequest{ReservationId: v.ReservationId})
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}
	}
}

func (*Show) GetShows(context context.Context, req *proto.Request, rsp *proto.ShowResponse) error {
	for id, v := range Shows {
		rsp.Value = append(rsp.Value, &proto.ShowRequest{Id: id, CinemaHall: v.Cinemahall, Movie: v.Movie})
	}
	return nil
}

func doesCinemahallExist(cinemahall string) bool {
	var client client.Client
	cinemahallService := proto.NewCinemahallService("cinemahall", client)
	res, err := cinemahallService.GetCinemahalls(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("#SHOW_FIND_CINE: Error: %s", err)
		return false
	}
	for _, cinema := range res.Value {
		if cinemahall == cinema.Name {
			//Cinemahall does exist
			return true
		}
	}
	return false
}

func doesMovieExist(movieTitle string) bool {
	var client client.Client
	movieService := proto.NewMovieService("movie", client)
	res, err := movieService.GetMovies(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return false
	}
	for _, movie := range res.Value {
		if movieTitle == movie.MovieTitle {
			// Movie does exists
			return true
		}
	}
	return false
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

//Start Service for Show class
func StartShowService(context context.Context) {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("show"),
		micro.Version("latest"),
		micro.Context(context),
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterShowHandler(service.Server(), new(Show))
	fmt.Println("Show Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
