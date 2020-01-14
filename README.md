# Reservierungssystem für ein Kino

## Beschreibung
Es befinden sich 6 Microservices, die in user, movie, cinemahall, showing, reservation und client aufgeteilt sind.
Die Kommunikation der Microservices wird über micro bereitgestellt.
Die Services können alle entweder einzeln oder über 'startServicesLocal.sh' gestartet werden.
Zudem ist es möglich, über 'startServicesDocker' die aktuellsten Dockercontainer aus develop auszuführen.

## Getting started

-   das Github Repository klonen:

    ```
    git clone https://github.com/ob-vss-ws19/blatt-4-consal.git kino && cd kino
    ```

-   Bauen des Clients:

    ```
    go build -o client cli/client.go
    ```

-   Starten des Clients:

    ```
    ./cli fill
    ```

-  Starten der Services Lokal

    ```
    ./startServicesLocal.sh
    ```

  Starten der Services Über Dockercontainer

    ```
    ./startServicesDocker.sh
    ```

## CLI-Commands

-   Über den Client können jedem Service (cinema, movie, reservation, showing und user) Daten hinzugefügt (add), gelöscht (delete) und aufgelistet (get) werden.
Einzige Ausnahme bietet hierbei reservation, hier ist es nicht möglich einfach eine Reservierung hinzuzufügen, diese muss zunächst beantragt (check) werden und anschließend gebucht (make).

    ```
    Commands:

    us = user
    mv = movie
    cm = cinemahall
    sw = show
    rv = reservation

    client <SERVICE> <FUNCTION> <PARAMS>

    <SERVICE>
    us
        <FUNCTION>
        - add <PARAMS>: name. | Example: client us add user1
        - delete <PARAMS>: name. | Example: client us delete user1
        - get: | Example: client us get
    mv
        <FUNCTION>
        - add <PARAMS>: title. | Example: client mv add movie1
        - delete <PARAMS>: title. | Example: client mv delete movie1
        - get: | Example: client mv get

    cm
        <FUNCTION>
        - add <PARAMS>: name. | Example: client cm add cine1
        - delete <PARAMS>: name. | Example: client cm delete cine1
        - get: | Example: client cm get
    sw
        <FUNCTION>
        - add <PARAMS>: movie cm. | Example: client sw add movie1 cine1
        - delete <PARAMS>: showingID. | Example: client sw delete 4
        - get: | Example: client sw get
    rv
        <FUNCTION>
        - check <PARAMS>: user showingID seats. | Example: client rv check user1 2 4
        Requests a reservation.
      - make <PARAMS>: reservationID. | Example: client rv make 1
        Books a reservation.
      - delete <PARAMS>: reservationID. | Example: client rv delete 1
      - get: | Example: client reservation get
    fill
      - Fills services with some data. | Example: client fill
      ```
