FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o reservationservice/main reservationservice/main.go

FROM iron/go
COPY --from=builder /app/reservationservice/main /app/reservationservice
EXPOSE 3004
ENTRYPOINT [ "/app/reservationservice" ]
