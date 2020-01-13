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
	Id           int32
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

func (rv *Reservation) ReservationInquiry(context context.Context, req *proto.ReservationRequest, res *proto.Response) error {
	if rv.reservations == nil {
		rv.Id = 1
		rv.reservations = make(map[int32]*ReservationRequest)
	}
	if !showExists(req.Show) {
		return makeResponse(res, fmt.Sprintf("#RESERVATION_INQUIRY: Show %d does not exist yet.", req.Show))
	}
	if !userExists(req.UserName) {
		return makeResponse(res, fmt.Sprintf("#RESERVATION_INQUIRY: User %s does not exist yet.", req.UserName))
	}

	return nil
}

func (rv *Reservation) MakeReservation(context context.Context, req *proto.ReservationRequest, res *proto.Response) error {
	return nil
}

func (rv *Reservation) DeleteReservation(ctx context.Context, req *proto.ReservationRequest, res *proto.Response) error {
	return nil
}

func (rv *Reservation) GetReservations(ctx context.Context, req *proto.Request, res *proto.ReservationResponse) error {
	return nil
}

func userExists(userName string) bool {
	var client client.Client
	tmpShow := proto.NewUserService("user", client)
	res, err := tmpShow.GetUsers(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, user := range res.Value {
		if user.Name == userName {
			return true
		}
	}
	return false
}

func showExists(showId int32) bool {
	var client client.Client
	tmpShow := proto.NewShowService("show", client)
	res, err := tmpShow.GetShows(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, show := range res.Value {
		if show.Id == showId {
			return true
		}
	}
	return false
}

func makeResponse(res *proto.Response, message string) error {
	res.Success = true
	res.Message = message
	return nil
}

// Start Service for reservation class
func StartReservationService(context context.Context, port int64) {
	// Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("reservation"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(context),
	)
	// Init will parse the command line flags
	service.Init()
	// Register handler
	proto.RegisterReservationHandler(service.Server(), new(Reservation))
	fmt.Println("Reservation Service starting...")
	// Run the Server
	if err := service.Run(); err != nil {
		// Print error message if there is any
		fmt.Println(err)
	}
}
