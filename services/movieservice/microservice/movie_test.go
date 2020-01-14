package microservice

import (
	"blatt-4-consal/proto"
	cinemahallservice "blatt-4-consal/services/cinemahallservice/microservice"
	reservationservice "blatt-4-consal/services/reservationservice/microservice"
	showservice "blatt-4-consal/services/showservice/microservice"
	userservice "blatt-4-consal/services/userservice/microservice"

	"context"
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
var reservation proto.ReservationService
var user proto.UserService

func TestMovie(t *testing.T) {
	tmpContext, cancel = context.WithCancel(context.Background())
	go cinemahallservice.StartCinemaService(tmpContext, true)
	time.Sleep(500 * time.Millisecond)
	go StartMovieService(tmpContext, true)
	time.Sleep(500 * time.Millisecond)
	go showservice.StartShowService(tmpContext, true)
	time.Sleep(500 * time.Millisecond)
	go reservationservice.StartReservationService(tmpContext, true)
	time.Sleep(500 * time.Millisecond)
	go userservice.StartUserService(tmpContext, true)
	time.Sleep(500 * time.Millisecond)

	var cli client.Client

	//cinemahall := proto.NewCinemaService("cinema",cli)
	movie := proto.NewMovieService("movie",cli)
	//showing := proto.NewShowingService("showing",cli)

	//add movie 1
	req1 := &proto.MovieRequest{MovieTitle: "Herr der Ringe"}
	rsp, err := movie.AddMovie(context.TODO(), req1)
	assert.Nil(t, err);
	assert.True(t,rsp.Success)




}
