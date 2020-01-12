package userservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

type User struct {
	Users map[string]bool
}

func (us *User) AddUser(ctx context.Context, req *proto.UserRequest, rsp *proto.Response) error {
	// Erstelle neue Userliste falls noch keine existiert.
	if us.Users == nil {
		us.Users = make(map[string]bool)
	}
	// kontrollieren ob Benutzer schon existiert.
	if _, exists := us.Users[req.Name]; exists {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("#ADD_USER_FAIL: User '%s' does exist already.", req.Name)
		return nil
	}
	// Setze neuen User in die Map
	us.Users[req.Name] = true
	rsp.Success = true
	rsp.Message = fmt.Sprintf("#ADD_USER: Created new User '%s'.", req.Name)
	return nil
}

func (us *User) DeleteUser(context context.Context, req *proto.UserRequest, res *proto.Response) error {
	if _, exists := us.Users[req.Name]; !exists {
		res.Success = false
		res.Message = fmt.Sprintf("#DELETE_USER_FAIL: User %s doesn't exist yet", req.Name)
		return nil
	}
	deleteCorrespondingReservations(req.Name)
	delete(us.Users, req.Name)
	res.Success = true
	res.Message = fmt.Sprintf("#DELETE_USER: User %s deleted successfully", req.Name)
	return nil
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
	for user := range us.Users {
		//only key used. Value remains unused
		res.Value = append(res.Value, &proto.UserRequest{Name: user})
	}
	return nil
}

//Start Service for user class
func StartUserService(context context.Context) {
	//Create a new Service. Add name address and context
	var port int32 = 8096
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
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
