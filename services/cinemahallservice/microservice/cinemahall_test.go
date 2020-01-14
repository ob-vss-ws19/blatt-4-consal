package microservice

//
//import (
//	"blatt-4-consal/proto"
//	movieservice "blatt-4-consal/services/movieservice/microservice"
//	showservice "blatt-4-consal/services/showservice/microservice"
//	"context"
//	"fmt"
//	"github.com/micro/go-micro"
//	"github.com/micro/go-micro/client"
//	"github.com/stretchr/testify/assert"
//	_ "github.com/stretchr/testify/assert"
//
//	"testing"
//	"time"
//)
//
//var cli client.Client
//var tmpContext context.Context
//var cancel context.CancelFunc
//var cinemahall proto.CinemahallService
//var movie proto.MovieService
//var show proto.ShowService
//
//func init() {
//	fmt.Println("Starting Cinemahall Microservice")
//	cli, tmpContext, cancel = initialize()
//	cinemahall = proto.NewCinemahallService("cinemahall", cli)
//	show = proto.NewShowService("show", cli)
//	movie = proto.NewMovieService("movie", cli)
//	fix()
//}
//
//func initialize() (client.Client, context.Context, context.CancelFunc) {
//	tmpContext, cancel = context.WithCancel(context.Background())
//	go StartService(tmpContext, "cinemahall", 3000)
//	sleep()
//	go StartService(tmpContext, "show", 3001)
//	sleep()
//	go StartService(tmpContext, "movie", 3002)
//	sleep()
//	var client client.Client
//	return client, tmpContext, cancel
//}
//
//func StartService(context context.Context, servicename string, port int32) {
//	//Create a new Service. Include name, version, address and context
//
//	service := micro.NewService(
//		micro.Name(servicename),
//		micro.Version("latest"),
//		micro.Context(context), //needed
//		micro.Address(fmt.Sprintf(":%v", port)),
//	)
//	service.Init()
//
//	//Register handler
//	switch servicename {
//	case "cinemahall":
//		proto.RegisterCinemahallHandler(service.Server(), new(Cinemahall))
//	case "movie":
//		proto.RegisterMovieHandler(service.Server(), new(movieservice.Movie))
//	case "show":
//		proto.RegisterShowHandler(service.Server(), new(showservice.Show))
//	}
//	fmt.Println("Service starting...")
//	//Run the Server
//	if err := service.Run(); err != nil {
//		//Print error message if there is any
//		fmt.Println(err)
//	}
//}
//
//func getNewCinemahall(name string, SeatRows int32, SeatRowCapacity int32) *proto.CinemahallRequest {
//	cinema := &proto.CinemahallRequest{
//		Name:            name,
//		SeatRows:        SeatRows,
//		SeatRowCapacity: SeatRowCapacity,
//	}
//	return cinema
//}
//
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
//
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
//
//func sleep() {
//	time.Sleep(1000 * time.Millisecond)
//}
//
//func fix() {
//	cinemahall.GetCinemahalls(tmpContext, &proto.Request{})
//}
