package microservice

import (
	"blatt-4-consal/proto"
	microservice2 "blatt-4-consal/services/cinemahallservice/microservice"
	"blatt-4-consal/services/movieservice/microservice"
	microservice3 "blatt-4-consal/services/showservice/microservice"
	microservice4 "blatt-4-consal/services/userservice/microservice"
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
var parallelIdReservation int32 = 1
var parallelIdShow int32 = 1

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

// Zähle Id immer hoch bevor CheckReservation, um mitzuzählen
func countIdReservation(add int32) {
	parallelIdReservation = parallelIdReservation + add
}

// Zähle Id immer hoch bevor CheckReservation, um mitzuzählen
func countIdShow(add int32) {
	parallelIdShow = parallelIdShow + add
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
		proto.RegisterUserHandler(service.Server(), new(microservice4.User))
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
		proto.RegisterReservationHandler(service.Server(), new(Reservation))
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

func TestReservation_AddDeleteReservation(t *testing.T) {
	testnum := "1"
	user = proto.NewUserService("user", cli)
	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)
	reservation = proto.NewReservationService("reservation", cli)

	us := "user" + testnum
	mv := "movie" + testnum
	cm := "cinemahall" + testnum

	user1 := &proto.UserRequest{Name: us}
	movie1 := &proto.MovieRequest{MovieTitle: mv}
	cinemahall1 := &proto.CinemahallRequest{Name: cm, SeatRows: 10, SeatRowCapacity: 10}
	show1 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	checkReservation1 := &proto.ReservationRequest{UserName: us, Show: parallelIdShow, Seats: 5}
	makeReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}
	deleteReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}

	fix()
	user.AddUser(context.TODO(), user1)
	movie.AddMovie(context.TODO(), movie1)
	cinemahall.AddCinemahall(context.TODO(), cinemahall1)
	show.AddShow(context.TODO(), show1)
	countIdShow(1)
	res1, err1 := reservation.ReservationInquiry(context.TODO(), checkReservation1)
	countIdReservation(1)
	res2, err2 := reservation.MakeReservation(context.TODO(), makeReservation1)
	res3, err3 := reservation.DeleteReservation(context.TODO(), deleteReservation1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)

	assert.Nil(t, err3)
	assert.True(t, res3.Success)
}

func TestReservation_DeleteMovie_Of_Reservation(t *testing.T) {
	testnum := "2"
	user = proto.NewUserService("user", cli)
	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)
	reservation = proto.NewReservationService("reservation", cli)

	us := "user" + testnum
	mv := "movie" + testnum
	cm := "cinemahall" + testnum

	user1 := &proto.UserRequest{Name: us}
	movie1 := &proto.MovieRequest{MovieTitle: mv}
	cinemahall1 := &proto.CinemahallRequest{Name: cm, SeatRows: 10, SeatRowCapacity: 10}
	show1 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	checkReservation1 := &proto.ReservationRequest{UserName: us, Show: parallelIdShow, Seats: 5}
	makeReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}
	deleteReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}

	fix()
	user.AddUser(context.TODO(), user1)
	movie.AddMovie(context.TODO(), movie1)
	cinemahall.AddCinemahall(context.TODO(), cinemahall1)
	show.AddShow(context.TODO(), show1)
	countIdShow(1)
	res1, err1 := reservation.ReservationInquiry(context.TODO(), checkReservation1)
	countIdReservation(1)
	res2, err2 := reservation.MakeReservation(context.TODO(), makeReservation1)
	res3, err3 := movie.DeleteMovie(context.TODO(), movie1)
	res4, err4 := reservation.DeleteReservation(context.TODO(), deleteReservation1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)

	assert.Nil(t, err3)
	assert.True(t, res3.Success)

	assert.Nil(t, err4)
	assert.False(t, res4.Success)
}

func TestReservation_DeleteCinemahall_Of_Reservation(t *testing.T) {
	testnum := "3"
	user = proto.NewUserService("user", cli)
	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)
	reservation = proto.NewReservationService("reservation", cli)

	us := "user" + testnum
	mv := "movie" + testnum
	cm := "cinemahall" + testnum

	user1 := &proto.UserRequest{Name: us}
	movie1 := &proto.MovieRequest{MovieTitle: mv}
	cinemahall1 := &proto.CinemahallRequest{Name: cm, SeatRows: 10, SeatRowCapacity: 10}
	deleteCinemahall1 := &proto.CinemahallRequest{Name: cm}
	show1 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	checkReservation1 := &proto.ReservationRequest{UserName: us, Show: parallelIdShow, Seats: 5}
	makeReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}
	deleteReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}

	fix()
	user.AddUser(context.TODO(), user1)
	movie.AddMovie(context.TODO(), movie1)
	cinemahall.AddCinemahall(context.TODO(), cinemahall1)
	show.AddShow(context.TODO(), show1)
	countIdShow(1)
	res1, err1 := reservation.ReservationInquiry(context.TODO(), checkReservation1)
	countIdReservation(1)
	res2, err2 := reservation.MakeReservation(context.TODO(), makeReservation1)
	res3, err3 := cinemahall.DeleteCinemahall(context.TODO(), deleteCinemahall1)
	res4, err4 := reservation.DeleteReservation(context.TODO(), deleteReservation1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)

	assert.Nil(t, err3)
	assert.True(t, res3.Success)

	assert.Nil(t, err4)
	assert.False(t, res4.Success)
}

