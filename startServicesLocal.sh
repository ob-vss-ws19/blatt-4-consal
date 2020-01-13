echo "Start User"
go run ./services/userservice/main.go &
sleep 0.2
echo "Start Movieservice"
go run ./services/movieservice/main.go &
sleep 0.2
echo "Start Cinemahall"
go run ./services/cinemahallservice/main.go &
sleep 0.2
echo "Start Show"
go run ./services/showservice/main.go &
sleep 0.2
echo "Start Reservation"
go run ./services/reservationservice/main.go &
