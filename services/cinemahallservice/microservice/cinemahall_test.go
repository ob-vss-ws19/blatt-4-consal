package microservice

import (
	"blatt-4-consal/proto"
	movieservice "blatt-4-consal/services/movieservice/microservice"
	reservationservice "blatt-4-consal/services/reservationservice/microservice"
	showservice "blatt-4-consal/services/showservice/microservice"
	userservice "blatt-4-consal/services/userservice/microservice"
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var cli client.Client
var tmpContext context.Context
var cancel context.CancelFunc
var cinemahall proto.CinemahallService
var movie proto.MovieService
var show proto.ShowService

func TestCinemahall (t *testing.T) {
	tmpContext, cancel = context.WithCancel(context.Background())
	go StartCinemaService(tmpContext,true)
	time.Sleep(500 * time.Millisecond)
	go showservice.StartShowService(tmpContext,true)
	time.Sleep(500 * time.Millisecond)
	go movieservice.StartMovieService(tmpContext,true)
	time.Sleep(500 * time.Millisecond)
	go reservationservice.StartReservationService(tmpContext,true)
	time.Sleep(500 * time.Millisecond)
	go userservice.StartUserService(tmpContext,true)

	var cli client.Client

	fmt.Println("Starting Cinemahall Microservice")
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)
	movie = proto.NewMovieService("movie", cli)

	// test Add Cinema
	//req1 := getNewCinemahall("testCinemaFirst", 4, 4)
	req1 := getNewCinemahall("testcinemafirst",4,10)
	rsp,err := cinemahall.AddCinemahall(tmpContext,req1)
	assert.Nil(t,err)
	assert.True(t,rsp.Success)
}

func getNewCinemahall(name string, SeatRows int32, SeatRowCapacity int32) *proto.CinemahallRequest {
	cinema := &proto.CinemahallRequest{
		Name:            name,
		SeatRows:        SeatRows,
		SeatRowCapacity: SeatRowCapacity,
	}
	return cinema
}

//func getCinemahallForDelete(name string) *proto.CinemahallRequest {
//	cinema := &proto.CinemahallRequest{
//		Name: name,
//	}
//	return cinema
//}
//
//func TestCreateCinemahall(t *testing.T) {
//	request1 := getNewCinemahall("kino1", 5, 5)
//	sleep()
//	res, err := cinemahall.AddCinemahall(tmpContext, request1)
//
//	assert.Nil(t, err)
//	assert.True(t, res.Success)
//}
//
//func TestCreateDoubleCinemahall(t *testing.T) {
//	request1 := getNewCinemahall("kino2", 5, 5)
//	sleep()
//	res1, err1 := cinemahall.AddCinemahall(tmpContext, request1)
//	sleep()
//	res2, err2 := cinemahall.AddCinemahall(tmpContext, request1)
//
//	assert.Nil(t, err1)
//	assert.True(t, res1.Success)
//
//	assert.Nil(t, err2)
//	assert.False(t, res2.Success)
//	cancel()
//}
//
//func TestCreateTripleCinemahall(t *testing.T) {
//
//	request1 := getNewCinemahall("kino3", 5, 5)
//	sleep()
//	res1, err1 := cinemahall.AddCinemahall(tmpContext, request1)
//	sleep()
//	res2, err2 := cinemahall.AddCinemahall(tmpContext, request1)
//	sleep()
//	res3, err3 := cinemahall.AddCinemahall(tmpContext, request1)
//
//	assert.Nil(t, err1)
//	assert.True(t, res1.Success)
//
//	assert.Nil(t, err2)
//	assert.False(t, res2.Success)
//
//	assert.Nil(t, err3)
//	assert.False(t, res3.Success)
//}

//func TestDeleteAddCinemahall(t *testing.T) {
//
//	request1 := getNewCinemahall("kino4", 5, 5)
//	request2 := getCinemahallForDelete("kino4")
//	request3 := &proto.MovieRequest{
//		MovieTitle: "film4",
//	}
//	request4 := &proto.ShowRequest{
//		CinemaHall: "kino4",
//		Movie:      "film4",
//	}
//	sleep()
//	res1, err1 := cinemahall.AddCinemahall(tmpContext, request1)
//	sleep()
//	movie.AddMovie(tmpContext, request3)
//	sleep()
//	show.AddShow(tmpContext, request4)
//	sleep()
//	res4, err4 := cinemahall.DeleteCinemahall(tmpContext, request2)
//	sleep()
//
//	assert.Nil(t, err1)
//	assert.True(t, res1.Success)
//
//	assert.Nil(t, err4)
//	assert.True(t, res4.Success)
//}

func sleep() {
	time.Sleep(1000 * time.Millisecond)
}

func fix() {
	cinemahall.GetCinemahalls(tmpContext, &proto.Request{})
}
