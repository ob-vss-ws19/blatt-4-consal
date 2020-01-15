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

	//cinemahall := proto.NewCinemahallService("cinema",cli)
	//show := proto.NewShowService("showing",cli)
	movie := proto.NewMovieService("movie", cli)

	//add first movie
	req1 := getNewMovie("Herr der Ringe")
	rsp1, err1 := movie.AddMovie(tmpContext, req1)
	assert.Nil(t, err1);
	assert.True(t, rsp1.Success)

	// add first movie again -> error
	rsp2,err2 := movie.AddMovie(tmpContext,req1);
	assert.Nil(t, err2)
	assert.False(t,rsp2.Success)
}

func getNewMovie(name string, ) *proto.MovieRequest {
	movie := &proto.MovieRequest{
		MovieTitle: name,
	}
	return movie
}
