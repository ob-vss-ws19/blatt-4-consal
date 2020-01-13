FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/userservice/main services/userservice/main.go

FROM iron/go
COPY --from=builder /app/services/userservice/main /app/services/userservice
EXPOSE 3000
ENTRYPOINT [ "/app/userservice" ]
