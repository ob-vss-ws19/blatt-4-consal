package microservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
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
		return makeFailedResponse(res, fmt.Sprintf("#CHECK_RESERV_FAIL: Show %d does not exist yet.", req.Show))
	}
	if !userExists(req.UserName) {
		return makeFailedResponse(res, fmt.Sprintf("#CHECK_RESERV_FAIL: User %s does not exist yet.", req.UserName))
	}
	// is there enough seats
	availableSeats := availableSeats(req.Show, rv.reservations)
	if availableSeats < req.Seats {
		return makeFailedResponse(res, fmt.Sprintf("#CHECK_RESERV_FAIL: Only %d seats available.", availableSeats))
	}

	rv.reservations[rv.Id] = &ReservationRequest{seats: req.Seats, user: req.UserName, show: req.Show, reserved: false}
	makeResponse(res, fmt.Sprintf("#CHECK_RESERV: %d seats available for the show %d. Your reservation ID is %d.", req.Seats, req.Show, rv.Id))
	rv.Id++
	return nil
}

func (rv *Reservation) MakeReservation(context context.Context, req *proto.ReservationRequest, res *proto.Response) error {
	// synchronize the method.
	rv.mux.Lock()
	rv.mux.Unlock()

	if _, exists := rv.reservations[req.ReservationId]; !exists {
		return makeFailedResponse(res, fmt.Sprintf("#MAKE_RESERV_FAIL: There is no Reservation with ID: %d", req.ReservationId))
	}
	if rv.reservations[req.ReservationId].reserved {
		return makeFailedResponse(res, fmt.Sprintf("#MAKE_RESERV_FAIL: The Reservation with ID: %d is already booked.", req.ReservationId))
	}
	availableSeats := availableSeats(rv.reservations[req.ReservationId].show, rv.reservations)
	if availableSeats < rv.reservations[req.ReservationId].seats {
		return makeFailedResponse(res, fmt.Sprintf("#MAKE_RESERV_FAIL: There are not enough available seats for reservation with the ID %d. Available are: %d. You want to reserve: %d", req.ReservationId, availableSeats, rv.reservations[req.ReservationId].seats))
	}
	rv.reservations[req.ReservationId].reserved = true
	return makeResponse(res, fmt.Sprintf("#MAKE_RESERV: Reservation with ID %d succeed for the show %d with %d seats.", req.ReservationId, rv.reservations[req.ReservationId].show, rv.reservations[req.ReservationId].seats))
}

func (rv *Reservation) DeleteReservation(ctx context.Context, req *proto.ReservationRequest, res *proto.Response) error {
	if _, exists := rv.reservations[req.ReservationId]; !exists {
		return makeFailedResponse(res, fmt.Sprintf("#DELETE_RESERV_FAIL: Reservation with the ID %d doens't exist.", req.ReservationId))
	}
	delete(rv.reservations, req.ReservationId)
	return makeResponse(res, fmt.Sprintf("#DELETE_RESERV: Reservation with the ID %d has been deleted.", req.ReservationId))
}

func (rv *Reservation) GetReservations(ctx context.Context, req *proto.Request, res *proto.ReservationResponse) error {
	for id, v := range rv.reservations {
		res.Value = append(res.Value, &proto.ReservationRequest{ReservationId: id, UserName: v.user, Seats: v.seats, Show: v.show, Reserved: v.reserved})
	}
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

func availableSeats(showId int32, reservations map[int32]*ReservationRequest) int32 {
	var numSeats int32 = -1
	var client client.Client

	// get all shows
	tmpShow := proto.NewShowService("show", client)
	res1, err := tmpShow.GetShows(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Println(err)
		return -1
	}

	// get all cinema
	tmpCinemahall := proto.NewCinemahallService("cinemahall", client)
	res2, err := tmpCinemahall.GetCinemahalls(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Println(err)
		return -1
	}

	var cinemahallName = ""

	// find cinemahall of show
	for _, show := range res1.Value {
		if show.Id == showId {
			cinemahallName = show.CinemaHall
		}
	}

	// calculate number of all seats
	for _, cinemahall := range res2.Value {
		if cinemahall.Name == cinemahallName {
			numSeats = cinemahall.SeatRowCapacity * cinemahall.SeatRows
			break
		}
	}

	// subtract number of available seats
	for _, reservation := range reservations {
		if reservation.show == showId {
			if reservation.reserved {
				numSeats = numSeats - reservation.seats
			}
		}
	}
	return numSeats
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
