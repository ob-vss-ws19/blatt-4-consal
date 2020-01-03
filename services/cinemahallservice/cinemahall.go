package cinemahallservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

type CinemaRequest struct {
	CinemaName       string
	SeatRows         int32
	SeatRowsCapacity int32
}

type Cinema struct {
}

//initialize a map using built in function make
var cinemas = make(map[string]*CinemaRequest)

//functions for cinema class
func (cm *Cinema) AddCinema(ctx context.Context, req *proto.CinemaRequest, rsp *proto.Response) error {
	//A two-value assignment tests for the existence of a key
	//if ok true -> key exists in the map
	if _, ok := cinemas[req.Name]; ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Cinema %s does already exist", req.Name)
		return nil
	}
	//Cinema doesn't exist. Add new one
	cinemas[req.Name] = &CinemaRequest{SeatRows: req.SeatRows, SeatRowsCapacity: req.SeatRowCapacity}
	rsp.Success = true
	rsp.Message = fmt.Sprintf("New Cinema %s added", req.Name)
	return nil
}

func deleteCorrespondingShows(cinemahall string) {
	var client client.Client
	show := proto.NewShowService("showing", client)

	rsp, err := show.GetShows(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	//Iterate through DATA struc (shows) and call delete show
	for _, v := range rsp.Value {
		if cinemahall == v.CinemaHall {
			_, err := show.DeleteShow(context.TODO(), &proto.ShowRequest{Id: v.Id})
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}
	}
}

func (cm *Cinema) DeleteCinema(ctx context.Context, req *proto.CinemaRequest, rsp *proto.Response) error {
	if _, ok := cinemas[req.Name]; !ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Cinema %s does not exist", req.Name)
		return nil
	}
	//Cinema does exist
	//create Show service client and delete corresponding shows
	deleteCorrespondingShows(req.Name)
	//delete cinemahall from map
	delete(cinemas, req.Name)
	rsp.Success = true
	rsp.Message = fmt.Sprint("Cinema %s was deleted", req.Name)
	return nil
}

func (cm *Cinema) GetCinemas(ctx context.Context, req *proto.Request, rsp *proto.CinemaResponse) error {
	for k, v := range cinemas {
		rsp.Value = append(rsp.Value, &proto.CinemaRequest{Name: k, SeatRowCapacity: v.SeatRowsCapacity, SeatRows: v.SeatRows})
	}
	return nil
}

//Start Service for movie class
func StartCinemaService() {
	//Create a new Service. Include name, version, address and context
	var port int32 = 8082
	service := micro.NewService(
		micro.Name("cinema"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(nil), //needed
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
