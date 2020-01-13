FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/movieservice/main services/movieservice/main.go

FROM iron/go
COPY --from=builder /app/services/movieservice/main /app/services/movieservice
EXPOSE 3001
ENTRYPOINT [ "/app/movieservice" ]
