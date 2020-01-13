FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/reservationservice/main services/reservationservice/main.go

FROM iron/go
COPY --from=builder /app/services/reservationservice/main /app/services/reservationservice
EXPOSE 3004
ENTRYPOINT [ "/app/reservationservice" ]
