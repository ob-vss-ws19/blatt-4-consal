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

func TestCinemahall(t *testing.T) {
	tmpContext, cancel = context.WithCancel(context.Background())
	go StartCinemaService(tmpContext, true)
	time.Sleep(500 * time.Millisecond)
	go showservice.StartShowService(tmpContext, true)
	time.Sleep(500 * time.Millisecond)
	go movieservice.StartMovieService(tmpContext, true)
	time.Sleep(500 * time.Millisecond)
	go reservationservice.StartReservationService(tmpContext, true)
	time.Sleep(500 * time.Millisecond)
	go userservice.StartUserService(tmpContext, true)

	var cli client.Client

	fmt.Println("Starting Cinemahall Microservice")
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
	show = proto.NewShowService("show", cli)
	movie = proto.NewMovieService("movie", cli)

	// test Add Cinema first
	req1 := getNewCinemahall("testcinemafirst", 4, 10)
	rsp1, err1 := cinemahall.AddCinemahall(tmpContext, req1)
	assert.Nil(t, err1)
	assert.True(t, rsp1.Success)

	// test Add Cinema second
	req2 := getNewCinemahall("testcinemasecond", 4, 8)
	rsp2, err2 := cinemahall.AddCinemahall(tmpContext, req2)
	assert.Nil(t, err2)
	assert.True(t, rsp2.Success)

	// test Add Cinema first again -> error
	rsp3, err3 := cinemahall.AddCinemahall(tmpContext, req1)
	assert.Nil(t, err3)
	assert.False(t, rsp3.Success)

	// check if Cinemas are added
	emptyReq := &proto.Request{}
	rsp4, err4 := cinemahall.GetCinemahalls(tmpContext, emptyReq)
	assert.Nil(t, err4)
	assert.Len(t, rsp4.Value, 2)

	// delete cinema 1
	deleteReq := &proto.CinemahallRequest{Name: "testcinemafirst"}
	deleteRsp, err5 := cinemahall.DeleteCinemahall(tmpContext, deleteReq)
	assert.Nil(t, err5)
	assert.True(t, deleteRsp.Success)

}

func getNewCinemahall(name string, SeatRows int32, SeatRowCapacity int32) *proto.CinemahallRequest {
	cinema := &proto.CinemahallRequest{
		Name:            name,
		SeatRows:        SeatRows,
		SeatRowCapacity: SeatRowCapacity,
	}
	return cinema
}

func sleep() {
	time.Sleep(1000 * time.Millisecond)
}

func fix() {
	cinemahall.GetCinemahalls(tmpContext, &proto.Request{})
}
