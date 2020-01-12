package main

import (
	"blatt-4-consal/services/movieservice"
	"blatt-4-consal/services/userservice"
	"context"
	"time"
	//"blatt-4-consal/services/cinemahallservice"
	//"blatt-4-consal/services/reservationservice"
	//"blatt-4-consal/services/showservice"
)

func main() {
	//start services asynchronously for cinema, movie, reservation, show and user
	go movieservice.StartMovieService(context.TODO())
	time.Sleep(300 * time.Millisecond)

	userservice.StartUserService(context.TODO())

	//cinemahallservice.StartCinemaService(context.TODO())

	// go reservationservice.StartReservationService()
	//showservice.StartReservationService()
	// time.Sleep(300*time.Millisecond)

}
