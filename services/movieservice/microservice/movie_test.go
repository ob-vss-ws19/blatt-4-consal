package microservice

import (
	"blatt-4-consal/proto"
	microservice2 "blatt-4-consal/services/cinemahallservice/microservice"
	microservice4 "blatt-4-consal/services/reservationservice/microservice"
	microservice3 "blatt-4-consal/services/showservice/microservice"
	"blatt-4-consal/services/userservice/microservice"
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
		proto.RegisterUserHandler(service.Server(), new(microservice.User))
		printInfo(microservicename)
	case "movie":
		proto.RegisterMovieHandler(service.Server(), new(Movie))
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

func TestMovie_AddMovie(t *testing.T) {
	testnum := "1"
	movie = proto.NewMovieService("movie", cli)
	req1 := &proto.MovieRequest{MovieTitle: "movie" + testnum}

	fix()
	res1, err := movie.AddMovie(context.TODO(), req1)

	assert.Nil(t, err)
	assert.True(t, res1.Success)
}

func TestMovie_AddDoubleMovie(t *testing.T) {
	testnum := "2"
	movie = proto.NewMovieService("movie", cli)
	req1 := &proto.MovieRequest{MovieTitle: "movie" + testnum}

	fix()
	res1, err1 := movie.AddMovie(context.TODO(), req1)
	res2, err2 := movie.AddMovie(context.TODO(), req1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.False(t, res2.Success)
}

func TestMovie_AddTripleMovie(t *testing.T) {
	testnum := "3"
	movie = proto.NewMovieService("movie", cli)
	req1 := &proto.MovieRequest{MovieTitle: "movie" + testnum}

	fix()
	res1, err1 := movie.AddMovie(context.TODO(), req1)
	res2, err2 := movie.AddMovie(context.TODO(), req1)
	res3, err3 := movie.AddMovie(context.TODO(), req1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.False(t, res2.Success)

	assert.Nil(t, err3)
	assert.False(t, res3.Success)
}

func TestMovie_DeleteMovie_WithoutShow(t *testing.T) {
	testnum := "4"
	movie = proto.NewMovieService("movie", cli)

	mv := "movie" + testnum

	movie4 := &proto.MovieRequest{MovieTitle: mv}

	fix()
	res1, err1 := movie.AddMovie(context.TODO(), movie4)
	res2, err2 := movie.DeleteMovie(context.TODO(), movie4)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)
}

func TestMovie_DeleteMovie_WithShow(t *testing.T) {
	testnum := "5"

	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)

	mv := "movie" + testnum
	cm := "cinemahall" + testnum

	movie5 := &proto.MovieRequest{MovieTitle: mv}
	cinemahall5 := &proto.CinemahallRequest{Name: cm, SeatRows: 10, SeatRowCapacity: 10}
	show5 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	deleteShow5 := &proto.ShowRequest{Id: 1}

	fix()
	cinemahall.AddCinemahall(context.TODO(), cinemahall5)
	res1, err1 := movie.AddMovie(context.TODO(), movie5)
	show.AddShow(context.TODO(), show5)

	res2, err2 := movie.DeleteMovie(context.TODO(), movie5)
	res3, err3 := show.DeleteShow(context.TODO(), deleteShow5)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)

	assert.Nil(t, err3)
	assert.False(t, res3.Success)
}

func fix() {
	movie.GetMovies(context.TODO(), &proto.Request{})
}
