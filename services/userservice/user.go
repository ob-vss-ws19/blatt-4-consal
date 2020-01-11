package userservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
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
		rsp.Message = fmt.Sprintf("User with name %s already exists.", req.Name)
	}
	// Setze neuen User in die Map
	us.Users[req.Name] = true
	rsp.Success = true
	rsp.Message = fmt.Sprintf("# Created new User #{req.Name}.")
	return nil
}

func (us *User) DeleteUser(ctx context.Context, req *proto.UserRequest, rsp *proto.Response) error {

	return nil
}

func (us *User) GetUsers(ctx context.Context, req *proto.Request, rsp *proto.UserResponse) error {
	return nil
}

// Start Service for user class
func StartUserService(context context.Context) {
	var port int64 = 8091
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("User"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(context),
	)
	// Init will parse the command line flags
	service.Init()
	// Register handler
	proto.RegisterUserHandler(service.Server(), new(User))
	fmt.Println("User Service is starting...")
	// Run the Server
	if err := service.Run(); err != nil {
		// Print error message if there is any
		fmt.Println(err)
	}
}