func TestReservation_DeleteUser_Of_Reservation(t *testing.T) {
	testnum := "4"
	user = proto.NewUserService("user", cli)
	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)
	reservation = proto.NewReservationService("reservation", cli)

	us := "user" + testnum
	mv := "movie" + testnum
	cm := "cinemahall" + testnum

	user1 := &proto.UserRequest{Name: us}
	movie1 := &proto.MovieRequest{MovieTitle: mv}
	cinemahall1 := &proto.CinemahallRequest{Name: cm, SeatRows: 10, SeatRowCapacity: 10}
	show1 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	checkReservation1 := &proto.ReservationRequest{UserName: us, Show: parallelIdShow, Seats: 5}
	makeReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}
	deleteReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}

	fix()
	user.AddUser(context.TODO(), user1)
	movie.AddMovie(context.TODO(), movie1)
	cinemahall.AddCinemahall(context.TODO(), cinemahall1)
	show.AddShow(context.TODO(), show1)
	countIdShow(1)
	res1, err1 := reservation.ReservationInquiry(context.TODO(), checkReservation1)
	countIdReservation(1)
	res2, err2 := reservation.MakeReservation(context.TODO(), makeReservation1)
	res3, err3 := user.DeleteUser(context.TODO(), user1)
	res4, err4 := reservation.DeleteReservation(context.TODO(), deleteReservation1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.True(t, res2.Success)

	assert.Nil(t, err3)
	assert.True(t, res3.Success)

	assert.Nil(t, err4)
	assert.False(t, res4.Success)
}

func TestReservation_NotEnoughSeats(t *testing.T) {
	testnum := "5"
	user = proto.NewUserService("user", cli)
	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)
	reservation = proto.NewReservationService("reservation", cli)

	us := "user" + testnum
	mv := "movie" + testnum
	cm := "cinemahall" + testnum

	user1 := &proto.UserRequest{Name: us}
	movie1 := &proto.MovieRequest{MovieTitle: mv}
	cinemahall1 := &proto.CinemahallRequest{Name: cm, SeatRows: 2, SeatRowCapacity: 4}
	show1 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	checkReservation1 := &proto.ReservationRequest{UserName: us, Show: parallelIdShow, Seats: 10}
	makeReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}

	fix()
	user.AddUser(context.TODO(), user1)
	movie.AddMovie(context.TODO(), movie1)
	cinemahall.AddCinemahall(context.TODO(), cinemahall1)
	show.AddShow(context.TODO(), show1)
	countIdShow(1)
	res1, err1 := reservation.ReservationInquiry(context.TODO(), checkReservation1)
	countIdReservation(1)
	res2, err2 := reservation.MakeReservation(context.TODO(), makeReservation1)

	assert.Nil(t, err1)
	assert.False(t, res1.Success)

	assert.Nil(t, err2)
	assert.False(t, res2.Success)
}

func TestReservation_FirstUserWasFasterAndNotEnoughSeats(t *testing.T) {
	testnum := "6"
	user = proto.NewUserService("user", cli)
	movie = proto.NewMovieService("movie", cli)
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)
	reservation = proto.NewReservationService("reservation", cli)

	us := "user" + testnum
	mv := "movie" + testnum
	cm := "cinemahall" + testnum

	user1 := &proto.UserRequest{Name: us}
	movie1 := &proto.MovieRequest{MovieTitle: mv}
	cinemahall1 := &proto.CinemahallRequest{Name: cm, SeatRows: 2, SeatRowCapacity: 5}
	show1 := &proto.ShowRequest{CinemaHall: cm, Movie: mv}
	checkReservation1 := &proto.ReservationRequest{UserName: us, Show: parallelIdShow, Seats: 8}
	makeReservation1 := &proto.ReservationRequest{ReservationId: parallelIdReservation}

	usx := "user" + testnum + testnum

	user2 := &proto.UserRequest{Name: usx}
	checkReservation2 := &proto.ReservationRequest{UserName: usx, Show: parallelIdShow, Seats: 5} // only 2 can be set in this scene
	makeReservation2 := &proto.ReservationRequest{ReservationId: parallelIdReservation}

	fix()
	user.AddUser(context.TODO(), user1)
	user.AddUser(context.TODO(), user2)
	movie.AddMovie(context.TODO(), movie1)
	cinemahall.AddCinemahall(context.TODO(), cinemahall1)
	show.AddShow(context.TODO(), show1)
	countIdShow(1)
	res11, err11 := reservation.ReservationInquiry(context.TODO(), checkReservation1)
	res12, err12 := reservation.ReservationInquiry(context.TODO(), checkReservation2)
	countIdReservation(1)
	countIdReservation(1)
	res21, err21 := reservation.MakeReservation(context.TODO(), makeReservation1)
	res22, err22 := reservation.MakeReservation(context.TODO(), makeReservation2)

	assert.Nil(t, err11)
	assert.True(t, res11.Success)

	assert.Nil(t, err12)
	assert.True(t, res12.Success)

	assert.Nil(t, err21)
	assert.True(t, res21.Success)

	assert.Nil(t, err22)
	assert.False(t, res22.Success)
}

func fix() {
	reservation.GetReservations(context.TODO(), &proto.Request{})
}
