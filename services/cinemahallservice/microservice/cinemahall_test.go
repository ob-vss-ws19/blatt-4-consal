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

func TestCinemahall(t *testing.T) {
	tmpContext, cancel := context.WithCancel(context.Background())
	servicename := "cinemahall"
	go StartService(tmpContext, servicename)
	time.Sleep(300 * time.Millisecond)

	var client client.Client
	cinemahall := proto.NewCinemahallService("cinemahall", client)
	req := &proto.CinemahallRequest{
		Name:            "kino1",
		SeatRows:        5,
		SeatRowCapacity: 5,
	}
	cinemahall.GetCinemahalls(tmpContext, &proto.Request{})
	res, err := cinemahall.AddCinemahall(tmpContext, req)
	assert.Nil(t, err)
	//assert.True(t, res.Success)

	fmt.Print(cinemahall)
	fmt.Print(res)

	cancel()
}
