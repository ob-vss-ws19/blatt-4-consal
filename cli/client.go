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
	cm proto.CinemahallService
	mv proto.MovieService
	rv proto.ReservationService
	sw proto.ShowService
	us proto.UserService

	help = flag.Bool("help", false, "help")
)

func main() {

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("client <SERVICE> <FUNCTION> <PARAMS>")

		fmt.Println("<SERVICE>")
		fmt.Println("us")
		fmt.Println("	<FUNCTION>")
		fmt.Println("	- add <PARAMS>: name. | Example: client us add user1")
		fmt.Println("	- delete <PARAMS>: name. | Example: client us delete user1")
		fmt.Println("	- get: | Example: client us get")
		fmt.Println("")

		fmt.Println("cm")
		fmt.Println("	<FUNCTION>")
		fmt.Println("	- add <PARAMS>: name. | Example: client cm add cine1")
		fmt.Println("	- delete <PARAMS>: name. | Example: client cm delete cine1")
		fmt.Println("	- get: | Example: client cm get")
		fmt.Println("")

		fmt.Println("mv")
		fmt.Println("	<FUNCTION>")
		fmt.Println("	- add <PARAMS>: title. | Example: client mv add movie1")
		fmt.Println("	- delete <PARAMS>: title. | Example: client mv delete movie1")
		fmt.Println("	- get: | Example: client mv get")
		fmt.Println("")

		fmt.Println("rv")
		fmt.Println("	<FUNCTION>")
		fmt.Println("	- request <PARAMS>: user showingID seats. | Example: client rv request user1 2 4")
		fmt.Println("    Requests a reservation.")
		fmt.Println("  - book <PARAMS>: reservationID. | Example: client rv book 1")
		fmt.Println("    Books a reservation.")
		fmt.Println("  - delete <PARAMS>: reservationID. | Example: client rv delete 1")
		fmt.Println("  - get: | Example: client reservation get")
		fmt.Println("")

		fmt.Println("sw")
		fmt.Println("	<FUNCTION>")
		fmt.Println("	- add <PARAMS>: movie cm. | Example: client sw add movie1 cine1")
		fmt.Println("	- delete <PARAMS>: showingID. | Example: client sw delete 4")
		fmt.Println("	- get: | Example: client sw get")
		fmt.Println("")

		fmt.Println("fill")
		fmt.Println("  - Fills services with some data. | Example: client fill")
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
	case "us":
		us = proto.NewUserService("user", service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "add":
			// Füge neuen Benutzer hinzu
			us.GetUsers(context.TODO(), &proto.Request{})
			fmt.Println(us.AddUser(context.TODO(), &proto.UserRequest{
				Name: flag.Arg(2),
			}))
		case "delete":
			// Lösche Benutzer
			us.GetUsers(context.TODO(), &proto.Request{})
			information(us.DeleteUser(context.TODO(), &proto.UserRequest{
				Name: flag.Arg(2),
			}))
		case "get":
			// Gebe alle Benutzer aus
			us.GetUsers(context.TODO(), &proto.Request{})
			informationUser(us.GetUsers(context.TODO(), &proto.Request{}))
		}
	case "mv":
		mv = proto.NewMovieService("movie", service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "add":
			mv.GetMovies(context.TODO(), &proto.Request{})
			information(mv.AddMovie(context.TODO(), &proto.MovieRequest{
				MovieTitle: flag.Arg(2),
			}))
		case "delete":
			mv.GetMovies(context.TODO(), &proto.Request{})
			information(mv.DeleteMovie(context.TODO(), &proto.MovieRequest{
				MovieTitle: flag.Arg(2),
			}))
		case "get":
			mv.GetMovies(context.TODO(), &proto.Request{})
			informationMovie(mv.GetMovies(context.TODO(), &proto.Request{}))
		}
	case "cm":
		cm = proto.NewCinemahallService(("cinemahall"), service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "add":
			cm.GetCinemahalls(context.TODO(), &proto.Request{})
			information(cm.AddCinemahall(context.TODO(), &proto.CinemahallRequest{
				Name:            flag.Arg(2),
				SeatRows:        stringToInt(flag.Arg(3)),
				SeatRowCapacity: stringToInt(flag.Arg(4)),
			}))
		case "delete":
			cm.GetCinemahalls(context.TODO(), &proto.Request{})
			information(cm.DeleteCinemahall(context.TODO(), &proto.CinemahallRequest{
				Name: flag.Arg(2),
			}))
		case "get":
			cm.GetCinemahalls(context.TODO(), &proto.Request{})
			informationCinemahall(cm.GetCinemahalls(context.TODO(), &proto.Request{}))
		}
	case "rv":
		rv = proto.NewReservationService("reservation", service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "check":
			rv.ReservationInquiry(context.TODO(), &proto.ReservationRequest{
				UserName: flag.Arg(2),
				Show:     stringToInt(flag.Arg(3)),
				Seats:    stringToInt(flag.Arg(4)),
			})
		case "make":
			rv.MakeReservation(context.TODO(), &proto.ReservationRequest{
				ReservationId: stringToInt(flag.Arg(2)),
			})
		case "delete":
			rv.DeleteReservation(context.TODO(), &proto.ReservationRequest{
				ReservationId: stringToInt(flag.Arg(2)),
			})
		case "get":
			rv.GetReservations(context.TODO(), &proto.Request{})
		}
	case "sw":
		sw = proto.NewShowService("show", service.Client())
		secondFlag := flag.Arg(1)
		switch secondFlag {
		case "add":
			sw.GetShows(context.TODO(), &proto.Request{})
			information(sw.AddShow(context.TODO(), &proto.ShowRequest{
				Movie:      flag.Arg(2),
				CinemaHall: flag.Arg(3),
			}))
		case "delete":
			sw.GetShows(context.TODO(), &proto.Request{})
			information(sw.DeleteShow(context.TODO(), &proto.ShowRequest{
				Id: stringToInt(flag.Arg(2)),
			}))
		case "get":
			sw.GetShows(context.TODO(), &proto.Request{})
			informationShow(sw.GetShows(context.TODO(), &proto.Request{}))
		}

	case "fill":
		us = proto.NewUserService("user", service.Client())
		mv = proto.NewMovieService("movie", service.Client())
		cm = proto.NewCinemahallService("cinemahall", service.Client())
		sw = proto.NewShowService("show", service.Client())

		us.GetUsers(context.TODO(), &proto.Request{})
		information(us.AddUser(context.TODO(), &proto.UserRequest{
			Name: "Benutzer1",
		}))
		information(us.AddUser(context.TODO(), &proto.UserRequest{
			Name: "Benutzer2",
		}))
		information(us.AddUser(context.TODO(), &proto.UserRequest{
			Name: "Benutzer3",
		}))
		information(us.AddUser(context.TODO(), &proto.UserRequest{
			Name: "Benutzer4",
		}))
		information(mv.AddMovie(context.TODO(), &proto.MovieRequest{
			MovieTitle: "Spiderman",
		}))
		information(mv.AddMovie(context.TODO(), &proto.MovieRequest{
			MovieTitle: "Batman",
		}))
		information(cm.AddCinemahall(context.TODO(), &proto.CinemahallRequest{
			Name:            "Kino1",
			SeatRows:        10,
			SeatRowCapacity: 10,
		}))
		information(cm.AddCinemahall(context.TODO(), &proto.CinemahallRequest{
			Name:            "Kino2",
			SeatRows:        15,
			SeatRowCapacity: 15,
		}))
		information(sw.AddShow(context.TODO(), &proto.ShowRequest{
			CinemaHall: "Kino1",
			Movie:      "Spiderman",
		}))
		information(sw.AddShow(context.TODO(), &proto.ShowRequest{
			CinemaHall: "Kino2",
			Movie:      "Batman",
		}))

	default:
		// Falls falsch benutzt, Usagemöglichkeiten anzeigen
		flag.Usage()
		return
	}
}

func informationShow(res *proto.ShowResponse, error error) {
	if error == nil {
		fmt.Println(res)
	}
}

func informationCinemahall(res *proto.CinemahallResponse, error error) {
	if error == nil {
		fmt.Println(res)
	}
}

func informationUser(res *proto.UserResponse, error error) {
	if error == nil {
		fmt.Println(res)
	}
}

func informationMovie(res *proto.MovieResponse, error error) {
	if error == nil {
		fmt.Println(res)
	}
}

func information(res *proto.Response, error error) {
	if error != nil {
		fmt.Println(error)
		return
	}

	if res.Success {
		fmt.Printf("# %s\n", res.Message)
	}
}

func stringToInt(toParse string) int32 {
	newInt, _ := strconv.Atoi(toParse)
	return int32(newInt)
}
