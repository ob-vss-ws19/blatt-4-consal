package microservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

type User struct {
}

var Users = make(map[string]bool)

func (us *User) AddUser(ctx context.Context, req *proto.UserRequest, res *proto.Response) error {
	// kontrollieren ob Benutzer schon existiert.
	if _, exists := Users[req.Name]; exists {
		return makeFailedResponse(res, fmt.Sprintf("#ADD_USER_FAIL: User '%s' does exist already.", req.Name))
	}
	// Setze neuen User in die Map
	Users[req.Name] = true
	return makeResponse(res, fmt.Sprintf("#ADD_USER: Created new User '%s'.", req.Name))
}

func (us *User) DeleteUser(context context.Context, req *proto.UserRequest, res *proto.Response) error {
	if _, exists := Users[req.Name]; !exists {
		return makeFailedResponse(res, fmt.Sprintf("#DELETE_USER_FAIL: User %s doesn't exist yet", req.Name))
	}
	deleteCorrespondingReservations(req.Name)
	delete(Users, req.Name)
	return makeResponse(res, fmt.Sprintf("#DELETE_USER: User %s deleted successfully", req.Name))

}

func deleteCorrespondingReservations(userName string) {
	var client client.Client
	reservationService := proto.NewReservationService("reservation", client)
	res, err := reservationService.GetReservations(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("#DELETE_USER_ERROR: %s", err)
		return
	}
	//Iterate through DATA struc (Reservations) and call delete reservation
	for _, reservation := range res.Value {
		if userName == reservation.UserName {
			_, err := reservationService.DeleteReservation(context.TODO(), &proto.ReservationRequest{ReservationId: reservation.ReservationId})
			if err != nil {
				fmt.Printf("#DELETE_USER_ERROR: %s", err)
			}
		}
	}
}

func (us *User) GetUsers(context context.Context, req *proto.Request, res *proto.UserResponse) error {
	for user := range Users {
		//only key used. Value remains unused
		res.Value = append(res.Value, &proto.UserRequest{Name: user})
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

func StartUserService(context context.Context) {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Context(context),
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterUserHandler(service.Server(), new(User))
	fmt.Println("User Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
