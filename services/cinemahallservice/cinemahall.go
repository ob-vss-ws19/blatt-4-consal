package cinemahallservice

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
func (cm *Cinemahall) AddCinemahall(ctx context.Context, req *proto.CinemahallRequest, rsp *proto.Response) error {
	//A two-value assignment tests for the existence of a key
	//if ok true -> key exists in the map
	if _, ok := cinemas[req.Name]; ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Cinemahall %s does already exist", req.Name)
		return nil
	}
	//Cinema doesn't exist. Add new one
	cinemas[req.Name] = &CinemahallRequest{SeatRows: req.SeatRows, SeatRowsCapacity: req.SeatRowCapacity}
	rsp.Success = true
	rsp.Message = fmt.Sprintf("New Cinemahall %s added", req.Name)
	return nil
}

func deleteCorrespondingShows(cinemahallName string) {
	var client client.Client
	showService := proto.NewShowService("show", client)

	rsp, err := showService.GetShows(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	//Iterate through DATA struc (shows) and call delete show
	for _, v := range rsp.Value {
		if cinemahallName == v.CinemaHall {
			_, err := showService.DeleteShow(context.TODO(), &proto.ShowRequest{Id: v.Id})
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}
	}
}

func (cm *Cinemahall) DeleteCinemahall(ctx context.Context, req *proto.CinemahallRequest, rsp *proto.Response) error {
	if _, ok := cinemas[req.Name]; !ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Cinemahall %s does not exist", req.Name)
		return nil
	}
	//Cinema does exist
	//create Show service client and delete corresponding shows
	deleteCorrespondingShows(req.Name)
	//delete cinemahall from map
	delete(cinemas, req.Name)
	rsp.Success = true
	rsp.Message = fmt.Sprint("Cinemahall %s was deleted", req.Name)
	return nil
}

func (cm *Cinemahall) GetCinemahalls(ctx context.Context, req *proto.Request, rsp *proto.CinemahallResponse) error {
	for k, v := range cinemas {
		rsp.Value = append(rsp.Value, &proto.CinemahallRequest{Name: k, SeatRowCapacity: v.SeatRowsCapacity, SeatRows: v.SeatRows})
	}
	return nil
}

//Start Service for movie class
func StartCinemaService() {
	//Create a new Service. Include name, version, address and context
	var port int32 = 8081
	service := micro.NewService(
		micro.Name("cinemahall"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(nil), //needed
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
