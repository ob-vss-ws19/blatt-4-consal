package main

import (
	"blatt-4-consal/services/cinemahallservice"
	"blatt-4-consal/services/movieservice"
	"blatt-4-consal/services/reservationservice"
	"blatt-4-consal/services/showservice"
	"blatt-4-consal/services/userservice"
)

func main() {
	//TODO
	//start services asynchronously for cinema, movie, reservation, show and user
	go cinemahallservice.StartCinemaService()
	go movieservice.StartMovieService()
	go reservationservice.StartReservationService()
	go showservice.StartReservationService()
	go userservice.StartMovieService()
}
