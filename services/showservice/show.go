package showservice

import (
	"blatt-4-consal/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type Show struct {
	Shows map[string]*ShowRequest
	Id    int32
}

type ShowRequest struct {
	Movie  string
	Cinema string
}

func (sw *Show) AddShowing(ctx context.Context, req *proto.ShowRequest, rsp *proto.Response) error {
	return nil
}

func (sw *Show) DeleteShowing(ctx context.Context, req *proto.ShowRequest, rsp *proto.Response) error {
	return nil
}

func (sw *Show) GetShowings(ctx context.Context, req *proto.Request, rsp *proto.ShowResponse) error {
	return nil
}

//Start Service for Show class
func StartReservationService() {
	//Create a new Service. Add name address and context
	service := micro.NewService(
		micro.Name("show"),
	)
	// Init will parse the command line flags
	service.Init()
	//Register handler
	proto.RegisterShowHandler(service.Server(), new(Show))
	fmt.Println("Reservation Service starting...")
	//Run the Server
	if err := service.Run(); err != nil {
		//Print error message if there is any
		fmt.Println(err)
	}
}
