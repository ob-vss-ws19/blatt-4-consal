package reservationservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"sync"
)

type Reservation struct {
	reservations map[int32]*ReservationData
	nextID       int32
	mux          sync.RWMutex
}

type ReservationData struct {
	showing int32
	seats   int32
	booked  bool
	user    string
}

//initialize a map using built in function make
var reservations = make(map[int32]*ReservationData)

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
