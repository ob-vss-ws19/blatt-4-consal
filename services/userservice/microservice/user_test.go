package microservice

import (
	"blatt-4-consal/proto"
	microservice2 "blatt-4-consal/services/cinemahallservice/microservice"
	"blatt-4-consal/services/movieservice/microservice"
	microservice4 "blatt-4-consal/services/reservationservice/microservice"
	microservice3 "blatt-4-consal/services/showservice/microservice"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

var cli client.Client
var user proto.UserService
var mux sync.RWMutex

func init() {
	go StartService("user", 3000)
	time.Sleep(1 * time.Second)
	go StartService("movie", 3001)
	time.Sleep(1 * time.Second)

	go StartService("cinemahall", 3002)
	//go StartService("show")
	//go StartService("reservation")
	time.Sleep(2 * time.Second)
}

func StartService(microservicename string, port int32) {
	mux.Lock()
	context, _ := context.WithCancel(context.Background())
	service := micro.NewService(
		micro.Name(microservicename),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(context),
	)

	switch microservicename {
	case "user":
		proto.RegisterUserHandler(service.Server(), new(User))
		printInfo(microservicename)
	case "movie":
		proto.RegisterMovieHandler(service.Server(), new(microservice.Movie))
		printInfo(microservicename)
	case "cinemahall":
		proto.RegisterCinemahallHandler(service.Server(), new(microservice2.Cinemahall))
		printInfo(microservicename)
	case "show":
		proto.RegisterShowHandler(service.Server(), new(microservice3.Show))
		printInfo(microservicename)
	case "reservation":
		proto.RegisterReservationHandler(service.Server(), new(microservice4.Reservation))
		printInfo(microservicename)
	}
	mux.Unlock()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func printInfo(microservicename string) {
	fmt.Printf("Starting %sservice\n", microservicename)
}

func TestUser_AddUser(t *testing.T) {
	user = proto.NewUserService("user", cli)
	req1 := &proto.UserRequest{Name: "user1"}

	fix()
	res1, err := user.AddUser(context.TODO(), req1)

	assert.Nil(t, err)
	assert.True(t, res1.Success)
}

func TestUser_AddDoubleUser(t *testing.T) {
	user = proto.NewUserService("user", cli)
	req1 := &proto.UserRequest{Name: "user2"}

	fix()
	res1, err1 := user.AddUser(context.TODO(), req1)
	res2, err2 := user.AddUser(context.TODO(), req1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.False(t, res2.Success)
}

func TestUser_AddTripleUser(t *testing.T) {
	user = proto.NewUserService("user", cli)
	req1 := &proto.UserRequest{Name: "user3"}

	fix()
	res1, err1 := user.AddUser(context.TODO(), req1)
	res2, err2 := user.AddUser(context.TODO(), req1)
	res3, err3 := user.AddUser(context.TODO(), req1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.False(t, res2.Success)

	assert.Nil(t, err3)
	assert.False(t, res3.Success)
}

func TestUser_DeleteUser(t *testing.T) {

}

func fix() {
	user.GetUsers(context.TODO(), &proto.Request{})
}
