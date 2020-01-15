package microservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

var cli client.Client
var user proto.UserService

func init() {
	go StartService("user")
}

func StartService(microservicename string) {
	context, _ := context.WithCancel(context.Background())
	service := micro.NewService(
		micro.Name(microservicename),
		micro.Context(context),
	)
	proto.RegisterUserHandler(service.Server(), new(User))
	fmt.Printf("Starting %sservice", microservicename)
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func TestUser_AddUser(t *testing.T) {
	user = proto.NewUserService("user", cli)
	req1 := &proto.UserRequest{
		Name: "user1",
	}
	user.GetUsers(context.TODO(), &proto.Request{})
	res1, err := user.AddUser(context.TODO(), req1)
	assert.Nil(t, err)
	assert.True(t, res1.Success)
}
