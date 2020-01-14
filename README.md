# Reservierungssystem für ein Kino mit Microservices

## Funktionsbeschreibung
Es befinden sich 5 Microservices, die in `user, movie, cinemahall, showing` und `reservation` aufgeteilt sind.
Um auf die Microservices über die Kommandozeile zu benutzen, verwendet man den `client`.
Die Services können alle entweder einzeln oder über `startServicesLocal.sh` gestartet werden.
Zudem ist es möglich, über `startServicesDocker` die aktuellsten Dockercontainer aus develop auszuführen.

__Die Kommunikation zwischen den Microservices wird [hier](https://github.com/ob-vss-ws19/blatt-4-consal/blob/development/Protocol.md) genauer erklärt.__
## Let's Start!

-   das Github Repository klonen:

    ```
    git clone https://github.com/ob-vss-ws19/blatt-4-consal.git kino && cd kino
    ```

-  Starten der Services Lokal

    ```
    bash startServicesLocal.sh
    ```

-  Starten der Services Über Dockercontainer

    ```
    bash startServicesDocker.sh
    ```

-   Bauen des Clients Lokal:

    ```
    go build -o client cli/client.go
    ```

-   Starten des Clients und mit Beispieldaten füllen:

    ```
    ./client fill
    ```

-   Erhalte alle User:

    ```
    ./client us get
    ```

-   Füge neue Movie hinzu mit den Namen `Zohan`:

    ```
    ./client mv add Zohan
    ```
    `Siehe CLI-Commands für weitere ausführbare Befehle.`

-   Bauen des Clients über Docker und mit Beispieldaten füllem:

    ```
    docker run terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-consal:PR-3-client fill
    ```


## CLI-Commands

-   Über den Client können jedem Service `user, movie, cinemahall, showing` und `reservation` Daten hinzugefügt `add`, gelöscht `delete` und aufgelistet `get` werden.
Einzige Ausnahme bietet hierbei reservation, hier ist es nicht möglich einfach eine Reservierung hinzuzufügen, diese muss zunächst beantragt `check` werden und anschließend gebucht `make`.

### Befehlausführung

    ./client [Service] [Function] [Parameter]

oder


    docker run terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-consal:PR-3-client [Service] [Function] [Parameter]


- alle Befehle
```
Services: us, mv, cm, sw, rv
Functions: add, delete, get, ...

us
    Functions:
    - add <Parameter>: name. Example: client us add user1
    - delete <Parameter>: name. Example: client us delete user1
    - get: Example: client us get

mv
    Functions:
    - add <Parameter>: title. Example: client mv add movie1
    - delete <Parameter>: title. Example: client mv delete movie1
    - get: Example: client mv get

cm
    Functions:
    - add <Parameter>: name. Example: client cm add cine1
    - delete <Parameter>: name. Example: client cm delete cine1
    - get: Example: client cm get

sw
    Functions:
    - add <Parameter>: movie cm. Example: client sw add movie1 cine1
    - delete <Parameter>: showingID. Example: client sw delete 4
    - get: Example: client sw get

rv
    Functions:
    - check <Parameter>: user showingID seats. Example: client rv check user1 2 4
    Requests a reservation.
    - make <Parameter>: reservationID. Example: client rv make 1
    Books a reservation.
    - delete <Parameter>: reservationID. Example: client rv delete 1
    - get: Example: client reservation get

fill
    - Fills services with some data. Example: client fill
```
