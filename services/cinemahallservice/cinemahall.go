package cinemahallservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type CinemaRequest struct {
	CinemaName       string
	SeatRows         int32
	SeatRowsCapacity int32
}

type Cinema struct {
	//Cinemas map[string]*CinemaRequest
}

//initialize a map using built in function make
var cinemas = make(map[string]*CinemaRequest)

//functions for cinema class
func (cm *Cinema) AddCinema(ctx context.Context, req *proto.CinemaRequest, rsp *proto.Response) error {
	//A two-value assignment tests for the existence of a key
	//if ok true -> key exists in the map
	if _, ok := cinemas[req.Name]; ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Cinema %s already exists", req.Name)
		return nil
	}
	//Cinema doesn't exist. Add new one
	cinemas[req.Name] = &CinemaRequest{SeatRows: req.SeatRows, SeatRowsCapacity: req.SeatRowCapacity}
	rsp.Success = true
	rsp.Message = fmt.Sprintf("New Cinema %s added", req.Name)
	return nil
}

func (cm *Cinema) DeleteCinema(ctx context.Context, req *proto.CinemaRequest, rsp *proto.Response) error {
	if _, ok := cinemas[req.Name]; !ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Cinema %s does not exist", req.Name)
		return nil
	}
	//Cinema does exist
	//Delete shows for cinema aswell?
	//TODO: conjuction between cinema and show
	// Setup and the client
	/*	func runClient(service micro.Service) {
		// Create new greeter client
		greeter := proto.NewGreeterService("greeter", service.Client())

		// Call the greeter
		rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "John"})
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print response
		fmt.Println(rsp.Greeting)
	}*/
	//create Show service client to get shows available and delete the ones
	//for this cinema
/*	var client client.Client
	show := proto.NewShowService("showing", client)

	rsp, err := show.GetShows()
		if err == nil {

		}*/
	//TODO: handle nil. Iterate thorugh DATA struc and call delete showing

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
	//Create a new Service. Add name address and context
	var port int32 = 8081
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
