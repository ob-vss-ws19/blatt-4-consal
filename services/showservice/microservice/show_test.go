package microservice

import (
	"blatt-4-consal/proto"
	microservice3 "blatt-4-consal/services/cinemahallservice/microservice"
	microservice2 "blatt-4-consal/services/movieservice/microservice"
	microservice4 "blatt-4-consal/services/reservationservice/microservice"
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
var parallelId int32 = 1

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
		proto.RegisterMovieHandler(service.Server(), new(microservice2.Movie))
		printInfo(microservicename)
	case "cinemahall":
		proto.RegisterCinemahallHandler(service.Server(), new(microservice3.Cinemahall))
		printInfo(microservicename)
	case "show":
		proto.RegisterShowHandler(service.Server(), new(Show))
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

func TestShow_AddShow(t *testing.T) {
	countId(1)
	testnum := "1"
	var seatnum int32 = 10
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	movie = proto.NewMovieService("movie", cli)
	show = proto.NewShowService("show", cli)

	cm := "cinemahall" + testnum
	mv := "movie" + testnum

	cinemahall1 := &proto.CinemahallRequest{Name: cm, SeatRows: seatnum, SeatRowCapacity: seatnum}
	movie1 := &proto.MovieRequest{MovieTitle: mv}
	req1 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	//deleteCinemahall1 := &proto.CinemahallRequest{Name: cm + testnum}
	//deleteShow1 := &proto.ShowRequest{Id: 1}

	fix()
	movie.AddMovie(context.TODO(), movie1)
	cinemahall.AddCinemahall(context.TODO(), cinemahall1)
	res1, err1 := show.AddShow(context.TODO(), req1)
	fmt.Println(res1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)
}

func TestShow_AddDoubleShow(t *testing.T) {
	countId(2)
	testnum := "2"
	var seatnum int32 = 10
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	movie = proto.NewMovieService("movie", cli)
	show = proto.NewShowService("show", cli)

	cm := "cinemahall" + testnum
	mv := "movie" + testnum

	cinemahall2 := &proto.CinemahallRequest{Name: cm, SeatRows: seatnum, SeatRowCapacity: seatnum}
	movie2 := &proto.MovieRequest{MovieTitle: mv}
	req1 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	//deleteCinemahall1 := &proto.CinemahallRequest{Name: cm + testnum}
	//deleteShow1 := &proto.ShowRequest{Id: 1}

	fix()
	movie.AddMovie(context.TODO(), movie2)
	cinemahall.AddCinemahall(context.TODO(), cinemahall2)
	res1, err1 := show.AddShow(context.TODO(), req1)
	res2, err2 := show.AddShow(context.TODO(), req1)

	fmt.Println(res1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)
}

func TestShow_AddTripleShow(t *testing.T) {
	countId(3)
	testnum := "3"
	var seatnum int32 = 10
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	movie = proto.NewMovieService("movie", cli)
	show = proto.NewShowService("show", cli)

	cm := "cinemahall" + testnum
	mv := "movie" + testnum

	cinemahall3 := &proto.CinemahallRequest{Name: cm, SeatRows: seatnum, SeatRowCapacity: seatnum}
	movie3 := &proto.MovieRequest{MovieTitle: mv}
	req1 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	//deleteCinemahall1 := &proto.CinemahallRequest{Name: cm + testnum}
	//deleteShow1 := &proto.ShowRequest{Id: 1}

	fix()
	movie.AddMovie(context.TODO(), movie3)
	cinemahall.AddCinemahall(context.TODO(), cinemahall3)
	res1, err1 := show.AddShow(context.TODO(), req1)
	res2, err2 := show.AddShow(context.TODO(), req1)
	res3, err3 := show.AddShow(context.TODO(), req1)

	fmt.Println(res1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)

	assert.Nil(t, err3)
	assert.True(t, res3.Success)
}

func TestShow_DeleteCinemahall_Of_Show(t *testing.T) {
	countId(1)
	testnum := "4"
	var seatnum int32 = 10
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	movie = proto.NewMovieService("movie", cli)
	show = proto.NewShowService("show", cli)

	cm := "cinemahall" + testnum
	mv := "movie" + testnum

	addCinemahall4 := &proto.CinemahallRequest{Name: cm, SeatRows: seatnum, SeatRowCapacity: seatnum}
	deleteCinemahall4 := &proto.CinemahallRequest{Name: cm}
	movie4 := &proto.MovieRequest{MovieTitle: mv}
	show4 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	deleteShow4 := &proto.ShowRequest{Id: parallelId}

	fix()
	movie.AddMovie(context.TODO(), movie4)
	res1, err1 := cinemahall.AddCinemahall(context.TODO(), addCinemahall4)
	show.AddShow(context.TODO(), show4)

	res2, err2 := cinemahall.DeleteCinemahall(context.TODO(), deleteCinemahall4)
	res3, err3 := show.DeleteShow(context.TODO(), deleteShow4)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)

	assert.Nil(t, err3)
	assert.False(t, res3.Success)
}

func TestShow_DeleteMovie_Of_Show(t *testing.T) {
	countId(1)
	testnum := "5"
	var seatnum int32 = 10
	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)

	mv := "movie" + testnum
	cm := "cinemahall" + testnum

	addCinemahall5 := &proto.CinemahallRequest{Name: cm, SeatRows: seatnum, SeatRowCapacity: seatnum}
	movie5 := &proto.MovieRequest{MovieTitle: mv}
	show5 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	deleteShow5 := &proto.ShowRequest{Id: parallelId}

	fix()
	res1, err1 := movie.AddMovie(context.TODO(), movie5)
	cinemahall.AddCinemahall(context.TODO(), addCinemahall5)
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

func TestShow_DeleteShow(t *testing.T) {

	testnum := "6"
	var seatnum int32 = 10
	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)

	mv := "movie" + testnum
	cm := "cinemahall" + testnum

	addCinemahall6 := &proto.CinemahallRequest{Name: cm, SeatRows: seatnum, SeatRowCapacity: seatnum}
	movie6 := &proto.MovieRequest{MovieTitle: mv}
	show6 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	deleteShow6 := &proto.ShowRequest{Id: parallelId}

	fix()
	movie.AddMovie(context.TODO(), movie6)

	cinemahall.AddCinemahall(context.TODO(), addCinemahall6)
	res1, err1 := show.AddShow(context.TODO(), show6)
	countId(1)
	res2, err2 := show.DeleteShow(context.TODO(), deleteShow6)
	res3, err3 := show.DeleteShow(context.TODO(), deleteShow6)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)

	assert.Nil(t, err3)
	assert.False(t, res3.Success)
}

func fix() {
	show.GetShows(context.TODO(), &proto.Request{})
}

// Zähle Id immer hoch bevor AddShow, um mitzuzählen (braucht man beim löschen einer Show)
func countId(add int32) {
	parallelId = parallelId + add
}
