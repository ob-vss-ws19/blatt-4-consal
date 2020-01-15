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
var movie proto.MovieService
var cinemahall proto.CinemahallService
var show proto.ShowService
var reservation proto.ReservationService
var mux sync.RWMutex

func init() {
	go StartService("user", 3000)
	time.Sleep(1 * time.Second)
	go StartService("movie", 3001)
	time.Sleep(1 * time.Second)
	go StartService("cinemahall", 3002)
	time.Sleep(1 * time.Second)
	go StartService("show", 3003)
	time.Sleep(1 * time.Second)
	go StartService("reservation", 3004)
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

func TestUser_DeleteUser_WithoutReservation(t *testing.T) {
	user = proto.NewUserService("user", cli)

	us := "user5"

	user5 := &proto.UserRequest{Name: us}

	fix()
	res1, err1 := user.AddUser(context.TODO(), user5)
	res2, err2 := user.DeleteUser(context.TODO(), user5)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)
}

func TestUser_DeleteUser_WithReservation(t *testing.T) {
	user = proto.NewUserService("user", cli)
	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)
	reservation = proto.NewReservationService("reservation", cli)

	us := "user4"
	mv := "movie4"
	cm := "cinemahall4"

	user4 := &proto.UserRequest{Name: us}
	movie4 := &proto.MovieRequest{MovieTitle: mv}
	cinemahall4 := &proto.CinemahallRequest{Name: cm, SeatRows: 10, SeatRowCapacity: 10}
	show4 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	checkReservation4 := &proto.ReservationRequest{UserName: us, Show: 1, Seats: 5}
	makeReservation4 := &proto.ReservationRequest{ReservationId: 1}
	deleteReservation4 := &proto.ReservationRequest{ReservationId: 1}

	fix()
	movie.AddMovie(context.TODO(), movie4)
	cinemahall.AddCinemahall(context.TODO(), cinemahall4)
	show.AddShow(context.TODO(), show4)
	reservation.ReservationInquiry(context.TODO(), checkReservation4)
	reservation.MakeReservation(context.TODO(), makeReservation4)

	res1, err1 := user.AddUser(context.TODO(), user4)
	res2, err2 := user.DeleteUser(context.TODO(), user4)
	res3, err3 := reservation.DeleteReservation(context.TODO(), deleteReservation4)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)

	assert.Nil(t, err3)
	assert.False(t, res3.Success)

}

func fix() {
	user.GetUsers(context.TODO(), &proto.Request{})
}
