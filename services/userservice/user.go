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
	return nil
}

func (us *User) DeleteUser(ctx context.Context, req *proto.UserRequest, rsp *proto.Response) error {
	return nil
}

func (us *User) GetUsers(ctx context.Context, req *proto.Request, rsp *proto.UserResponse) error {
	return nil
}

//Start Service for user class
func StartMovieService() {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("user"),
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
