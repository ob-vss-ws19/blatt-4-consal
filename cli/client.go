package main

import (
	"blatt-4-consal/proto"
	"context"
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

	help = flag.Bool("help", false, "help")
)

func main() {

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("client.exe SERVICE FUNCTION PARAMS")
		fmt.Println("SERVICE")
		fmt.Println(" cm")
		fmt.Println("  FUNCTION")
		fmt.Println("  -add PARAMS: name. Example: client.exe cinema add hall1")
		fmt.Println("  -delete PARAMS: name. Example: client.exe cinema delete hall1")
		fmt.Println("  -get: Example: client.exe cinema get")
		fmt.Println(" mv")
		fmt.Println("  FUNCTION")
		fmt.Println("  -add PARAMS: title. Example: client.exe movie add shrek")
		fmt.Println("  -delete PARAMS: title. Example: client.exe movie delete shrek")
		fmt.Println("  -get: Example: client.exe movie get")
		fmt.Println(" rv")
		fmt.Println("  FUNCTION")
		fmt.Println("  -request PARAMS: user showingID seats. Example: client.exe reservation request sepp 2 4")
		fmt.Println("   Requests a reservation.")
		fmt.Println("  -book PARAMS: reservationID. Example: client.exe reservation book 1")
		fmt.Println("   Books a reservation.")
		fmt.Println("  -delete PARAMS: reservationID. Example: client.exe reservation delete 1")
		fmt.Println("  -get: Example: client.exe reservation get")
		fmt.Println(" sw")
		fmt.Println("  FUNCTION")
		fmt.Println("  -add PARAMS: movie cinema. Example: client.exe showing add shrek hall1")
		fmt.Println("  -delete PARAMS: showingID. Example: client.exe showing delete 4")
		fmt.Println("  -get: Example: client.exe showing get")
		fmt.Println(" us")
		fmt.Println("  FUNCTION")
		fmt.Println("  -add PARAMS: name. Example: client.exe user add sepp")
		fmt.Println("  -delete PARAMS: name. Example: client.exe user delete sepp")
		fmt.Println("  -get: Example: client.exe user get")
		fmt.Println(" fill")
		fmt.Println("  -Fills services with some data. Example: client.exe fill")
		return
	}

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	service := micro.NewService(micro.Name("client"))
	service.Init(micro.Address(fmt.Sprintf(": #{8091}")))

	//TODO: Switch Cases
	num := "cinema"
	switch num {
	case "cinema":
		cm = proto.NewCinemaService(("cinema"), service.Client())
		fmt.Print("hi")
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
