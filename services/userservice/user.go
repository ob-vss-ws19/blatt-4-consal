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
		rsp.Message = fmt.Sprintf("# User '#{req.Name}' does exist already.")
	}
	// Setze neuen User in die Map
	us.Users[req.Name] = true
	rsp.Success = true
	rsp.Message = fmt.Sprintf("# Created new User '#{req.Name}'.")
	return nil
}

func deleteCorrespondingReservations(userName string) {
	var client client.Client
	reservationService := proto.NewReservationService("reservation", client)

	rsp, err := reservationService.GetReservations(context.TODO(), &proto.Request{})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	//Iterate through DATA struc (Reservations) and call delete reservation
	for _, v := range rsp.Value {
		if userName == v.UserName {
			_, err := reservationService.DeleteReservation(context.TODO(), &proto.ReservationRequest{ReservationId: v.ReservationId})
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}
	}
}

func (us *User) DeleteUser(ctx context.Context, req *proto.UserRequest, rsp *proto.Response) error {
	if _, ok := us.Users[req.Name]; !ok {
		rsp.Success = false
		rsp.Message = fmt.Sprintf("User %s doesn't exist", req.Name)
		return nil
	}
	deleteCorrespondingReservations(req.Name)
	delete(us.Users, req.Name)
	rsp.Success = true
	rsp.Message = fmt.Sprintf("User %s was deleted", req.Name)
	return nil
}

func (us *User) GetUsers(ctx context.Context, req *proto.Request, rsp *proto.UserResponse) error {
	for k := range us.Users {
		//only key used. Value remains unused
		rsp.Value = append(rsp.Value, &proto.UserRequest{Name: k})
	}
	return nil
}

//Start Service for user class
func StartUserService() {
	//Create a new Service. Add name address and context
	var port int32 = 8084
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(nil),
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
