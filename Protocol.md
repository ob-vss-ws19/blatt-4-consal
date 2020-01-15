# Kommunikation zwischen Services

## Beschreibung
Die Microservices kommunizieren über messages, die von `go/micro` gehandelt werden.

### Request an ein Microservice

Wenn der Service über `micro.NewService` mit dem Namen `reservation` und einen leeren Context `context.TODO()` initialisiert und gestartet wurde,
kann man mit

    var client client.Client
    reservationService := proto.NewReservationService("reservation", client)
    res, err := reservationService.GetReservations(context.TODO(), &proto.Request{})

Requests an den Microservice mit dem Namen `reservation` schicken.
Hier Beispielsweise ein Request, um alle Reservierungen zu erhalten.

### Response an ein Microservice

Um ein Reponse zu senden, verwendet man die in den Parametern enthaltene `res *proto.ShowResponse`.

Beispielsweise wäre dies eine Antwort, um alle Reservierungen an den Client als message zurückzugeben.

    var allReservations string

    [...]

    res.Success = true
    res.Message = allReservations

### Aufbau und Regeln der Services

<img src="https://github.com/ob-vss-ws19/blatt-4-consal/blob/development/Blatt4%20(2).png" alt="Blatt4 Ueberblick"/>

Objekte in Services `user`, `movie` und `cinemahall` können unabhängig von anderen Microservices erstellt werden.

Um einen Eintrag in `show` zu erzeugen, muss mindestens ein `movie` und ein `cinemahall` vorhanden sein.

Um einen Eintrag in `reservation` zu erzeugen, muss mindestens ein `show` und ein `user` vorhanden sein.

### Verfügbare Funktionen

#### Funktion
- Funktionsname [parameter1] [parameter2] `Erklärung`

#### User

- AddUser [name] `Fügt einen neuen [user] hinzu`
- DeleteUser [name] `Löscht den [user] und die damit verknüpften [reservations]`
- GetUser `Gibt alle erstellten [user] zurück`

#### Movie

- AddMovie [name] `Fügt eine neue [movie] hinzu`
- DeleteMovie [name] `Löscht den [movie] und die damit verknüpften [shows]`
- GetMovie `Gibt alle erstellten [movies] zurück`

#### Cinemahall

- AddCinemahall [name] [seatrows] [seatrowcapacity] `Fügt eine neue [cinemahall] hinzu mit der Größe [seatrows] * [seatrowcapacity]`
- DeleteCinemahall [name] `Löscht den [cinemahall] und die verknüpften [shows]`
- GetCinemahall `Gibt alle erstellten [cinemahalls] zurück`

#### Show
- AddShow [moviename] [cinemahallname] `Verknüpft [moviename] & [cinemahallname] zu einer [show]`
- DeleteShow [Id] `Löscht eine [show] mit der id [Id]`
- GetShows `Gibt alle erstellten [shows] zurück`

#### Reservation

- CheckReservation [user] [movie] [seats] `Fragt eine Reservation für [user] für den Film [movie] mit [seats] Sitzplätzen an`
- MakeReservation [reservationId] `Reserviert endgültig die Reservation mitder id [reservationId]`
- DeleteReservation [reservationId] `Löscht eine [reservation] mit der id [id]`
- GetReservations`Gibt alle erstellten [reservations] zurück`
