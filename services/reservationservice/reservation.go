package reservationservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type Reservation struct {
	reservationNumber string // int or string?
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


func (rv *Reservation) MakeReservation(ctx context.Context, req *proto.ReservationRequest, rsp *proto.Response) error {
	return nil
}

func (rv *Reservation) DeleteReservation(ctx context.Context, req *proto.ReservationRequest, rsp *proto.Response) error {
	return nil
}

func (rv *Reservation) CheckReservation(ctx context.Context, req *proto.ReservationRequest, rsp *proto.Response) error {
	return nil
}

func (rv *Reservation) GetReservations(ctx context.Context, req *proto.Request, rsp *proto.ReservationResponse) error {
	return nil
}

//Start Service for movie class
func StartReservationService() {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("reservation"),
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

