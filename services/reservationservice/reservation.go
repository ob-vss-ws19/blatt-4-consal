package reservationservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"sync"
)

type Reservation struct {
	reservations map[int32]*ReservationRequest
	nextID       int32
	mux          sync.RWMutex
}

type ReservationRequest struct {
	user          string
	reservationId int32
	show          int32
	seats         int32
	reserved      bool
}

//initialize a map using built in function make

var reservations = make(map[int32]*ReservationRequest)
var reservationNumber = 1

func doesUserExists(userName string) bool {
	var client client.Client
	userService := proto.NewUserService("user", client)
	rsp, err := userService.GetUsers(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return false
	}
	//user does exist
	for _, v := range rsp.Value {
		if userName == v.Name {
			return true
		}
	}
	return false
}

func doesShowExist(showId int32) bool {
	var client client.Client
	showService := proto.NewShowService("show", client)
	rsp, err := showService.GetShows(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return false
	}
	//user does exist
	for _, v := range rsp.Value {
		if showId == v.Id {
			return true
		}
	}
	return false
}

func (rv *Reservation) MakeReservation(ctx context.Context, req *proto.ReservationRequest, rsp *proto.Response) error {
	//check if user already exists or made a reservation
	if !doesUserExists(req.UserName) {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("User %s does not exist", req.UserName)
		return nil
	}
	//check if show does exist
	if !doesShowExist(req.Show) {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Show %s does not exist", req.Show)
	}

	//check if there are free seats

	//Add new reservation
	if _, ok := reservations[req.ReservationId]; ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("Reservation with id: %s does already exist", req.ReservationId)
		return nil
	}

	reservations[req.ReservationId] = &ReservationRequest{
		user: req.UserName, reservationId: 123} //TODO
	rsp.Success = true
	rsp.Message = fmt.Sprintf("New Reservation with id: %s added", req.ReservationId)
	return nil
}

func (rv *Reservation) DeleteReservation(ctx context.Context, req *proto.ReservationRequest, rsp *proto.Response) error {
	return nil
}

func (rv *Reservation) ReservationInquiry(ctx context.Context, req *proto.ReservationRequest, rsp *proto.Response) error {
	return nil
}

func (rv *Reservation) GetReservations(ctx context.Context, req *proto.Request, rsp *proto.ReservationResponse) error {
	return nil
}

//Start Service for movie class
func StartReservationService(context context.Context, port int64) {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("reservation"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(context), //needed
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterReservationHandler(service.Server(), new(Reservation))
	fmt.Println("Reservation Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
