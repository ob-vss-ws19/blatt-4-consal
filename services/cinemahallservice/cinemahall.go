package cinemahallservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type Cinema struct {
	CinemaName       string
	SeatRows         int32
	SeatRowsCapacity int32
}

type CinemaRequest struct {
	//TODO: add map
}

//functions for cinema class
func (cm *Cinema) AddCinema(ctx context.Context, req *proto.CinemaRequest, rsp *proto.Response) error {
	rsp.Greeting = "Hello" + req.Name

	return nil
}

func (cm *Cinema) DeleteCinema(ctx context.Context, req *proto.CinemaRequest, rsp *proto.Response) error {
	rsp.Greeting = "Hello" + req.Name
	return nil
}

func (cm *Cinema) GetCinemas(ctx context.Context, req *proto.Request, rsp *proto.CinemaResponse) error {
	return nil
}

//Start Service for movie class
func StartCinemaService() {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("cinema"),
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterCinemaHandler(service.Server(), new(Cinema))
	fmt.Println("Cinema Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
