package cli

import (
	"context"
	"blatt-4-consal/proto"
	"flag"
	"fmt"
	"github.com/micro/go-micro"
)

var (
	cm proto.CinemaService
	mv proto.MovieService
	rv proto.ReservationService
	sw proto.ShowService
	us proto.UserService
)

func main() {
	//TODO: Switch Cases
	service := micro.NewService(micro.Name("client"))
	service.Init()
	switch nil {
	case "cinema":
		cm = proto.NewCinemaService(("cinema"), service.Client())
		//service.Init();
		switch flag.Arg(1) {
		//TODO
		case "add":
			// Call the service method TODO: remove values
			rsp, err := cm.AddCinema(context.TODO(), &proto.CinemaRequest{Name: "Cinemaxx", SeatRows: 12, SeatRowCapacity: 5})
			if err != nil {
				fmt.Println(err)
				return
			}
			// Print response
			if rsp.Success {
				fmt.Printf("Success: %s", rsp.Message)
			} else {
				fmt.Printf("Error: %s", rsp.Message)
			}
		case "delete":
		case "get":
		}
	case "movie":
		//TODO: Add more cases and implementations
	default:
		return
	}

}
