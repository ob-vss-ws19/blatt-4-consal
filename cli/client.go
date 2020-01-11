package main

import (
	"blatt-4-consal/proto"
	"context"
	"flag"
	"fmt"
	"github.com/micro/go-micro"
	"strconv"
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
	service.Init(micro.Address(fmt.Sprintf(":%v", 8091)))

	firstFlag := flag.Arg(0)
	switch firstFlag {
	case "cm":
		cm = proto.NewCinemaService(("cinema"), service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "add":
			cm.AddCinema(context.TODO(), &proto.CinemaRequest{
				Name:            flag.Arg(2),
				SeatRows:        stringToInt(flag.Arg(3)),
				SeatRowCapacity: stringToInt(flag.Arg(4)),
			})
		case "delete":
			cm.DeleteCinema(context.TODO(), &proto.CinemaRequest{
				Name: flag.Arg(2),
			})
		case "get":
			cm.GetCinemas(context.TODO(), &proto.Request{})
		}
	case "mv":
		mv = proto.NewMovieService("movie", service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "add":
			mv.AddMovie(context.TODO(), &proto.MovieRequest{
				MovieTitle: flag.Arg(2),
			})
		case "delete":
			mv.DeleteMovie(context.TODO(), &proto.MovieRequest{
				MovieTitle: flag.Arg(2),
			})
		case "get":
			mv.GetMovies(context.TODO(), &proto.Request{})
		}
	case "rv":
		rv = proto.NewReservationService("reservation", service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "check":
			rv.CheckReservation(context.TODO(), &proto.ReservationRequest{
				User:    flag.Arg(2),
				Showing: stringToInt(flag.Arg(3)),
				Seats:   stringToInt(flag.Arg(4)),
			})
		case "make":
			rv.MakeReservation(context.TODO(), &proto.ReservationRequest{
				ReservationID: stringToInt(flag.Arg(2)),
			})
		case "delete":
			rv.DeleteReservation(context.TODO(), &proto.ReservationRequest{
				ReservationID: stringToInt(flag.Arg(2)),
			})
		case "get":
			rv.GetReservations(context.TODO(), &proto.Request{})
		}
	case "sw":
		sw = proto.NewShowService("show", service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "add":
			sw.AddShow(context.TODO(), &proto.ShowRequest{
				Movie:      flag.Arg(2),
				CinemaHall: flag.Arg(3),
			})
		case "delete":
			sw.DeleteShow(context.TODO(), &proto.ShowRequest{
				Id: stringToInt(flag.Arg(2)),
			})
		case "get":
			sw.GetShows(context.TODO(), &proto.Request{})
		}
	case "us":
		us = proto.NewUserService("User", service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "add":
			// Füge neuen Benutzer hinzu
			fmt.Print(us.AddUser(context.TODO(), &proto.UserRequest{
				Name: flag.Arg(2),
			}))

			//case "delete":
			//	// Lösche Benutzer
			//	us.DeleteUser(context.TODO(), &proto.UserRequest{
			//		Name: flag.Arg(2),
			//	})
			//case "get":
			//	// Gebe alle Benutzer aus
			//	us.GetUsers(context.TODO(), &proto.Request{})
			//}

		case "fill":
		}
	default:
		// Falls falsch benutzt, Usagemöglichkeiten anzeigen
		flag.Usage()
		return
	}

}

func stringToInt(toParse string) int32 {
	newInt, _ := strconv.Atoi(toParse)
	return int32(newInt)
}
