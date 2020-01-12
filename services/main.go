package main

import (
	"blatt-4-consal/services/cinemahallservice"
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
	go userservice.StartUserService(context.TODO(), 3000)
	sleep()
	go movieservice.StartMovieService(context.TODO(), 3001)
	sleep()
	cinemahallservice.StartCinemaService(context.TODO(), 3002)

	// go reservationservice.StartReservationService()
	//showservice.StartReservationService()
	// time.Sleep(300*time.Millisecond)
}

func sleep() {
	time.Sleep(300 * time.Millisecond)
}
