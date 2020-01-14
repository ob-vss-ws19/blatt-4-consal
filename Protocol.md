# Kommunikation zwischen Services

## Beschreibung
Die Microservices kommunizieren über messages, die von `go/micro` gehandelt werden.

### Request an ein Microservice

Wenn der Service über `micro.NewService` mit dem Namen `reservation` und einem leeren Context `context.TODO()` initialisiert und gestartet wurde,
kann man mit

    var client client.Client
    reservationService := proto.NewReservationService("reservation", client)
    res, err := reservationService.GetReservations(context.TODO(), &proto.Request{})

Requests an den Microservice mit den Namen `reservation` schicken.
Hier Beispielsweise ein Request, um alle Reservationen zu erhalten.

### Response an ein Microservice

Um ein Reponse zu senden, verwendet man die in den Parametern enthaltene `res *proto.ShowResponse`.

Beispielsweise wäre dies eine Antwort, als alle Reservationen an den Client als message zurückzugeben.

    var allReservations string

[...]

    res.Success = true
    res.Message = allReservations


