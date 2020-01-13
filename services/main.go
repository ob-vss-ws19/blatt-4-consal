package main

import (
	"blatt-4-consal/services/cinemahallservice"
	"blatt-4-consal/services/movieservice"
	"blatt-4-consal/services/reservationservice"
	"blatt-4-consal/services/showservice"
	"blatt-4-consal/services/userservice"
	"context"
	"time"
	//"blatt-4-consal/services/cinemahallservice"
	//"blatt-4-consal/services/reservationservice"
	//"blatt-4-consal/services/showservice"
)

func main() {
	// start services asynchronously for user, movie, cinemahall, show and reservationservice
	go userservice.StartUserService(context.TODO(), 3000)
	sleep()
	go movieservice.StartMovieService(context.TODO(), 3001)
	sleep()
	go cinemahallservice.StartCinemaService(context.TODO(), 3002)
	sleep()
	go showservice.StartReservationService(context.TODO(), 3003)
	sleep()
	reservationservice.StartReservationService(context.TODO(), 3004)
}

func sleep() {
	time.Sleep(500 * time.Millisecond)
}
