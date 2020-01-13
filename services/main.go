package main

import (
	microservice3 "blatt-4-consal/services/cinemahallservice/microservice"
	microservice2 "blatt-4-consal/services/movieservice/microservice"
	microservice5 "blatt-4-consal/services/reservationservice/microservice"
	microservice4 "blatt-4-consal/services/showservice/microservice"
	microservice1 "blatt-4-consal/services/userservice/microservice"
	"context"
	"time"
)

func main() {
	// start services asynchronously for user, movie, cinemahall, show and reservationservice
	go microservice1.StartUserService(context.TODO(), 3000)
	sleep()
	go microservice2.StartMovieService(context.TODO(), 3001)
	sleep()
	go microservice3.StartCinemaService(context.TODO(), 3002)
	sleep()
	go microservice4.StartShowService(context.TODO(), 3003)
	sleep()
	microservice5.StartReservationService(context.TODO(), 3004)
}

func sleep() {
	time.Sleep(500 * time.Millisecond)
}
