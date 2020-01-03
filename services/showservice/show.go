package showservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type Show struct {
	Id int32
}

type ShowRequest struct {
	Movie      string
	Cinemahall string
}

var shows = make(map[int32]*ShowRequest)

func (sw *Show) AddShow(ctx context.Context, req *proto.ShowRequest, rsp *proto.Response) error {
	return nil
}

func (sw *Show) DeleteShow(ctx context.Context, req *proto.ShowRequest, rsp *proto.Response) error {
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
	service := micro.NewService(
		micro.Name("show"),
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterShowHandler(service.Server(), new(Show))
	fmt.Println("Reservation Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
