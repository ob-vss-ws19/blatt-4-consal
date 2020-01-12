package main

import (
	"blatt-4-consal/services/movieservice"
	"time"

	//"blatt-4-consal/services/cinemahallservice"
	//"blatt-4-consal/services/movieservice"
	//"blatt-4-consal/services/reservationservice"
	//"blatt-4-consal/services/showservice"
	"blatt-4-consal/services/userservice"
	"context"
)

func main() {
	//start services asynchronously for cinema, movie, reservation, show and user
	go userservice.StartUserService(context.TODO())
	time.Sleep(1000 * time.Millisecond)
	movieservice.StartMovieService(context.TODO())
	//cinemahallservice.StartCinemaService(context.TODO())

	// go reservationservice.StartReservationService()
	// go showservice.StartReservationService()
	// time.Sleep(300*time.Millisecond)

}
