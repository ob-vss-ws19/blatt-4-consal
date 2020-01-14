package microservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var cli client.Client
var tmpContext context.Context
var cancel context.CancelFunc
var cinemahall proto.CinemahallService

func initialize() (client.Client, context.Context, context.CancelFunc) {
	tmpContext, cancel := context.WithCancel(context.Background())
	go StartService(tmpContext, "cinemahall")
	var client client.Client
	return client, tmpContext, cancel
}

func StartService(context context.Context, servicename string) {
	//Create a new Service. Include name, version, address and context
	service := micro.NewService(
		micro.Name(servicename),
		micro.Version("latest"),
		micro.Context(context), //needed
	)
	//Register handler
	proto.RegisterCinemahallHandler(service.Server(), new(Cinemahall))
	fmt.Println("Cinemahall Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}

func getNewCinemahall(name string, SeatRows int32, SeatRowCapacity int32) *proto.CinemahallRequest {
	cinema := &proto.CinemahallRequest{
		Name:            name,
		SeatRows:        SeatRows,
		SeatRowCapacity: SeatRowCapacity,
	}
	return cinema
}

func init() {
	fmt.Println("Starting Cinemahall Microservice")
	cli, tmpContext, cancel = initialize()
	cinemahall = proto.NewCinemahallService("cinemahall", cli)
}

func TestCreateCinemahall(t *testing.T) {

	request1 := getNewCinemahall("kino1", 5, 5)
	fix()
	sleep()
	res, err := cinemahall.AddCinemahall(tmpContext, request1)

	assert.Nil(t, err)
	assert.True(t, res.Success)
}

func TestCreateDoubleCinemahall(t *testing.T) {

	request1 := getNewCinemahall("kino2", 5, 5)
	sleep()
	res1, err1 := cinemahall.AddCinemahall(tmpContext, request1)
	sleep()
	res2, err2 := cinemahall.AddCinemahall(tmpContext, request1)

	assert.Nil(t, err1)
	assert.True(t, res1.Success)

	assert.Nil(t, err2)
	assert.False(t, res2.Success)
	cancel()
}

//func TestCreateSameCinemahall(t *testing.T) {
//
//	request1 := getNewCinemahall("kino2", 5, 5)
//
//	//cinemahall2 := proto.NewCinemahallService("cinemahall", cli)
//	//request2 := getNewCinemahall("kino2", 5, 5)
//
//
//	cinemahall.GetCinemahalls(tmpContext, &proto.Request{})
//	res1, err1 := cinemahall.AddCinemahall(tmpContext, request1)
//	res2, err2 := cinemahall.AddCinemahall(tmpContext, request1)
//
//
//	assert.Nil(t, err1)
//	assert.True(t, res1.Success)
//
//	assert.Nil(t, err2)
//	assert.False(t, res2.Success)
//	cancel()
//}

func sleep() {
	time.Sleep(300 * time.Millisecond)
}

func fix() {
	cinemahall.GetCinemahalls(tmpContext, &proto.Request{})
}
