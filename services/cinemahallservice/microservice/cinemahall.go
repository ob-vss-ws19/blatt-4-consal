package microservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

type CinemahallRequest struct {
	CinemaName       string
	SeatRows         int32
	SeatRowsCapacity int32
}

type Cinemahall struct {
}

//initialize a map using built in function make
var cinemas = make(map[string]*CinemahallRequest)

//functions for cinema class
func (cm *Cinemahall) AddCinemahall(ctx context.Context, req *proto.CinemahallRequest, res *proto.Response) error {
	// A two-value assignment tests for the existence of a key
	//if ok true -> key exists in the map
	if _, exists := cinemas[req.Name]; exists {
		return makeFailedResponse(res, fmt.Sprintf("#ADD_CINE_FAIL: Cinemahall %s does already exist", req.Name))
	}
	//Cinema doesn't exist. Add new one
	cinemas[req.Name] = &CinemahallRequest{SeatRows: req.SeatRows, SeatRowsCapacity: req.SeatRowCapacity}
	return makeResponse(res, fmt.Sprintf("#ADD_CINE: Movie %s added", req.Name))
}

func (cm *Cinemahall) DeleteCinemahall(ctx context.Context, req *proto.CinemahallRequest, res *proto.Response) error {
	if _, exists := cinemas[req.Name]; !exists {
		return makeFailedResponse(res, fmt.Sprintf("#DELETE_CINE_FAIL: Cinema %s doesn't exist yet", req.Name))
	}
	//Cinema does exist
	//create Show service client and delete corresponding shows
	//delete cinemahall from map

	deleteCorrespondingShows(req.Name)
	delete(cinemas, req.Name)
	return makeResponse(res, fmt.Sprintf("#DELETE_MOVIE: User %s deleted successfully", req.Name))
}

func deleteCorrespondingShows(cinemahallName string) {
	var client client.Client
	showService := proto.NewShowService("show", client)

	res, err := showService.GetShows(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("#DELETE_CINE_ERROR: %s", err)
	}
	//Iterate through DATA struc (shows) and call delete show
	for _, show := range res.Value {
		if cinemahallName == show.CinemaHall {
			_, err := showService.DeleteShow(context.TODO(), &proto.ShowRequest{Id: show.Id})
			if err != nil {
				fmt.Printf("#DELETE_CINE_ERROR: %s", err)
			}
		}
	}
}

func (cm *Cinemahall) GetCinemahalls(ctx context.Context, req *proto.Request, rsp *proto.CinemahallResponse) error {
	for k, v := range cinemas {
		rsp.Value = append(rsp.Value, &proto.CinemahallRequest{Name: k, SeatRowCapacity: v.SeatRowsCapacity, SeatRows: v.SeatRows})
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

//Start Service for cinema class
func StartCinemaService(context context.Context, port int64) {
	//Create a new Service. Include name, version, address and context
	service := micro.NewService(
		micro.Name("cinemahall"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(context), //needed
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterCinemahallHandler(service.Server(), new(Cinemahall))
	fmt.Println("Cinemahall Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
