package showservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

type Show struct {
	Id int32
}

type ShowRequest struct {
	Movie      string
	Cinemahall string
}

var shows = make(map[int32]*ShowRequest)

func doesMovieExist(movieTitle string) bool {
	var client client.Client
	movieService := proto.NewMovieService("movie", client)

	rsp, err := movieService.GetMovies(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return false
	}
	//user does exist
	for _, v := range rsp.Value {
		if movieTitle == v.MovieTitle {
			return true
		}
	}
	return false
}

func doesCinemahallExist(cinemahall string) bool {
	var client client.Client
	cinemahallService := proto.NewCinemahallService("cinemahall", client)

	rsp, err := cinemahallService.GetCinemahalls(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return false
	}
	//Cinemahall does exist
	for _, v := range rsp.Value {
		if cinemahall == v.Name {
			return true
		}
	}
	return false
}

func (sw *Show) AddShow(ctx context.Context, req *proto.ShowRequest, rsp *proto.Response) error {
	if !doesMovieExist(req.Movie) {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Movie %s does not exist", req.Movie)
		return nil
	}
	if !doesCinemahallExist(req.CinemaHall) {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Cinemahall %s does not exist", req.CinemaHall)
		return nil
	}

	shows[sw.Id] = &ShowRequest{Movie: req.Movie, Cinemahall: req.CinemaHall}
	rsp.Success = true
	rsp.Message = fmt.Sprintf("New Show with ID %s in Cinema %s with Movie Title %s added", req.Id, req.CinemaHall, req.Movie)
	sw.Id++
	return nil
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

func (sw *Show) DeleteShow(ctx context.Context, req *proto.ShowRequest, rsp *proto.Response) error {
	if _, ok := shows[req.Id]; !ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Show %s does not exist", req.Id)
		return nil
	}
	//create Reservation service client and delete corresponding Reservations
	deleteCorrespondingReservations(req.Id)
	//delete cinemahall from map
	delete(shows, req.Id)
	rsp.Success = true
	rsp.Message = fmt.Sprintf("Show with id: %s was deleted", req.Id)
	return nil
}

func (sw *Show) GetShows(ctx context.Context, req *proto.Request, rsp *proto.ShowResponse) error {
	for k, v := range shows {
		rsp.Value = append(rsp.Value, &proto.ShowRequest{Id: k, CinemaHall: v.Cinemahall, Movie: v.Movie})
	}
	return nil
}

//Start Service for Show class
func StartReservationService() {
	//Create a new Service. Add name address and context
	var port int32 = 8084
	service := micro.NewService(
		micro.Name("show"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(nil),
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
